package reset

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"Caesar/internal/relation"
	"Caesar/pkg/utils"
)

func SetupHitsOfZeroInFile(fileName string) {

	var eachJson []relation.EachPath
	var newJson []relation.EachPath

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		println(fileName + " open failed")
		return
	}

	if err := json.Unmarshal(bytes, &eachJson); err != nil {
		println("json Unmarshal is failed")
		return
	}

	for _, v := range eachJson {
		v.Hits = 0
		newJson = append(newJson, v)
	}

	// 最后面4个空格，让json格式更美观
	result, _ := json.MarshalIndent(newJson, "", "    ")

	_ = ioutil.WriteFile(fileName, result, 0644)

}

func SetupHitsOfZeroInDocument(document string) {

	assetsDir := document

	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(assetsDir)
	if err != nil {
		println("Can not open dir")
		return
	}

	for i := range fileInfoList {
		var dictPath string
		if !strings.HasSuffix(dictPath, "/") {
			dictPath = assetsDir + "/"
		} else {
			dictPath = assetsDir
		}

		fileName := dictPath + fileInfoList[i].Name()

		var eachJson []relation.EachPath
		var newJson []relation.EachPath

		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			println(fileName + " open failed")
			continue
		}

		if err := json.Unmarshal(bytes, &eachJson); err != nil {
			println("json Unmarshal is failed")
			continue
		}

		for _, v := range eachJson {
			v.Hits = 0
			newJson = append(newJson, v)
		}

		// 最后面4个空格，让json格式更美观
		result, _ := utils.CustomMarshal(newJson)

		_ = ioutil.WriteFile(fileName, []byte(result), 0644)

	}

}
