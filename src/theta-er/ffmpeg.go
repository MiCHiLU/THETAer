package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

func AddMetaData(filename string, values map[string]string) (os.FileInfo, error) {
	// ffprobe R0000000_er.MP4 -show_format -print_format json
	cmd := exec.Command("ffprobe", filename, "-hide_banner", "-show_format", "-print_format", "json")
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
