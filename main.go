package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Starting HTTP Server on port 8080")
	http.HandleFunc("/person/", handleUrlPerson)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleUrlPerson(writer http.ResponseWriter, request *http.Request) {
	name, found := getUrlParam(request.URL, "name")
	if found {
		n, err := fmt.Fprintf(writer, "Hello %v", name)
		if n <= 0 {
			log.Fatal("Unable to write to responseWriter: ", err)
		}
	} else {
		writer.WriteHeader(http.StatusAccepted)
		n, err := fmt.Fprint(writer, "Hello Anonymous Person")
		if n <= 0 {
			log.Fatal("Unable to write to responseWriter: ", err)
		}
	}
}

func getUrlParam(theUrl *url.URL, param string) (string, bool) {
	urlMap, _ := url.ParseQuery(theUrl.RawQuery)
	if urlMap.Has(param) {
		return urlMap.Get(param), true
	}

	return "", false
}
