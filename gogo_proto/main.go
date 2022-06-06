package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"log"
)

func main() {

	req := &SearchRequestParam{
		QueryText: "test text",
		Limit:     10,
		Type:      SearchRequestParam_PC,
	}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalln("Marshal err: ", err)
	}
	// send data
	fmt.Println(string(data))

	var respData []byte
	var result = SearchRequestParam{}
	if err = proto.Unmarshal(respData, &result); err != nil {
		fmt.Println(result)
	} else {
		log.Fatalln("Unmarshal err : err")
	}
}
