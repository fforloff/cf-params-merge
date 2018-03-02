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

	res, err := merge.MergeParams(cfTemplate, paramFilesArray)

	resj, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}

	//fmt.Fprintf("%s", resj)
	fmt.Printf("%s", resj)
}
