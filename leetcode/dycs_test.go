package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"log"
)

//go test -v

func TestAdd(t *testing.T) {
	sum := Add(1,2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}

/*
运行这个单元测试，就可以看到我们访问/sendjsonAPI的结果里，并且我们没有启动任何HTTP服务就达到了目的。
这个主要利用httptest.NewRecorder()创建一个http.ResponseWriter，模拟了真实服务端的响应，这种响应时通过调用http.DefaultServeMux.ServeHTTP方法触发的。
 */
func init()  {
	Routes()
}

func TestSendJSON(t *testing.T){
	req,err:=http.NewRequest(http.MethodGet,"/sendjson",nil)
	if err!=nil {
		t.Fatal("创建Request失败")
	}

	rw:=httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw,req)

	log.Println("code:",rw.Code)

	log.Println("body:",rw.Body.String())
}