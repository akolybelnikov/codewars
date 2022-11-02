package egg_drop_test

import (
	e "github.com/akolybelnikov/codewars/go/egg_drop"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/big"
)

func testerIntString(a, b int64, c string) {
	Expect(e.Height(big.NewInt(a), big.NewInt(b)).String()).To(Equal(c))
}
func testerInt(a, b, c int64) {
	Expect(e.Height(big.NewInt(a), big.NewInt(b)).String()).To(Equal(big.NewInt(c).String()))
}

var _ = Describe("Should work", func() {
	It("should work for some basic tests", func() {
		testerInt(1, 3, 3)
		testerInt(1, 51, 51)
		testerInt(2, 1, 1)
		testerInt(2, 20, 210)
		testerInt(3, 20, 1350)
		testerInt(4, 20, 6195)
	})
	It("should work for some advanced cases", func() {
		testerIntString(30, 400, "1549835246319296875396604627597789530400430471")
		testerIntString(310, 512, "13407798997994793023909825703579562145240298498207798340676479471840456202906831071791573160709230447579657844888754638680952175985107213464211003140484530")
		testerIntString(494, 500, "3273390607896141870013189696827599152216642046043064789483291368096133796404674554883270092325904157150886684127560071009217256545885393053070689036899")
	})
})
