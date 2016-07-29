package main

import (
	"time"
)

const (
	LayoutCreationTime = "2006-01-02 15:04:05"
)

type ProbeData struct {
	Format *ProbeFormat `json:"format,omitempty"`
}

type ProbeFormat struct {
	//Filename         string            `json:"filename"`
	//NBStreams        int               `json:"nb_streams"`
	//NBPrograms       int               `json:"nb_programs"`
	//FormatName       string            `json:"format_name"`
	//FormatLongName   string            `json:"format_long_name"`
	//StartTimeSeconds float64           `json:"start_time,string"`
	DurationSeconds float64 `json:"duration,string"`
	//Size             uint64            `json:"size,string"`
	//BitRate          uint64            `json:"bit_rate,string"`
	//ProbeScore       float64           `json:"probe_score"`
	Tags *ProbeFormatTags `json:"tags"`
}

type ProbeFormatTags struct {
	CreationTime string `json:"creation_time"`
}

//func (f ProbeFormat) StartTime() time.Duration {
//	return time.Duration(f.StartTimeSeconds * float64(time.Second))
//}

func (f ProbeFormat) Duration() time.Duration {
	return time.Duration(f.DurationSeconds * float64(time.Second))
}

func (f ProbeFormat) CreationTime() (time.Time, error) {
	t, err := time.Parse(LayoutCreationTime, f.Tags.CreationTime)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (f ProbeFormat) Date() (time.Time, error) {
	t, err := f.CreationTime()
	if err != nil {
		return t, err
	}
	return t.Add(-f.Duration()), nil
}
