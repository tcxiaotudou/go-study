package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/**
{
    "company":"武汉闪达科技",
    "username":"sendcloud_admin",
    "timestamp":1661927656950,
    "page":1,
    "size":10,
    "signature":"c679640e792cede5ae9c5c26f66a0c96ec6041b5"
}'
*/

type Request struct {
	Company   string `json:"company"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Timestamp int64  `json:"timestamp"`
	Username  string `json:"username"`
	Signature string `json:"signature"`
}

func HMACSHA1(key, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func main() {
	timestamp := time.Now().UnixNano() / 1e6
	str := fmt.Sprintf("company=武汉闪达科技&page=1&size=10&timestamp=%d&username=", timestamp)
	fmt.Println(str)
	quest := Request{
		Company:   "武汉闪达科技",
		Username:  "",
		Timestamp: timestamp,
		Page:      1,
		Size:      10,
		Signature: HMACSHA1("746310e38d994cbfa4dc9669251eab66", str),
	}
	body, _ := json.Marshal(quest)
	fmt.Println(string(body))
	resp, err := http.Post("http://scdata.sendcloud.io/crm/query", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		jsonStr := string(body)
		log.Println("Response: ", jsonStr)
	} else {
		fmt.Println("Get failed with error: ", resp.Status)
	}
}
