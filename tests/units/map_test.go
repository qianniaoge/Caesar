package units

import (
	models "Caesar/internal/relation"
	"encoding/json"
	"io/ioutil"
	"testing"
)

/*
此单元测试可以将字典的hits重制为0, 慎用
*/
func TestMap(t *testing.T) {

	assetsDir := "../../assets/directory/"

	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(assetsDir)
	if err != nil {
		println("Can not open dir")
	}
	for i := range fileInfoList {
		fileName := assetsDir + fileInfoList[i].Name()

		var eachJson []models.EachPath
		var newJson []models.EachPath

		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			println(fileName + " open failed")
		}

		if err := json.Unmarshal(bytes, &eachJson); err != nil {
			println("json Unmarshal is failed")
		}

		for _, v := range eachJson {
			v.Hits = 0
			newJson = append(newJson, v)
		}

		// 最后面4个空格，让json格式更美观
		result, _ := json.MarshalIndent(newJson, "", "    ")

		_ = ioutil.WriteFile(fileName, result, 0644)

	}

}
