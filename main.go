package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func processRoot(w http.ResponseWriter, req *http.Request) {
	log.Print("/ request received")
}

func processSleep(w http.ResponseWriter, req *http.Request) {
	s, err := strconv.Atoi(req.URL.Query().Get("ms"))
	if err != nil || s < 1 {
		s = 0
		log.Print(fmt.Sprintf("/sleep request received", s))
	} else {
		log.Print(fmt.Sprintf("/sleep?s=%d request received", s))
	}

	fmt.Fprintf(w, "delayed for %ds", s)
	time.Sleep(time.Millisecond * time.Duration(s) )
	log.Print("end")
}

func main() {
	http.HandleFunc("/sleep", processSleep)
	http.HandleFunc("/", processRoot)

	log.Print("server start listening for :8080")
	http.ListenAndServe(":8080", nil)
}
