package library

import (
	"bufio"
	"os"
	"strings"

	"Caesar/pkg/record"
	"Caesar/pkg/utils"
)

func GetTargets(targetString string) []string {
	/*
		获取目标列表
	*/
	var targetList []string
	if strings.HasSuffix(targetString, ".txt") {

		file, err := os.Open(targetString)
		if err != nil {
			record.Logger.Error("can not open file " + targetString)
			os.Exit(1)
		}

		defer func() { _ = file.Close() }()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineText := scanner.Text()
			targetList = append(targetList, utils.ConvertAddress(lineText))
		}

	} else {
		targetList = append(targetList, utils.ConvertAddress(targetString))

	}

	return utils.RemoveDuplicateElement(targetList)

}
