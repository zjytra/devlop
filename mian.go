package main

import (
	"fmt"
	"github.com/zjytra/devlop/xcontainer/vector"
	"sync"
	"time"
)

var wg sync.WaitGroup
var syncvec *vector.SyncVector
var safeVec *vector.SafeVector
func main() {
	syncvec = vector.NewSyncVector(0)
	safeVec = vector.NewSafeVector(0)
	wg.Add(5)
	go Custom()

	go Add()
	go Add()
	go Add()
	go Add()
	wg.Wait()

}


func Custom(){
	defer  wg.Done()
	for  {

		data := syncvec.WaitPop()
		fmt.Println("Custom",data)
		if syncvec.IsClose() {
			break
		}
	}
}

func Add(){
	defer  wg.Done()
	for i := 0;i <10 ;i++ {
		if i % 2 == 0 {
			time.Sleep(time.Second * 2)
		}
		if i == 5 {
			syncvec.Close()
		}
		syncvec.PushBack(i)
	}
	time.Sleep(time.Second * 2)
	syncvec.Close()
}