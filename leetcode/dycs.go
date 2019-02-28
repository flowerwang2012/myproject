package main

import (
	"net/http"
	"encoding/json"
)

//单元测试
func Add(a,b int) int{
	return a+b
}

//模拟调用
/*
单元测试的原则，就是你所测试的函数方法，不要受到所依赖环境的影响，比如网络访问等，因为有时候我们运行单元测试的时候，并没有联网，那么总不能让单元测试因为这个失败吧？所以这时候模拟网络访问就有必要了。
针对模拟网络访问，标准库了提供了一个httptest包，可以让我们模拟http的网络调用，下面举个例子了解使用。
首先我们创建一个处理HTTP请求的函数，并注册路由
非常简单，这里是一个/sendjsonAPI，当我们访问这个API时，会返回一个JSON字符串。现在我们对这个API服务进行测试，但是我们又不能时时刻刻都启动着服务，所以这里就用到了外部终端对API的网络访问请求。
*/
func Routes(){
	http.HandleFunc("/sendjson",SendJSON)
}

func SendJSON(rw http.ResponseWriter,r *http.Request){
	u := struct {
		Name string
	}{
		Name:"张三",
	}

	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)
}
