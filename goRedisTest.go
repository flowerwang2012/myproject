package main

import (
	"flag"
	"strings"
	"github.com/go-redis/redis"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

//redis哨兵模式测试 go-redis驱动
//nohup ./goRedisTest --master_name "mymaster" --sentinels "127.0.0.1:26380,127.0.0.1:26381,127.0.0.1:26382" > goRedisTest.log 2>&1 &
//curl 127.0.0.1:18999
var (
	client *redis.Client
)
func main() {
	masterNamePtr := flag.String("master_name", "", "master name")
	sentinelsPtr := flag.String("sentinels", "", "ip:port,ip:port")
	flag.Parse()
	masterName := *masterNamePtr
	sentinels := *sentinelsPtr

	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: masterName,
		SentinelAddrs: strings.Split(sentinels, ","),
	})
	if err := client.Ping().Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	e := echo.New()
	e.GET("/set", handleSet)
	e.GET("/del", handleDel)
	e.GET("/get", handleGet)
	e.Start("127.0.0.1:18999")
}

func handleSet(c echo.Context) error {
	s, err := client.Set("name", "wanghongfa", 0).Result()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, s)
}

func handleDel(c echo.Context) error {
	s, err := client.Del("name").Result()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, s)
}

func handleGet(c echo.Context) error {
	s, err := client.Get("name").Result()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, s)
}