package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
                        <link>http</link>
		</Person>`
	type email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type address struct {
		City, State string
	}

	type person struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Company string
		Email   []email
		Group   []string `xml:"Group>Value"`
		address
		Link string `xml:"link"`
	}
	var v person
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Println("error:%w", err)
		return
	}
	fmt.Printf("%+v\n", v)

	fmt.Println("-----------------------------")
	b := strings.NewReader(data)
	decoder := xml.NewDecoder(b)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "Person" {
				var v2 person
				decoder.DecodeElement(&v2, &se)
				fmt.Printf("%+v\n", v2)
			}
		}
	}

	fmt.Println("-----------------------------")
	b2 := strings.NewReader(data)
	decoder2 := xml.NewDecoder(b2)
	var v3 person
	decoder2.Decode(&v3)
	fmt.Printf("%+v\n", v3)
}
