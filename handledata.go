package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main() {
	rd, err := loadDataJson("./data.json")
	if err != nil {
		panic(err.Error())
	}
	//单位是天，每天的峰值，开始时间是当天的0:0:0
	minVo := rd.Datas[0]
	maxVo := rd.Datas[len(rd.Datas)-1]
	mints := int64(minVo["timestamp"].(float64)) // 开始时间0:0:0
	maxts := int64(maxVo["timestamp"].(float64)) // 截止时间点
	var daytime int64
	daytime = 1000 * 60 * 60 * 24     // 一天毫秒数
	days := (maxts - mints) / daytime // 天数
	ms := (maxts - mints) % daytime   // 剩余毫秒数
	times := ms / 1000 / 60 / 60      // 转成小时
	if times > 1 { // 如果大于一小时，天数加1
		days++
	}
	// 根据天数得到每天0:0:0的时间戳
	rs := new(ResultData)
	start := mints
	for d := 1; d <= int(days); d++ {
		list := make([]HM, 0)
		end := start + daytime
		for _, vo := range rd.Datas {
			ts := int64(vo["timestamp"].(float64))
			if ts >= start && ts < end {
				list = append(list, vo)
			}
		}
		start = end
		var hm HM
		hm = list[0]
		maxVal := int64(list[0]["hourTotalAccess"].(float64))
		for i := 1; i < len(list); i++ { // 求最大值
			val := int64(list[i]["hourTotalAccess"].(float64))
			if val > maxVal {
				maxVal = val
				hm = list[i]
			}
		}
		rs.Datas = append(rs.Datas, hm)
	}
	j, err := json.Marshal(rs)
	fmt.Println(string(j))
}

type HM map[string]interface{}
type ResultData struct {
	Datas []HM `json:"datas"`
}

func loadDataJson(filename string) (*ResultData, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read es query file failed:%s", err.Error())
	}
	rd := new(ResultData)
	if err = json.Unmarshal(raw, rd); err != nil {
		return nil, fmt.Errorf("json parse failed:%s", err.Error())
	}
	return rd, nil
}