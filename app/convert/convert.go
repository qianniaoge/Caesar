package convert

import (
	"io/ioutil"
	"strings"

	"Caesar/internal/relation"
	"Caesar/pkg/utils"
)

func TextToJsonOfPath(path string) {

	txtList := utils.WrapFuncGetFile(utils.GetFileFromDocument, path, "txt")

	for _, v := range txtList {

		var newSlice []relation.EachPath

		infos, err := utils.ReadLines(v)
		if err != nil {
			return
		}

		if len(infos) == 0 {
			return
		}

		for _, v := range utils.RemoveDuplicateElement(infos) {
			newTag := relation.EachPath{
				Hits: 0,
				Path: v,
			}
			newSlice = append(newSlice, newTag)

		}

		info, _ := utils.CustomMarshal(newSlice)
		_ = ioutil.WriteFile(strings.TrimSuffix(v, ".txt")+".json", []byte(info), 0644)

	}

}

func TextToJsonOfFile(fileName string) {

	var newSlice []relation.EachPath

	infos, err := utils.ReadLines(fileName)
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range utils.RemoveDuplicateElement(infos) {
		newTag := relation.EachPath{
			Hits: 0,
			Path: v,
		}
		newSlice = append(newSlice, newTag)

	}

	info, _ := utils.CustomMarshal(newSlice)
	_ = ioutil.WriteFile(strings.TrimSuffix(fileName, ".txt")+".json", []byte(info), 0644)

}
