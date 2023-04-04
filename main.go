package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	endpointURL = "https://jsonplaceholder.typicode.com/posts"
	interval    = 15 * time.Second
)

type Data struct {
	Water  int `json:"water"`
	Wind   int `json:"wind"`
	Status struct {
		Water string `json:"water"`
		Wind  string `json:"wind"`
	} `json:"status"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		waterStatus := getStatus(water)
		windStatus := getStatus(wind)

		data := Data{
			Water: water,
			Wind:  wind,
			Status: struct {
				Water string `json:"water"`
				Wind  string `json:"wind"`
			}{
				Water: waterStatus,
				Wind:  windStatus,
			},
		}

		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println(err)
			continue
		}

		resp, err := http.Post(endpointURL, "application/json", strings.NewReader(string(jsonData)))
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("POST response: %v\n", resp.Status)
		fmt.Printf("%s", jsonData)

		time.Sleep(interval)
	}
}

func getStatus(value int) string {
	if value < 5 {
		return "aman"
	} else if value >= 5 && value <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
