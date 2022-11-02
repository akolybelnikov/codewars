package ttt_checker_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTicTacToeChecker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TicTacToeChecker Suite")
}
