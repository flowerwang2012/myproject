package main

import "fmt"

//Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
// 因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样就在函数中就无法修改原内容数据；
// 有的是引用类型（指针、map、slice、chan等这些），这样就可以修改原内容数据。
// https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html

func main() {
	//与new不同的，make 返回类型的引用而不是指针
	m := make(map[string]int)
	m["num"] = 1
	fmt.Printf("原始map的值是：%+v\n", m)
	modifyMap(m)
	fmt.Printf("修改后map的值是：%+v\n", m)
	s := []string{"hello"}
	fmt.Printf("原始slice的值是：%+v\n", s)
	modifySlice(s)
	fmt.Printf("修改后slice的值是：%+v\n", s)
	c := make(chan int, 2)
	fmt.Printf("原始chan 未被读取的len：%d 容量：%d\n", len(c), cap(c))
	modifyChan(c)
	fmt.Printf("修改后chan 未被读取的len：%d 容量：%d\n", len(c), cap(c))

	i := "old interface"
	modifyInterface(i)
	fmt.Println(i)
	f := func() {
		fmt.Println("old func")
	}
	modifyFunc(f)
	f()
}

//func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap {
//    //省略无关代码
//}
//通过查看src/runtime/hashmap.go源代码发现，的确和我们猜测的一样，make函数返回的是一个hmap类型的指针*hmap。
//也就是说map===*hmap。 现在看func modify(p map)这样的函数，其实就等于func modify(p *hmap)
func modifyMap(m map[string]int) {
	m["num"] = 2
}

//func makechan(t *chantype, size int64) *hchan {
//    //省略无关代码
//}
//chan也是一个引用类型，和map相差无几，make返回的是一个*hchan
func modifyChan(c chan int) {
	c <- 1
}

//type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}
//
//type slice struct {
//	array unsafe.Pointer
//	len   int
//	cap   int
//}
//slice也是引用类型
func modifySlice(s []string) {
	s[0] = "hello world"
}

func modifyInterface(i interface{}) {
	i = "new interface"
}

func modifyFunc(f func()) {
	f = func() {
		fmt.Println("new func")
	}
}

