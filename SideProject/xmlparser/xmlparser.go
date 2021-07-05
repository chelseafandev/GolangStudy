package xmlparser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Options struct {
	Http  httpconnection  `xml:"http"`
	Https httpsconnection `xml:"https"`
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

	//output, _ := xml.Marshal(opts)
	//fmt.Println(string(output))

	//fmt.Println(opts)

	return nil
}
