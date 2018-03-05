package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfmingle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cfmingle Suite")
}
