package exper

import (
	"encoding/xml"
)

type USB struct {
	XMLName xml.Name `xml:"hostdev"`
	Mode    string   `xml:"mode,attr"`
	Type    string   `xml:"type,attr"`
	Managed string   `xml:"managed,attr"`
	Source  struct {
		Vendor struct {
			Id string `xml:"id,attr"`
		} `xml:"vendor"`
		Product struct {
			Id string `xml:"id,attr"`
		} `xml:"product"`
	} `xml:"source"`
	Address struct {
		Type string `xml:"type,attr"`
		Bus  string `xml:"bus,attr"`
		Port string `xml:"port,attr"`
	} `xml:"address"`
}

func NewUSB() USB {
	return USB{
		XMLName: xml.Name{Local: "hostdev"},
		Mode:    "subsystem",
		Type:    "usb",
		Managed: "yes",
	}
}

func Attach() {
	// TODO:
}

func Detach() {

}
