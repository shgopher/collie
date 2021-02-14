package mem

import (
	"github.com/golang/glog"
	"github.com/shirou/gopsutil/mem"
	"sync"
)

var (
	oldMem float64
)

func memPercent() (float64, error) {
	vm, err := mem.VirtualMemory()
	return vm.UsedPercent / 100, err
}

func init() {
	o := new(sync.Once)
	o.Do(func() {
		var err error
		oldMem, err = memPercent()
		if err != nil {
			glog.Error(err)
		}
	})
}

// return the difference between the old mem and the new one.
func MemDifference() (float64, error) {
	newMem, err := memPercent()
	return newMem - oldMem, err
}
