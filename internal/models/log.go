package models

type Log struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}
