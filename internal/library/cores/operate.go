package cores

import (
	"encoding/json"
	"io/ioutil"
	"sort"

	"Caesar/internal/relation"
	"Caesar/pkg/utils"
)

func ReadDict(info []string, dicts []map[string]string) []relation.TagPath {
	/*
		用来读取目录字典的数据，转换成列表的形式
	*/

	//// 错误回收
	//defer func() {
	//	if recover() != nil {
	//		println(recover())
	//	}
	//}()

	var allJson []relation.TagPath

	for _, v := range info {
		var eachJson []relation.EachPath

		dictPath := utils.GetValueFromList(v, dicts)
		bytes, err := ioutil.ReadFile(dictPath)
		if err != nil {
			println(dictPath + " open failed")
			//panic(dictPath + " open failed")
			continue
		}

		if err := json.Unmarshal(bytes, &eachJson); err != nil {
			println(" Unmarshal failed")
			continue
		}

		for _, y := range eachJson {
			mid := relation.TagPath{
				EachPath: y,
				Tag:      v,
			}
			allJson = append(allJson, mid)
		}

	}

	// 将每个json数据按照Hits进行排序
	sort.Slice(allJson,
		func(i, j int) bool {
			return allJson[i].Hits > allJson[j].Hits
		})

	return allJson

}
