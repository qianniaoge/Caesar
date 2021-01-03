package units

import (
	"Caesar/pkg/utils"
	"testing"
)

func TestSimilar(T *testing.T) {
	ratio := utils.ComputeLevenshteinPercentage("index.djangophp", "indexluffy.php")
	println(ratio)
}
