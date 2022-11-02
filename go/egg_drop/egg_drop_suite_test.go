package egg_drop_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEggDrop(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EggDrop Suite")
}
