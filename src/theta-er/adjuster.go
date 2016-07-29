package main

import (
	"flag"
	"log"
	"os"
)

var (
	sources []string
)

func init() {
	flag.Parse()
	sources = flag.Args()
}

func AdjustMetaData() {
	for _, source := range sources {
		fileInfo, err := os.Stat(source)
		if os.IsNotExist(err) {
			log.Println("File is not exist:", source)
			continue
		}
		if err != nil {
			log.Println("Error:", source, err)
			continue
		}
		if fileInfo.IsDir() {
			log.Println("It is directory:", source)
			continue
		}
		probeData, err := GetProbe(source)
		if err != nil {
			log.Println("Error:", source, err)
			continue
		}
		date, err := probeData.Format.Date()
		if err != nil {
			log.Println("Error:", source, err)
			continue
		}
		log.Println(source, date)
	}
}
