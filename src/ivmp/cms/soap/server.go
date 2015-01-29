package server

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"github.com/ziutek/soap"
	"log"
	"net/http"
	"time"
)

func SoapHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	fmt.Println("header:", r.Header)
	br := bufio.NewReader(r.Body)
	b := make([]byte, 1000)
	n, err := br.Read(b)
	fmt.Println(string(b[:n]))
	fmt.Println(err)

	var ele soap.Element
	error := xml.Unmarshal(b, &ele)
	if error != nil {
		fmt.Println("Unmarshal error")
	}
	fmt.Println(ele)
	ip, _ := ele.Get("login")
	fmt.Println(ip)
}

func Start() {
	// http.HandleFunc("/", SoapHandle)
	// err := http.ListenAndServe(":59000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	s := &http.Server{
		Addr:         ":59000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	http.HandleFunc("/", SoapHandle)
	log.Fatal(s.ListenAndServe())
}
