package report

import (
	"io/ioutil"
	"strconv"

	"Caesar/internal/relation"
	"Caesar/pkg/record"
	"Caesar/pkg/utils"
)

func ExportJson(results []relation.ResultPtah, savePath string) {

	/*
		@param ["http://127.0.0.1/index.php 200 示例"]
	*/

	var mapResults []map[string]string

	for _, v := range results {
		var result = make(map[string]string)
		result["path"] = v.Address
		result["code"] = strconv.Itoa(v.Code)
		result["title"] = v.Title
		result["length"] = strconv.Itoa(v.Length)
		mapResults = append(mapResults, result)
	}

	// 最后面4个空格，让json格式更美观
	result, err := utils.CustomMarshal(mapResults)

	if err != nil {
		record.Logger.Error(err.Error())
		return
	}

	if err := ioutil.WriteFile(savePath, []byte(result), 0644); err != nil {
		record.Logger.Error("Write file " + savePath + err.Error() + " error!")
		return
	}

}
