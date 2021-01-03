package units

import (
	"Caesar/pkg/utils"
	"testing"
)

func TestDocument(T *testing.T) {
	info := utils.WrapFuncGetFile(utils.GetFileFromDocument, "/Users/null/Desktop/denmo", "txt")

	for _, v := range info {

		println(v)
	}
}
