package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	exampleJsonString := `{
    "code":"000",
    "data":{
			"all_count":441353,
			"lists":[
				{
					"id":441353,
					"job_name":"经营日报-同步职位信息",
					"job_recall_time":"2021-03-13 15:05:04",
					"job_recall_content":"请求成功：great",
					"create_time":"2021-03-13 15:05:04"
				},
				{
					"id":441352,
					"job_name":"经营日报-Check张继学列表",
					"job_recall_time":"2021-03-13 15:05:00",
					"job_recall_content":"请求成功：OK",
					"create_time":"2021-03-13 15:05:00"
				}
			]
		},
		"msg":"获取列表成功",
		"success":true
	}`

	//fmt.Println(gjson.GetBytes(exampleJsonByte, "data.lists.#.job_name").Array())
	jsonCode := gjson.Get(exampleJsonString, "code")
	fmt.Println(jsonCode)

	jsonOneJobName := gjson.Get(exampleJsonString, "data.lists.#.job_name").Array()
	fmt.Println(jsonOneJobName)

	gjson.Get(exampleJsonString, "data.lists.1").ForEach(func(key, value gjson.Result) bool {
		fmt.Println(value)
		return true
	})
}
