package main

import (
	"fmt"
	"net/http"
	"sync"
)

func cityHandler(cityData map[string]int, mtx *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtx.Lock()
		defer mtx.Unlock()

		var result string
		var statusCode int

		switch r.Method {
		case http.MethodGet:
			result, statusCode = get(cityData, r)
		case http.MethodPost:
			result, statusCode = post(cityData, r)
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

func get(data map[string]int, r *http.Request) (string, int) {
// проверка наличия ключа

		//итерация по мапе
		/*for city, freq := range data {
			fmt.Fprintf(w, "%s %d times\n", city, freq)
		}*/
		var result string
		for city, freq := range data {
			result=result+fmt.Sprintf("%v %d times\n", city, freq)
		}
		//if city != "" {
		//	return fmt.Sprintf("%s %d times", city, data[city]), http.StatusOK
		//}


	// otherwise it prints all the data
	//result, err := json.Marshal(data)

	/*if err != nil {
		return err.Error(), http.StatusInternalServerError
	}*/

	return result, http.StatusOK
}

func post(data map[string]int, r *http.Request) (string, int) {

	city := r.URL.Query().Get("name")

	data[city]++
	fmt.Println(city)

	return "POST Done", http.StatusCreated
}

func main() {
	var mtx sync.Mutex
	data := make(map[string]int)

	http.HandleFunc("/cities", cityHandler(data, &mtx))
	http.ListenAndServe(":8000", nil)
}
