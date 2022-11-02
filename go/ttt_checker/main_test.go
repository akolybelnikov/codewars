package ttt_checker_test

import (
	c "github.com/akolybelnikov/codewars/go/ttt_checker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Checker", func() {
	testBoards := []struct {
		name   string
		board  [3][3]int
		result int
	}{
		{
			"not yet finished",
			[3][3]int{
				{0, 0, 1},
				{0, 1, 2},
				{2, 1, 0},
			},
			-1,
		},
		{
			"X wins row",
			[3][3]int{
				{1, 1, 1},
				{0, 2, 2},
				{0, 0, 0},
			},
			1,
		},
		{
			"O wins diagonal",
			[3][3]int{
				{1, 1, 2},
				{0, 2, 1},
				{2, 0, 0},
			},
			2,
		},
		{
			"X wins column",
			[3][3]int{
				{2, 1, 2},
				{2, 1, 1},
				{1, 1, 2},
			},
			1,
		},
		{
			"O wins column",
			[3][3]int{
				{2, 1, 2},
				{2, 0, 1},
				{2, 1, 2},
			},
			2,
		},
		{
			"draw",
			[3][3]int{
				{2, 1, 2},
				{2, 1, 1},
				{1, 2, 1},
			},
			0,
		},
		{
			"draw 2",
			[3][3]int{
				{0, 1, 0},
				{2, 1, 0},
				{1, 2, 2},
			},
			-1,
		},
	}
	for _, tb := range testBoards {
		tb := tb
		It(tb.name, func() {
			Expect(c.IsSolved(tb.board)).To(Equal(tb.result))
		})
	}
})
