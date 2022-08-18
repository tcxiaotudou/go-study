package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Effector func(context.Context) (string, error)

func Retry(effector Effector, retries int, delay time.Duration) Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				// Return when there is no error or the maximum amount
				// of retries is reached.
				return response, err
			}
			log.Printf("Function call failed, retrying in %v", delay)
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}

var count int

func GetPdfUrl(ctx context.Context) (string, error) {
	count++
	if count <= 3 {
		return "", errors.New("boom")
	} else {
		return "https://linktopdf.com", nil
	}
}

func main() {
	r := Retry(GetPdfUrl, 5, 2*time.Second)
	res, err := r(context.Background())
	fmt.Println(res, err)
}
