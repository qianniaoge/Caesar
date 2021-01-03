package units

import (
	"Caesar/internal/relation"
	"Caesar/pkg/utils"
	"testing"
)

/*
此测试方法可以讲普通路径文件转换成caesar能识别的字典
*/

func TestChange(t *testing.T) {

	var newSlice []relation.EachPath

	infos, err := utils.ReadLines("info.txt")
	if err != nil {
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

	println(info)
}
