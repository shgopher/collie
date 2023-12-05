package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/googege/collie/mem"
	"github.com/googege/gotools/id"
	"github.com/nfnt/resize"
)

var (
	root      string // 输入
	outPath   string // 输出
	outPutYes int    // 是否跟源文件保持一致的名称
	width     int    //宽度
	quality   int    // 质量
)

type XC struct {
	img  image.Image
	name string
}

// get file's path
func retrieveData(root string) (value chan string, err chan error) {
	err = make(chan error, 1)
	value = make(chan string)
	go func() {
		defer close(value)
		err <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// if the file is noe regular, it mean the file is done,you should return
			if !info.Mode().IsRegular() {
				return nil
			}
			value <- path
			return nil
		})
	}()
	return
}

// get file send to a chan.
func ReceiveData(file chan string, value chan io.Reader, wg *sync.WaitGroup) {
	for v := range file {
		dif, err := mem.MemDifference()
		if err != nil {
			fmt.Println(err)
		}
		if dif > 0.2 {
			time.Sleep(time.Second >> 1)
			fmt.Println("waiting for mem less.")
		}
		fi, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
		} else {
			value <- fi
		}
	}
	wg.Done()
}

// resize and create a new photo with only id name.
func DataProcessing(root string, outputFile string, wid int, q int) {
	reader := make(chan io.Reader)
	b := make(chan *XC)
	c := make(chan *XC)
	value, err := retrieveData(root)
	//
	wg := new(sync.WaitGroup)
	wg.Add(2)
	for i := 0; i < 2; i++ {
		mark(i, "获取文件路径：")
		go ReceiveData(value, reader, wg)
	}
	go func() {
		wg.Wait()
		close(reader)
	}()
	//
	wg1 := new(sync.WaitGroup)
	wg1.Add(32)
	for i := 0; i < 32; i++ {
		go func(i int) {
			defer wg1.Done()
			mark(i, "正在解析")
			for r := range reader {
				v, ok := r.(*os.File)
				if !ok {
					glog.Errorln("not photo")
				}
				_, name1 := filepath.Split(v.Name())
				// name1 是文件的名字，name是后缀
				name := findName(name1)
				if name == "" && name1 != ".DS_Store" {
					glog.Errorln("not file,the file name is ", name)
				}
				img, err := isJpg(name, r)
				if err != nil {
					glog.Errorln("无法读取文件：", name1, err)
				} else {
					b <- &XC{
						img:  img,
						name: name1,
					}
				}
			}
		}(i)
	}
	go func() {
		wg1.Wait()
		close(b)
	}()
	//
	wg2 := new(sync.WaitGroup)
	wg2.Add(32)
	for i := 0; i < 32; i++ {
		go func(i int) {
			mark(i, "正在压缩")
			defer wg2.Done()
			for i := range b {
				i.img = resize.Resize(uint(wid), 0, i.img, resize.NearestNeighbor)
				c <- i
			}
		}(i)
	}
	go func() {
		wg2.Wait()
		close(c)
	}()
	//
	wg3 := new(sync.WaitGroup)
	wg3.Add(32)
	for i := 0; i < 32; i++ {
		go func(i int) {
			mark(i, "正在创建文件：")
			defer wg3.Done()
			for i := range c {
				defaultName := ""
				if outPutYes == 0 {
					defaultName = i.name
				} else {
					defaultName = onlyID1() + ".jpeg"
				}
				file, err := os.Create(outputFile + "/" + defaultName)
				defer file.Close()
				stat, _ := file.Stat()
				fmt.Printf("成功输出文件:%s\n", stat.Name())
				if err != nil {
					fmt.Println(err)
				}
				if q < 20 {
					q = 20
				}
				if err := jpeg.Encode(file, i.img, &jpeg.Options{q}); err != nil {
					glog.Errorln("photo creating process is error:", err)
				}
			}
		}(i)
	}
	//
	if er := <-err; er != nil {
		fmt.Println("can not find file ,or no order to find", er)
	}
	//
	wg3.Wait()
}

// workNode is the computer's name if you have so many computers.
func onlyID() string {
	snow, err := id.NewSnowFlake(1)
	if err != nil {
		fmt.Println(err)
	}
	glog.V(1).Info("use snowFlake")
	return strconv.FormatInt(snow.GetID(), 10)
}
func onlyID1() string {
	u, err := id.NewUUID(id.VERSION_1, nil)
	if err != nil {
		glog.Error(err)
	}
	return u.String()
}
func findName(name string) string {
	name = strings.ToLower(name)
	v := name[len(name)-4:]
	v1 := name[len(name)-3:]
	if v == "jpeg" {
		return v
	}
	if v1 == "jpg" || v1 == "png" || v1 == "gif" {
		return v1
	}
	return ""
}
func isJpg(name string, r io.Reader) (image.Image, error) {
	switch name {
	case "jpeg", "jpg":
		return jpeg.Decode(r)
	case "png":
		return png.Decode(r)
	case "gif":
		return gif.Decode(r)
	default:
		return nil, fmt.Errorf("本程序只能压缩 jpg jpeg png 和gif，并且最后输出的都是jpeg文件，望周知")
	}
}

func mark(i int, name string) {
	if i == 0 {
		fmt.Printf("%s\n", name)
	}
}

func init() {
	flag.StringVar(&root, "r", "./test", "指定的输入路径，路径是指的图片所处的文件夹，文件夹中还有文件夹不影响，系统会找到你指定文件夹中的所以照片，包括文件夹中的文件夹里的图片")
	flag.StringVar(&outPath, "o", ".", "输出的路径")
	flag.IntVar(&width, "w", 0, "输出的照片尺寸，0是跟之前一样大，单位是px")
	flag.IntVar(&quality, "q", 75, "输出的照片质量，范围是从1 - 100")
	flag.IntVar(&outPutYes, "n", 0, "是否输出跟源文件相同的名称，0：是，1：不是，如果不是，系统会给出一个随机代码，默认是输出相同的名字")
	flag.Parse()
}
func main() {
	fmt.Println("声明：本程序来自GitHub：shgopher,欢迎关注公众号：科科人神；\n免费软件，如果使用期间出现任何后果，本软件不承担任何责任谢谢\n")
	fmt.Println("程序正式开始运行 🚀🚀🚀")
	DataProcessing(root, outPath, width, quality)
	fmt.Println("运行结束 ☕️ ☕ ☕\n")
	fmt.Printf("您可以打开%s去查看已经压缩好的文件\n", outPath)
}
