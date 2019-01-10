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
	fmt.Println(reflect.TypeOf(u1))
	if reflect.TypeOf(u1).Kind() == reflect.Ptr {
		fmt.Println("type ptr")
	}
	u2 := new(User)
	u2.Name = "sam"
	ru1 := reflect.ValueOf(u1).Elem()
	ru2 := reflect.ValueOf(u2).Elem()
	fmt.Println(ru1.NumField())
	for i := 0; i < ru1.NumField(); i ++ {
		v := ru1.Field(i)
		t := v.Type()
		if t.Kind() == reflect.String && v.String() != "" {
			fmt.Println(v.String())
			ru2.Field(i).SetString(v.String())
		}
	}
	fmt.Println(u2)
	//args := []reflect.Value{reflect.ValueOf("JOE")}
	mv := ru1.MethodByName("Hello")
	mv.Call(nil)
}