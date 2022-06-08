package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAPILogin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Login Suite")
}
 