package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/exec"
)

var (
	sources []string
)

func init() {
	flag.Parse()
	sources = flag.Args()
}

func GetProbes() {
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

func GetProbe(filename string) (*ProbeData, error) {
	// ffprobe R0000000_er.MP4 -show_format -print_format json
	cmd := exec.Command("ffprobe", filename, "-show_format", "-print_format", "json")
	cmd.Stderr = os.Stderr

	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	var v ProbeData
	err = json.NewDecoder(r).Decode(&v)
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return &v, nil
}
