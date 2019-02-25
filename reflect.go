package main

import (
	"fmt"
	"reflect"
)
//https://www.cnblogs.com/skymyyang/p/7690837.html
type User struct {
	Id   string
	Name string
	Age  string
}

func (u User) Hello() {
	fmt.Println(`
                How are you?
                Fine, thank you. And you?
                I'm fine too.
        `)
	//fmt.Println("Hello", "My name is", u.Name)
}

func main() {
	//u1 := User{"1", "OK", "12"}
	u1 := new(User)
	u1.Id = "1"
	u1.Age = "23"
	fmt.Println(reflect.TypeOf(u1)) //*main.User 获取任意对象的具体类型
	if reflect.TypeOf(u1).Kind() == reflect.Ptr { //判断是否属于指针类型
		fmt.Println("type ptr")
	}
	u2 := new(User)
	u2.Name = "sam"
	//fmt.Println(reflect.TypeOf(u1).Field(1).Name) //panic: reflect: Field of non-struct type
	//reflect.ValueOf函数返回的是一份值的拷贝，所以前提是我们是传入要修改变量的地址。其次需要我们调用Elem方法找到这个指针指向的值。
	v1 := reflect.ValueOf(u1).Elem()
	v2 := reflect.ValueOf(u2).Elem()
	t1 := reflect.TypeOf(*u1)
	for i := 0; i < v1.NumField(); i ++ {
		v := v1.Field(i)
		if v.Type().Kind() == reflect.String && v.String() != "" {
			fmt.Println(t1.Field(i).Name, v.String()) //字段名 字段值
			v2.Field(i).SetString(v.String())
		}
	}
	fmt.Println(u2)
	//args := []reflect.Value{reflect.ValueOf("JOE")}
	mv := v1.MethodByName("Hello")
	mv.Call(nil)
}