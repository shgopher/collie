package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/googege/gotools/id"
	"github.com/nfnt/resize"
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
)
var (
	root string
	outPath string
	width int
	quality int
)
func init(){
	flag.StringVar(&root,"root",".","root path")
	flag.StringVar(&outPath,"outPath",".","out put dir")
	flag.IntVar(&width,"width",0,"picture widh")
	flag.IntVar(&quality,"quality",75,"quality of the picture")
}
func main() {
	fmt.Println("collie is runing...🚀")
	DataProcessing(root, outPath, width, quality)
}

func retrieveData(root string) (value chan string, err chan error) {
	err = make(chan error, 1)
	value = make(chan string)
	go func() {
		defer close(value)
		err <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// if the file is noe regular, it mean the file is done,you should return .
			if !info.Mode().IsRegular() {
				return nil
			}
			value <- path
			return nil
		})
	}()
	return
}
func ReceiveData(file chan string, value chan io.Reader, wg *sync.WaitGroup) {
	for v := range file {
		fi, err := os.Open(v)
		if err != nil {
			glog.Errorln(err)
			continue
		}
		value <- fi
	}
	wg.Done()
}
func DataProcessing(root string, outputFile string, wid int, q int) {

	reader := make(chan io.Reader)
	b := make(chan image.Image)
	c := make(chan image.Image)
	value, err := retrieveData(root)
	wg := new(sync.WaitGroup)
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go ReceiveData(value, reader, wg)
	}
	go func() {
		wg.Wait()
		close(reader)
	}()
	ccc := make(chan bool, 20)
	wg1 := new(sync.WaitGroup)
	wg1.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg1.Done()
			for {
				select {
				case r, ok := <-reader:
					if ok {
						v,ok := r.(*os.File)
						if !ok {
							glog.Errorln("not file.")
						}
						_,name1 := filepath.Split(v.Name())
						name := findName(name1)
						if name == "" && name1 != ".DS_Store"{
							fmt.Println(name1)
							glog.Errorln("not file..")
						}
						img, err := isJpg(name,r)
						if err != nil {
							glog.Errorln(err)
						} else {
							b <- img
						}
					} else {
						ccc <- true
					}
				case <-ccc:
					return
				}
			}
		}()
	}
	go func() {
		wg1.Wait()
		close(b)
	}()
	ddd := make(chan bool, 20)
	wg2 := new(sync.WaitGroup)
	wg2.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg2.Done()
			for {
				select {
				case i, ok := <-b:
					if ok {
						c <- resize.Resize(uint(wid), 0, i, resize.NearestNeighbor)
					} else {
						ddd <- true
					}
				case <-ddd:
					return
				}
			}
		}()
	}
	go func() {
		wg2.Wait()
		close(c)
	}()
	ttt := make(chan bool, 20)
	wg3 := new(sync.WaitGroup)
	wg3.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg3.Done()
			for {
				select {
				case i, ok := <-c:
					if ok {
						file, err := os.Create(outputFile + "/" + onlyID() + ".jpeg")
						if err != nil {
							glog.Error(err)
						}
						if q < 20 {
							q = 20
						}
						jpeg.Encode(file, i, &jpeg.Options{
							q,
						})
					} else {
						ttt <- true
					}
				case <-ttt:
					return
				}
			}
		}()
	}
	if er := <-err; er != nil {
		glog.Error("错误over：", err)
	}
	wg3.Wait()
}

// workNode is the computer's name if you have so many computers.
func onlyID() string {
	snow, err := id.NewSnowFlake(1)
	if err != nil {
		glog.Error(err)
	}
	glog.V(1).Info("use snowFlake")
	return strconv.FormatInt(snow.GetID(), 10)
}
func onlyID1() string {
	u, err := id.NewUUID(id.VERSION_1, nil)
	if err != nil {
		glog.Error(err)
	}
	glog.V(1)
	return u.String()
}
func findName(name string)string {
	v := name[len(name)-4:]
	v1 := name[len(name)-3:]
	if v == "jpeg"{
		return v
	}
	if v1 == "jpg"|| v1 == "png" || v1 == "gif" {
		return v1
	}
	return ""
}
func isJpg(name string,r io.Reader,)(image.Image,error){
	name = strings.ToLower(name)
	switch name {
	case "jpeg","jpg":
		return jpeg.Decode(r)
	case "png":
		return png.Decode(r)
	case "gif":
		return gif.Decode(r)
	default:
		return nil,fmt.Errorf("just can use jpeg jpg png and gif")
	}
}