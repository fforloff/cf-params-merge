package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/awslabs/goformation"
	"github.com/spf13/cobra"
)

var (
	cfTemplate         string
	paramFilesArray    []string
	initFileArrayValue []string
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge parameters",
	Long:  `Merge Cloudformation parameters from parameters files and environment variables`,

	Run: mergeRun,
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	mergeCmd.PersistentFlags().StringVarP(&cfTemplate, "template", "t", "", "CloudFormation template (Required)")
	mergeCmd.MarkPersistentFlagRequired("template")
	mergeCmd.PersistentFlags().StringArrayVarP(&paramFilesArray, "param-file", "p", initFileArrayValue, "CloudFormation parameter file")

}

type param struct {
	ParameterKey     string `json:"ParameterKey"`
	ParameterValue   string `json:"ParameterValue,omitempty"`
	UsePreviousValue bool   `json:"UsePreviousValue,omitempty"`
	ResolvedValue    string `json:"ResolvedValue,omitempty"`
}

func thisParameterKeyHasValue(p []param, name string) (string, bool) {
	for _, v := range p {
		if v.ParameterKey == name {
			return v.ParameterValue, true
		}
	}
	return "", false
}

func getParamsFromTemplate(f string) (map[string]interface{}, error) {
	t, err := goformation.Open(string(f))
	if err != nil {
		return nil, err
	}
	return t.Parameters, nil
}

func getTemplateParamValue(i interface{}) string {
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

func getParamsFromFile(f string) ([]param, error) {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var p []param
	err = json.Unmarshal(raw, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getParamsFromFiles(fs []string) ([][]param, error) {
	var res [][]param
	reverseSlice(fs)
	for _, f := range fs {
		pp, err := getParamsFromFile(f)
		if err != nil {
			return res, err
		}
		res = append(res, pp)
	}
	return res, nil
}

func getParamValueFromFiles(pp [][]param, n string) string {
	var (
		v  string
		ok bool
	)
	for _, p := range pp {
		if v, ok = thisParameterKeyHasValue(p, n); ok {
			return v
		}
	}
	return ""
}

func getParamValueFromEnv(n string) string {
	var (
		v  string
		ok bool
	)
	if v, ok = os.LookupEnv(n); ok {
		return v
	}
	return ""
}

func mergeRun(cmd *cobra.Command, args []string) {

	fmt.Println(paramFilesArray)
	fmt.Println(cfTemplate)
	var (
		res                []param
		paramValue         string
		pv                 string
		err                error
		paramsFromTemplate map[string]interface{}
		paramsFromFiles    [][]param
	)

	paramsFromTemplate, err = getParamsFromTemplate(string(cfTemplate))
	if err != nil {
		panic(err)
	}
	if len(paramFilesArray) > 0 {
		paramsFromFiles, err = getParamsFromFiles(paramFilesArray)
		if err != nil {
			panic(err)
		}
	}

	for name, valueInterface := range paramsFromTemplate {

		pv = getTemplateParamValue(valueInterface)
		if len(pv) > 0 {
			paramValue = pv
		}

		pv = getParamValueFromFiles(paramsFromFiles, name)
		if len(pv) > 0 {
			paramValue = pv
		}

		pv = getParamValueFromEnv(name)
		if len(pv) > 0 {
			paramValue = pv
		}

		if len(paramValue) > 0 {
			param := param{
				ParameterKey:   name,
				ParameterValue: paramValue,
			}
			res = append(res, param)
		}
	}
	resj, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}

	//fmt.Fprintf("%s", resj)
	fmt.Printf("%s", resj)
}
