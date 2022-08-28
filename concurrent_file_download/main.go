package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	url     = "http://212.183.159.230/5MB.zip"
	workers = 5
)

type Part struct {
	Data  []byte
	Index int
}

func main() {
	var size int
	results := make(chan Part, workers)
	parts := [workers][]byte{}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Headers : ", resp.Header["Content-Length"])

	if header, ok := resp.Header["Content-Length"]; ok {
		fileSize, err := strconv.Atoi(header[0])
		if err != nil {
			log.Fatal("File size could not be determined: ", err)
		}
		size = fileSize / workers
	} else {
		log.Fatal("File size was not provided")
	}

	for i := 0; i < workers; i++ {
		go download(i, size, results)
	}

	counter := 0

	for part := range results {
		counter++
		parts[part.Index] = part.Data
		if counter == workers {
			break
		}
	}

	var file []byte

	for _, part := range parts {
		file = append(file, part...)
	}

	// Set permissions accordingly, 0700 may not
	// be the best choice
	err = os.WriteFile("./data.zip", file, 0700)

	if err != nil {
		log.Fatal(err)
	}
}

func download(index, size int, c chan Part) {
	client := &http.Client{}

	// calculate offset by multiplying
	// index with size
	start := index * size

	// Write data range in correct format
	// I'm reducing one from the end size to account for
	// the next chunk starting there
	dataRange := fmt.Sprintf("bytes=%d-%d", start, start+size-1)

	// if this is downloading the last chunk
	// rewrite the header. It's an easy way to specify
	// getting the rest of the file
	if index == workers-1 {
		dataRange = fmt.Sprintf("bytes=%d-", start)
	}

	log.Println(dataRange)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// code to restart download
		return
	}

	req.Header.Add("Range", dataRange)
	resp, err := client.Do(req)
	if err != nil {
		// code to restart download
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		// code to restart download
		return
	}

	c <- Part{
		Data:  body,
		Index: index,
	}
}
