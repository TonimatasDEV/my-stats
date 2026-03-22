package util

import (
	"io"
	"log"
	"net/http"
)

func IsNotNil(err error) bool {
	if err != nil {
		log.Println("Error making request:", err)
		return true
	}

	return false
}

func IsNotOk(resp *http.Response) bool {
	if resp.StatusCode != 200 {
		log.Printf("Error getting CF Widget data: %d\n", resp.StatusCode)
		CloseBody(resp.Body)
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
