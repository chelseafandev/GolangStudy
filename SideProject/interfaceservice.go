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
	XMLName httpconnection  `xml:"options"`
	Http    httpconnection  `xml:"http"`
	Https   httpsconnection `xml:"https"`
}

type httpconnection struct {
	Ip       string `xml:"ip"`
	Port     int    `xml:"port"`
	Location string `xml:"location"`
}

type httpsconnection struct {
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

	fmt.Println("options.Http.Ip:", options.Http.Ip)
	fmt.Println("options.Http.Port:", options.Http.Port)
	fmt.Println("options.Http.Location:", options.Http.Location)
	fmt.Println()

	fmt.Println("options.Https.Ip:", options.Https.Ip)
	fmt.Println("options.Https.Port:", options.Https.Port)
	fmt.Println("options.Https.Location:", options.Https.Location)
	fmt.Println("options.Https.Certpath:", options.Https.Certpath)
	fmt.Println("options.Https.Keypath:", options.Https.Keypath)
	fmt.Println()
}

func main() {
	fmt.Println("read config.xml!")
	readOptions("config.xml")

	fmt.Println("start https server!")
	https_server.StartServer(options.Https.Ip, options.Https.Port, options.Https.Location, options.Https.Certpath, options.Https.Keypath)
}
