package utils

import (
	"unicode/utf8"
)

/*
用来检查字符串相似度的函数
*/

func calculateLongestWord(a, b string) int {
	if len(a) >= len(b) {
		return len(a)
	}

	return len(b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func computeLevenshteinValue(a, b string) int {
	f := make([]int, utf8.RuneCountInString(b)+1)

	for j := range f {
		f[j] = j
	}

	for _, ca := range a {
		j := 1
		fj1 := f[0]
		f[0]++
		for _, cb := range b {
			mn := min(f[j]+1, f[j-1]+1)
			if cb != ca {
				mn = min(mn, fj1+1)
			} else {
				mn = min(mn, fj1)
			}

			fj1, f[j] = f[j], mn
			j++
		}
	}

	return f[len(f)-1]
}

func ComputeLevenshteinPercentage(a, b string) float64 {
	distance := computeLevenshteinValue(a, b)
	length := calculateLongestWord(a, b)

	percentFloat := 1.00 - float64(distance)/float64(length)

	return percentFloat
}
