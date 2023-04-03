package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	min := 1
	max := 100

	for range time.Tick(time.Minute * 15) {

		data := map[string]interface{}{
			"water": rand.Intn(max-min) + min,
			"wind":  rand.Intn(max-min) + min,
		}

		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts",
			bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var jsonData map[string]interface{}

		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			log.Println(err)
		}

		valueWater := jsonData["water"].(float64)

		valueWind := jsonData["wind"].(float64)

		statusWater := ""
		statusWind := ""

		if valueWater < 5 {
			statusWater = "aman"
		} else if valueWater >= 6 && valueWater <= 8 {
			statusWater = "siaga"
		} else if valueWater > 8 {
			statusWater = "bahaya"
		}

		if valueWind >= 7 && valueWind <= 15 {
			statusWind = "siaga"
		} else if valueWind > 15 {
			statusWind = "bahaya"
		} else {
			statusWind = "aman"
		}
		log.Println(string(body))
		log.Printf("water : %v meter per detik\n", valueWater)
		log.Printf("wind : %v meter\n", valueWind)
		log.Printf("status water : %v\n", statusWater)
		log.Printf("status wind : %v\n", statusWind)
	}
}
