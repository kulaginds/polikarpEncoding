package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		inputSeq     []int
		expOuputSeqs [][]int
	}{
		{
			inputSeq: []int{3, 2, 1, 1, 7, 2, 4, -1, 3, -1, 4, -1},
			expOuputSeqs: [][]int{
				{3, 1, 4},
				{2, 7},
				{1, 2, 3, 4},
			},
		},
		{
			inputSeq: []int{2, -1, 2, -1, 3, -1},
			expOuputSeqs: [][]int{
				{2},
				{},
				{2, 3},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			actualOutputSeqs := Decode(tc.inputSeq)

			a.Equal(tc.expOuputSeqs, actualOutputSeqs)
		})
	}
}
