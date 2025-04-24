package main

import (
	"log"
	"net/http"
	"time"
)

func RequestToSite() map[string]int {
	mapResponse := make(map[string]int)

	urls := make([]string, len(getUrls()))
	for i, j := range getUrls() {
		str, err := j.(string)
		if !err {
			log.Fatalf("Некоректный тип URL:%T", j)
		}
		urls[i] = str
	}

	client := http.Client{
		Timeout: time.Duration(getTimeout()) * time.Second,
	}
	for _, url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			log.Printf("Ошибка при запросе к %s: %v\n", url, err)
			continue
		}

		status := resp.StatusCode
		mapResponse[url] = status
		log.Printf("Запрос к сайту успешен: url: %s, status: %d ", url, status)
		resp.Body.Close()
	}
	return mapResponse
}
