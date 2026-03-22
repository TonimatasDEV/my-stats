package util

import (
	"io"
	"log"
	"net/http"
	"time"
)

func IsNotNil(err error) bool {
	if err != nil {
		log.Println("Error making request:", err)
		time.Sleep(1 * time.Minute)
		return true
	}

	return false
}

func IsNotOk(resp *http.Response) bool {
	if resp.StatusCode != 200 {
		log.Printf("Error getting CF Widget data: %d\n", resp.StatusCode)
		CloseBody(resp.Body)
		time.Sleep(1 * time.Minute)
		return true
	}

	return false
}

func CloseBody(body io.ReadCloser) {
	err := body.Close()

	if err != nil {
		log.Println("Error closing body:", err)
	}
}
