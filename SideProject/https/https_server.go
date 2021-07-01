package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RequestData struct {
	url       string
	user_id   string
	user_guid string
}

type Message struct {
	Name, Text string
}

func handleRequest(w http.ResponseWriter, req *http.Request) {

	log.Println("Method: " + req.Method)
	log.Println("Host: " + req.Host)
	log.Println("URL: " + req.URL.String())

	content_length := req.ContentLength
	body := make([]byte, content_length)
	req.Body.Read(body)

	log.Println("Body: " + string(body))

	dec := json.NewDecoder(strings.NewReader(string(body)))
	// read open bracket
	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	server_ip := "127.0.0.1"
	server_port := "18443"
	var conn_info string
	conn_info = server_ip + ":" + server_port

	http.HandleFunc("/test", handleRequest)
	err := http.ListenAndServeTLS(conn_info, "certs/server.crt", "certs/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
