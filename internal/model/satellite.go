package model

import (
	"bytes"
	"encoding/json"
)

type Satellite struct {
	Name string `json:"name"`
}

func (s *Satellite) Encode() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Satellite) Decode(in []byte) error {
	var res Satellite
	if err := json.NewDecoder(bytes.NewReader(in)).Decode(&res); err != nil {
		return err
	}
	*s = res
	return nil
}
