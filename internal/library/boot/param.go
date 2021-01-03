package boot

import (
	"io/ioutil"
	"os"
	"strings"

	"Caesar/pkg/record"
)

func GetFlag(document string) string {
	var stringKey []string
	var key string

	//获取指定目录下的路径文件
	fileInfoList, err := ioutil.ReadDir(document)
	if err != nil {
		record.Logger.Fatal(document + " can not read")
		os.Exit(1)
	}

	for i := range fileInfoList {
		if strings.HasSuffix(fileInfoList[i].Name(), ".json") {
			stringKey = append(stringKey, strings.Split(fileInfoList[i].Name(), ".json")[0])
		}
	}
	key = strings.Join(stringKey, ",")
	return key

}
