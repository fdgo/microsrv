package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	l.PushBack(36)
	l.PushBack(1)
	l.PushBack("beijing")
	l.PushBack(3)
	fmt.Println("original list:")
	prtList(l)
//------------------------
	delete(l,36)
//------------------------
	fmt.Println("deleted list:")
	prtList(l)
}
func delete(l *list.List, obj interface{})  {
	var next *list.Element
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		if e.Value == obj{
			l.Remove(e)
		}
	}
}
func deleteAll(l *list.List)  {
	var next *list.Element
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}

}
func prtList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Printf("\n")
}
