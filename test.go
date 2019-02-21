package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
)

func main() {
	//for i := 0; i < 10; i ++ {
	//	go Request()
	//}
	for i := 0; i < 100000; i++ {
		Request(i)
	}
	time.Sleep(3 * time.Second)
}

func Request(i int) {
	var (
		resp *http.Response
		err  error
	)
	fmt.Println(i)
	req, _ := http.NewRequest("GET", "https://app.gzgjj.gov.cn/user/login?yhhm=440823199701186225&yhmm=%E5%8D%9C%E7%A7%8B%E8%B4%B5&dllx=5&version=1.0&key=876CF952CE39D0695E9F01667055D8D4", nil)
	client := &http.Client{}
	if resp, err = client.Do(req); err != nil {
		fmt.Printf("request fail %s, %+v \n", err.Error(), req)
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("ebus http %s status code=%d \n", req.URL.Path, resp.StatusCode)
		panic(err)
	}

}