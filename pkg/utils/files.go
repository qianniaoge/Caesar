package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return ""
	}

	str := string(b) // convert content to a 'string'
	return str
}

func DeleteFile(path string) {
	// delete file
	if err := os.Remove(path); err != nil {
		return
	}

}
