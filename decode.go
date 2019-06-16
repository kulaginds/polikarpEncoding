package main

import "math"

func Decode(inputSeq []int) [][]int {
	var i, seqIndex int
	seqsCount := getSeqsCount(inputSeq)
	actualSeqsIndexes := make([]int, seqsCount)
	endedSeqs := make([]int, 0, seqsCount)
	outputSeqs := make([][]int, seqsCount)

	for i = 0; i < seqsCount; i++ {
		outputSeqs[i] = make([]int, 0, len(inputSeq))
		actualSeqsIndexes[i] = i
	}

	i = 0

	for {
		for _, seqIndex = range actualSeqsIndexes {
			if inputSeq[i] == -1 {
				endedSeqs = append(endedSeqs, seqIndex)
			} else {
				outputSeqs[seqIndex] = append(outputSeqs[seqIndex], inputSeq[i])
			}

			i++
		}

		if len(endedSeqs) > 0 {
			actualSeqsIndexes = diff(actualSeqsIndexes, endedSeqs)
			endedSeqs = endedSeqs[:0]

			if len(actualSeqsIndexes) == 0 {
				break
			}
		}
	}

	return outputSeqs
}

func getSeqsCount(inputSeq []int) int {
	count := 0

	for _, i := range inputSeq {
		if i == -1 {
			count++
		}
	}

	return count
}

func diff(seq1, seq2 []int) []int {
	newSeq := make([]int, 0, max(len(seq1), len(seq2)))

	var has bool
	for _, v1 := range seq1 {
		for _, v2 := range seq2 {
			if v1 == v2 {
				has = true
			}
		}

		if !has {
			newSeq = append(newSeq, v1)
		} else {
			has = false
		}
	}

	return newSeq
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
