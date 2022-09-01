package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Name         string       `json:"name"`
	Healthy      bool         `json:"healthy"`
	Dependencies []Dependency `json:"dependencies"`
	//Others			string			`json:"others"`
}

type Dependency struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Healthy bool   `json:"healthy"`
}

func statusCode() (statCode int) {
	requestURL := fmt.Sprintf("https://accelbyte.io/ghost/api/v4/admin/site/")
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	statCode = res.StatusCode

	return
}

func healthz(w http.ResponseWriter, req *http.Request) {
	statCode := statusCode()

	if statCode >= 400 {
		data := &Response{
			Name:    "wen-app",
			Healthy: false,
			Dependencies: []Dependency{
				{
					Name:    "redis-db",
					Url:     "ab.io",
					Healthy: false,
				},
				{
					Name:    "postgre-db",
					Url:     "ab.io",
					Healthy: false,
				},
			},
		}
		content, _ := json.Marshal(data)
		jsonStr := string(content)
		fmt.Fprintf(w, jsonStr)

	} else {
		data := &Response{
			Name:    "wen-app",
			Healthy: true,
			Dependencies: []Dependency{
				{
					Name:    "redis-db",
					Url:     "ab.io",
					Healthy: true,
				},
				{
					Name:    "postgre-db",
					Url:     "ab.io",
					Healthy: true,
				},
			},
		}
		content, _ := json.Marshal(data)
		jsonStr := string(content)
		fmt.Fprintf(w, jsonStr)
	}
}

func main() {

	http.HandleFunc("/healthz", healthz)

	http.ListenAndServe(":8118", nil)
}
