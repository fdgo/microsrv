package channel

import (
	"fmt"
	"errors"
	"time"
)

func ReadNoDataFromNoBufChWithSelect()  {
	bufch := make(chan int)
	if v,err := ReadWithSelect(bufch);err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("read: %d\n",v)
	}
}

func ReadNoDataFromBufChWithSelect()  {
	bufch := make(chan int, 1)
	if v,err := ReadWithSelect(bufch);err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("read: %d\n",v)
	}
}

func ReadWithSelect(ch chan int) (x int, err error) {
	timeout := time.NewTimer(time.Microsecond*500)
	select {
	case x= <-ch:
		return x,nil
	case<- timeout.C:
		return 0,errors.New("read time out!")
	}
}
//----------------------------------------------------------------------------------------------------------------------
func WriteNoBufChWithSelect()  {
	ch := make(chan int)
	if err := WriteChWithSelect(ch);err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("write success")
	}
}

func WirteBufChButFullWithSelect()  {
	ch := make(chan int,1)
	ch <- 100
	if err := WriteChWithSelect(ch); err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("write success")
	}
}

func WriteChWithSelect(ch chan int) error {
	timeout := time.NewTimer(time.Microsecond*500)
	select {
	case ch<-1:
		return nil
	case <-timeout.C:
		return errors.New("write time out!")

	}
}