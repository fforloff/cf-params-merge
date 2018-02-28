package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/fforloff/cfmingle/merge"
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

func mergeRun(cmd *cobra.Command, args []string) {

	fmt.Println(paramFilesArray)
	fmt.Println(cfTemplate)
	var (
		res                []merge.Param
		paramValue         string
		pv                 string
		err                error
		paramsFromTemplate map[string]interface{}
		paramsFromFiles    [][]merge.Param
	)

	paramsFromTemplate, err = merge.GetParamsFromTemplate(string(cfTemplate))
	if err != nil {
		panic(err)
	}
	if len(paramFilesArray) > 0 {
		paramsFromFiles, err = merge.GetParamsFromFiles(paramFilesArray)
		if err != nil {
			panic(err)
		}
	}

	for name, valueInterface := range paramsFromTemplate {

		pv = merge.GetTemplateParamValue(valueInterface)
		if len(pv) > 0 {
			paramValue = pv
		}

		pv = merge.GetParamValueFromFiles(paramsFromFiles, name)
		if len(pv) > 0 {
			paramValue = pv
		}

		pv = merge.GetParamValueFromEnv(name)
		if len(pv) > 0 {
			paramValue = pv
		}

		if len(paramValue) > 0 {
			param := merge.Param{
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
