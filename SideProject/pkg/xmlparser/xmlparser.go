package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Options struct {
	Http  httpconnection  `xml:"http"`
	Https httpsconnection `xml:"https"`
}

type httpconnection struct {
	Session []session `xml:"session"`
}

type httpsconnection struct {
	Ip       string `xml:"ip"`
	Port     int    `xml:"port"`
	Location string `xml:"location"`
	Certpath string `xml:"certpath"`
	Keypath  string `xml:"keypath"`
	Timeout  string `xml:"timeout"`
}

type session struct {
	Ip       string `xml:"ip"`
	Port     int    `xml:"port"`
	Location string `xml:"location"`
	Timeout  string `xml:"timeout"`
}

type ReadOptionsError struct {
	Status int
	Msg    string
}

func (roe *ReadOptionsError) Error() string {
	return roe.Msg
}

func ReadOptions(configpath string, opts *Options) error {
	xmlFile, err := os.Open(configpath)
	if err != nil {
		msg := "xml open fail"
		return &ReadOptionsError{Status: 1, Msg: msg}
	}
	defer xmlFile.Close()

	byteValue, err_readAll := ioutil.ReadAll(xmlFile)
	if err_readAll != nil {
		msg := "xml read fail"
		return &ReadOptionsError{Status: 2, Msg: msg}
	}

	err_Unmarshal := xml.Unmarshal(byteValue, &opts)
	if err_Unmarshal != nil {
		msg := "xml unmarshal fail"
		return &ReadOptionsError{Status: 3, Msg: msg}
	}

	// 읽어들인 옵션 정보 출력
	result, err_Marshal := xml.MarshalIndent(opts, "", "\t")
	if err_Marshal != nil {
		msg := "xml marshal fail"
		return &ReadOptionsError{Status: 4, Msg: msg}
	}

	fmt.Printf("\n============================ ReadOptions ============================\n")
	fmt.Println(string(result))
	fmt.Printf("============================ ReadOptions ============================\n\n")

	return nil
}
