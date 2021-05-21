package main

import (
	"fmt"
	"net/http"
	"sync"
)

func cityHandler(cityData map[string]int, mtx *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var result string
		var statusCode int

		switch r.Method {
		case http.MethodGet:
			result, statusCode = get(cityData, r,mtx)
		case http.MethodPost:
			result, statusCode = post(cityData, r, mtx)
		default:
			result = "Invalid request method"
			statusCode = http.StatusMethodNotAllowed
		}

		if statusCode >= http.StatusBadRequest {
			http.Error(w, result, statusCode)
		} else {
			w.WriteHeader(statusCode)
			fmt.Fprint(w, result)
		}
	}
}

func get(data map[string]int, r *http.Request,mtx *sync.Mutex) (string, int) {

		var result string
		mtx.Lock()
		for city, freq := range data {
			if freq >=2 && freq <=4 || (freq > 21 && (freq % 10 >=2) && freq % 10 <=4 ){
				result=result+fmt.Sprintf("%v - %d раза\n", city, freq)
			} else {
				result=result+fmt.Sprintf("%v - %d раз\n", city, freq)
			}
		}
		mtx.Unlock()

	return result, http.StatusOK
}

func post(data map[string]int, r *http.Request,mtx *sync.Mutex) (string, int) {

	city := r.URL.Query().Get("name")

	if city == ""{
		return "Error", http.StatusBadRequest
	}

	mtx.Lock()
	//defer mtx.Unlock()
	data[city]++
	mtx.Unlock()

	return "POST Done", http.StatusCreated
}

func main() {
	var mtx sync.Mutex
	data := make(map[string]int)

	http.HandleFunc("/cities", cityHandler(data, &mtx))
	http.ListenAndServe(":8000", nil)
}
