package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	ParameterKey   string `json:"ParameterKey"`
	ParameterValue string `json:"ParameterValue"`
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

func thisParameterKeyHasValue(p []param, param string) (string, bool) {
	for _, v := range p {
		if v.ParameterKey == param {
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

func mergeRun(cmd *cobra.Command, args []string) {
	// var res []param
	fmt.Println(paramFilesArray)
	fmt.Println(cfTemplate)

	// for paramFile := range paramFilesArray {
	// 	paramsFromFile, err := getParamsFromFile(string(paramFile))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(paramsFromFile)

	paramsFromTemplate, err := getParamsFromTemplate(string(cfTemplate))
	if err != nil {
		panic(err)
	}
	for p := range paramsFromTemplate {
		//var value string
		fmt.Println(p)
	}
	// 		// p, _ := properties.(map[string]interface{})

	// 		// if defaultValString, ok := p["Default"].(string); ok {
	// 		// 	value = defaultValString
	// 		// }
	// 		// if fileValString, ok := thisParameterKeyHasValue(paramsFromFile, name); ok {
	// 		// 	value = fileValString
	// 		// }
	// 		// if envValString, ok := os.LookupEnv(name); ok {
	// 		// 	value = envValString
	// 		// }
	// 		// if len(value) > 0 {
	// 		// 	param := param{
	// 		// 		ParameterKey:   name,
	// 		// 		ParameterValue: value,
	// 		// 	}
	// 		// 	res = append(res, param)
	// 		// }
	// 	}
	// }

	// resj, err := json.Marshal(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// //fmt.Fprintf("%s", resj)
	// fmt.Printf("%s", resj)
}
