package main

import (
	"fmt"
	"reflect"
)
//https://www.cnblogs.com/skymyyang/p/7690837.html
type User struct {
	Id   int
	Name string
	Age  int
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
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	//args := []reflect.Value{reflect.ValueOf("JOE")}
	mv.Call(nil)
}