package merge

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/awslabs/goformation"
)

//Param struct definition
type Param struct {
	ParameterKey     string `json:"ParameterKey"`
	ParameterValue   string `json:"ParameterValue,omitempty"`
	UsePreviousValue bool   `json:"UsePreviousValue,omitempty"`
	ResolvedValue    string `json:"ResolvedValue,omitempty"`
}

//ThisParameterKeyHasValue ...
func ThisParameterKeyHasValue(p []Param, name string) (string, bool) {
	for _, v := range p {
		if v.ParameterKey == name {
			return v.ParameterValue, true
		}
	}
	return "", false
}

//GetParamsFromTemplate ...
func GetParamsFromTemplate(f string) (map[string]interface{}, error) {
	t, err := goformation.Open(string(f))
	if err != nil {
		return nil, err
	}
	return t.Parameters, nil
}

// GetTemplateParamValue ...
func GetTemplateParamValue(i interface{}) string {
	dv := ""
	if vv, ok := i.(map[string]interface{}); ok {
		if v, ok := vv["Default"]; ok {
			switch vTyped := v.(type) {
			case string:
				dv = vTyped
			case int:
				dv = strconv.Itoa(vTyped)
			case int64:
				dv = strconv.FormatInt(vTyped, 10)
			case float64:
				dv = strconv.FormatFloat(vTyped, 'f', -1, 64)
			case bool:
				if vTyped {
					dv = "true"
				} else {
					dv = "false"
				}
			default:
				// need better error handling here
				os.Exit(1)
			}
		}
	}
	return dv
}

//GetParamsFromFile ...
func GetParamsFromFile(f string) ([]Param, error) {
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

//ReverseSlice ...
func ReverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//GetParamsFromFiles ...
func GetParamsFromFiles(fs []string) ([][]Param, error) {
	var res [][]Param
	ReverseSlice(fs)
	for _, f := range fs {
		pp, err := GetParamsFromFile(f)
		if err != nil {
			return res, err
		}
		res = append(res, pp)
	}
	return res, nil
}

//GetParamValueFromFiles ...
func GetParamValueFromFiles(pp [][]Param, n string) string {
	var (
		v  string
		ok bool
	)
	for _, p := range pp {
		if v, ok = ThisParameterKeyHasValue(p, n); ok {
			return v
		}
	}
	return ""
}

//GetParamValueFromEnv ...
func GetParamValueFromEnv(n string) string {
	var (
		v  string
		ok bool
	)
	if v, ok = os.LookupEnv(n); ok {
		return v
	}
	return ""
}
