package http

import (
	"fmt"
	"net/http"
	"time"
)

func UseHttp() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	for true {
		bytes := make([]byte, 1014)
		n, err := response.Body.Read(bytes)
		fmt.Println(string(bytes[:n]))
		if n == 0 || err != nil {
			break
		}
	}
}

func UseHttpClient() {
	// Get
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := httpClient.Get("https://google.com")
	if err != nil {
		return
	}
	fmt.Println(response.Body)
	defer response.Body.Close()
}
