package merge

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Merge", func() {
	Context("loadParamsFromAFile", func() {
		It("should load parameters from file", func() {
			expected := []Param{
				Param{ParameterKey: "Parameter1", ParameterValue: "parameter 1 value"},
				Param{ParameterKey: "Parameter3", ParameterValue: "43"},
			}
			result, err := loadParamsFromAFile("../test/params1.json")
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expected))
		})
	})
	Context("thisParameterIsInFiles", func() {
		var sampleData []Param
		var p Param
		var pEmpty Param
		BeforeEach(func() {
			p = Param{ParameterKey: "Parameter1", ParameterValue: "value1", UsePreviousValue: false, ResolvedValue: ""}
			sampleData = append(sampleData, p)
		})
		It("shoud return true if a parameter key is present has a value", func() {
			val, ok := thisParameterIsInFiles(sampleData, "Parameter1")
			Expect(ok).To(Equal(true))
			Expect(val).To(Equal(p))
		})
		It("shoud return false if a parameter key is not present", func() {
			val, ok := thisParameterIsInFiles(sampleData, "Parameter2")
			Expect(ok).To(Equal(false))
			Expect(val).To(Equal(pEmpty))
		})
	})
	Context("getParamsFromTemplate", func() {
		It("shall get parameters from a CloudFormation template", func() {
			p, err := getParamsFromTemplate("../test/template.json")
			Expect(err).To(BeNil())
			Expect(p["Parameter1"].(map[string]interface{})["Type"].(string)).To(Equal("String"))
			Expect(p["Parameter2"].(map[string]interface{})["Default"].(string)).To(Equal("string-value"))
			Expect(p["Parameter3"].(map[string]interface{})["Default"].(float64)).To(Equal(float64(42)))
		})
	})
})
