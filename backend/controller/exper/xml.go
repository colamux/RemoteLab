package exper

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func Get(item *interface{}, path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		}
	}
}
