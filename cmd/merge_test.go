package cmd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Merge", func() {
	Context("getParamsFromFile", func() {
		It("should load parameters from file", func() {
			expected := []param{
				param{ParameterKey: "StackName", ParameterValue: "blahStack"},
				param{ParameterKey: "Environment", ParameterValue: "dev"},
			}
			result, err := getParamsFromFile("../test/params1.json")
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expected))
		})
	})
	Context("thisParameterKeyHasValue", func() {
		var sampleData []param
		BeforeEach(func() {
			p := param{ParameterKey: "Parameter1", ParameterValue: "value1"}
			sampleData = append(sampleData, p)
		})
		It("shoud return true if a parameter key is present has a value", func() {
			val, ok := thisParameterKeyHasValue(sampleData, "Parameter1")
			Expect(ok).To(Equal(true))
			Expect(val).To(Equal("value1"))
		})
		It("shoud return false if a parameter key is not present", func() {
			val, ok := thisParameterKeyHasValue(sampleData, "Parameter2")
			Expect(ok).To(Equal(false))
			Expect(val).To(Equal(""))
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
	Context("getTemplateParamValue", func() {
		It("shall return a value, if exists", func() {
			var tests = []struct {
				in  map[string]interface{}
				out string
			}{
				{map[string]interface{}{"Default": 42}, "42"},
				{map[string]interface{}{"Default": "a string"}, "a string"},
				{map[string]interface{}{"Default": 42.42}, "42.42"},
			}

			for _, tt := range tests {
				v, err := getTemplateParamValue(tt.in)
				Expect(err).To(BeNil())
				Expect(v).To(Equal(tt.out))
			}
		})
		It("shall return an empty string if a value does not exists,", func() {
			testVal := map[string]interface{}{"Somekey": 42}
			v, err := getTemplateParamValue(testVal)
			Expect(err).To(BeNil())
			Expect(v).To(Equal(""))
		})
	})
	Context("reverseSlice", func() {
		It("shall return a reversed slice of strings", func() {
			slice := []string{"one", "two", "three", "four"}
			expectedSlice := []string{"four", "three", "two", "one"}
			reverseSlice(slice)
			Expect(slice).To(Equal(expectedSlice))
		})
	})
})
