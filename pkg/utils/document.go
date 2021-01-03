package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 递归获取路径下的所有文件
func GetFileFromDocument(rootPath string) []string {
	var fileList []string

	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fileList = append(fileList, path)

			return nil
		})
	if err != nil {

		return nil
	}

	return fileList
}

//利用装饰器增强一下读取目录下的文件的函数的功能
func WrapFuncGetFile(f func(rootPath string) []string, path string, fileType string) (newList []string) {
	// 获取特定类型后缀的文件
	fileList := f(path)

	for _, v := range fileList {
		if strings.HasSuffix(v, fileType) {
			newList = append(newList, v)
		}

	}

	return
}
