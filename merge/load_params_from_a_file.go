package merge

import (
	"encoding/json"
	"io/ioutil"
)

func loadParamsFromAFile(f string) ([]Param, error) {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var p []Param
	err = json.Unmarshal(raw, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
