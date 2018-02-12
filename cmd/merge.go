package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

type Param struct {
	ParameterKey   string `json:"ParameterKey"`
	ParameterValue string `json:"ParameterValue"`
}

func getParamsFromFile(f string) ([]Param, error) {
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

func thisParameterKeyHasValue(p []Param, param string) (string, bool) {
	for _, v := range p {
		if v.ParameterKey == param {
			return v.ParameterValue, true
		}
	}
	return "", false
}

func mergeRun(cmd *cobra.Command, args []string) {
	// fmt.Println(cfTemplate)
	// fmt.Println(paramFilesArray)

	var res []Param

	// Open a template from file (can be JSON or YAML)
	template, err := goformation.Open(string(cfTemplate))
	if err != nil {
		panic(err)
	}

	for param_file := range paramFilesArray {
		paramsFromFile, err := getParamsFromFile(string(param_file))
		if err != nil {
			panic(err)
		}
		log.Println(paramsFromFile)

		for name, properties := range template.Parameters {
			var value string
			p, _ := properties.(map[string]interface{})

			if default_val_string, ok := p["Default"].(string); ok {
				value = default_val_string
			}
			if file_val_string, ok := thisParameterKeyHasValue(paramsFromFile, name); ok {
				value = file_val_string
			}
			if env_val_string, ok := os.LookupEnv(name); ok {
				value = env_val_string
			}
			if len(value) > 0 {
				param := Param{
					ParameterKey:   name,
					ParameterValue: value,
				}
				res = append(res, param)
			}
		}
	}

	resj, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}

	//fmt.Fprintf("%s", resj)
	fmt.Printf("%s", resj)
}
