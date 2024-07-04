package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ReadResponseBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %s\n", err.Error())
		return ""
	}
	return string(body)
}

func CloseResponseBody(Body io.ReadCloser) {
	if err := Body.Close(); err != nil {
		fmt.Printf("Error closing response body: %v\n", err)
	}
}
