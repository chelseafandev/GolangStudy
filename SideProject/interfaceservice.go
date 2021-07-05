package main

import (
	https_server "GolangStudy/SideProject/https"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Options struct {
	XMLName       tlsconnection `xml:"options"`
	Tlsconnection tlsconnection `xml:"tlsconnection"`
}

type tlsconnection struct {
	Ip       string `xml:"ip"`
	Port     int    `xml:"port"`
	Location string `xml:"location"`
	Certpath string `xml:"certpath"`
	Keypath  string `xml:"keypath"`
}

var options Options

func readOptions(configpath string) {
	xmlFile, err := os.Open(configpath)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, err_readAll := ioutil.ReadAll(xmlFile)
	if err_readAll != nil {
		log.Fatal(err_readAll)

	}

	err_Unmarshal := xml.Unmarshal(byteValue, &options)
	if err_Unmarshal != nil {
		log.Fatal(err_Unmarshal)
	}
}

func main() {
	fmt.Println("start https server!")
	readOptions("config.xml")
	https_server.StartServer(options.Tlsconnection.Ip, options.Tlsconnection.Port, options.Tlsconnection.Location, options.Tlsconnection.Certpath, options.Tlsconnection.Keypath)
}
