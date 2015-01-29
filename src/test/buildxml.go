package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName xml.Name `xml:"person"`
		Id      int      `xml:"id,attr"`
		Name    string
		Company string
		Email   []Email
		Address
	}
	p := Person{Id: 2, Name: "Jack", Company: "google.com"}
	p.Address = Address{City: "Qingdao", State: "China"}
	p.Email = append(p.Email, Email{"home", "gre@example.com"}, Email{"work", "gre@work.com"})

	output, err := xml.MarshalIndent(&p, "-", "**")
	if err != nil {
		fmt.Println("marshal error")
		return
	}
	fmt.Println(string(output))
}
