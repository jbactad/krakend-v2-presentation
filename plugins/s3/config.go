package main

import "encoding/json"

type Config struct {
	Name   string `json:"name,omitempty"`
	Bucket string `json:"bucket,omitempty"`
}

func NewConfig(cfg map[string]interface{}) (c Config, err error) {
	j, err := json.Marshal(cfg)
	if err != nil {
		return
	}

	if err = json.Unmarshal(j, &c); err != nil {
		return
	}

	return c, nil
}
