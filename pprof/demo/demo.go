package demo

import (
	"log"
	"myproject/pprof/demo/data"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/flowerwang2012"))
		}
	}()
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal(err)
	}
}
