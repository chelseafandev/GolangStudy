package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/chelseafandev/resetapi/internal/https"
	"github.com/chelseafandev/resetapi/internal/xmlparser"
)

func main() {
	var options xmlparser.Options
	fmt.Println("read config.xml!")
	err := xmlparser.ReadOptions("config.xml", &options)
	if err != nil {
		var roe *xmlparser.ReadOptionsError
		if errors.As(err, &roe) {
			fmt.Println("msg:", roe.Msg, "code:", roe.Status)
		} else {
			log.Println(err.Error())
		}
		return
	}

	fmt.Println("start https server!!")
	https.StartServer(options.Https.Ip, options.Https.Port, options.Https.Location, options.Https.Certpath, options.Https.Keypath)
}
