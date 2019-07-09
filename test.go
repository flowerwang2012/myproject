package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
}
type CRJAuthRsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data interface{} `json:"data"`
}
type CRJPerson struct {
	FullName string `json:"fullName"`
	IdNum string `json:"idNum"`
	IdType string `json:"idType"`
	Nation string `json:"nation"`
	Sex string `json:"sex"`
	BirthDate string `json:"birthDate"`
	ExpiryDate string `json:"expiryDate"`
}

func main() {
	const (
		mutexLocked = 1 << iota // mutex is locked
		mutexWoken
		mutexStarving
		mutexWaiterShift = iota
	)
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
	str := "{\"data\":{\"fullName\":\"姓名\",\"idNum\":\"证件号码\",\"idType\":\"1001\",\"nation\":\"CHN\",\"sex\":\"1\",\"birthDate\":\"19930606\",\"expiryDate\":\"20281203\"},\"errcode\":0,\"errmsg\":\"\",\"hint\":\"1LR/F3fLDMYi\"}"
	rsp := new(CRJAuthRsp)
	if err := json.Unmarshal([]byte(str), rsp); err != nil {
		fmt.Println(err)
	}
	b, _ := json.Marshal(rsp.Data)
	p := new(CRJPerson)
	err := json.Unmarshal(b, p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	i := 0
	s := string(i)
	fmt.Println("s:", s)

	sli := []int{0, 1, 2}
	appendSlice(sli)
	fmt.Println("值传递", sli)
}

func appendSlice(i []int) {
	i[0] = 123
	i = append(i, 123)
	fmt.Println("after append", i)
}