package cmd

import (
	"fmt"

	"github.com/fforloff/cfmingle/merge"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MergeCmd", func() {
	Context("MergeParams", func() {
		It("shall load parameters from a template and parameters files and return a result", func() {
			cfTemplate := "../test/template.json"
			paramFilesArray := []string{"../test/params1.json", "../test/params2.json"}
			res, err := merge.MergeParams(cfTemplate, paramFilesArray)
			fmt.Println(res)
			Expect(err).Should(BeNil())
			Expect(res).ShouldNot(BeNil())

		})
	})
})
