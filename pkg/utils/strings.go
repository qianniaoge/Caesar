package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"Caesar/pkg/buoys"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

//生成随机字符串的函数
func GenRandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

	var (
		letterIdxBits       = 6
		letterIdxMask int64 = 1<<letterIdxBits - 1
		letterIdxMax        = 63 / letterIdxBits
		src                 = rand.NewSource(time.Now().UnixNano())
	)

	b := make([]byte, n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//  判断字符是否在字符列表中
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//  判断数字是否在列表中
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ConvertAddress(target string) string {
	// 检查目标是否是 http://target的格式
	/*
		@param: http://127.0.0.1/ || 127.0.0.1
		@return: http://127.0.0.1
	*/
	NotLine := "^(http://|https://).*"
	match, _ := regexp.MatchString(NotLine, target)

	if !match {
		target = "http://" + target
	}

	return strings.TrimSuffix(target, "/")
}

func GetValueFromList(key string, world []map[string]string) string {
	/*
	  从列表map中获取value
	  param: key1, [key1:value1, key2:value2, key3:value3]
	  return: value1
	*/
	for _, v := range world {

		if v, ok := v[key]; ok {
			return v
		}

	}

	return ""

}

func Input() string {
	/*
		类似Python3的input()函数
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "n"
	}

	return strings.ToLower(scanner.Text())

}

func GetRandomElement(lists []string) string {
	/*
		从slice选取一个随机值
	*/
	rand.Seed(time.Now().UnixNano())
	element := lists[rand.Intn(len(lists))]
	return element
}

// 探测网页编码
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytesHtml, err := r.Peek(1024)

	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytesHtml, "")

	return e

}

func GenStrings(param ...string) string {
	/*
		拼接字符串，基准测试的时候发现效率不如 +
	*/

	var build strings.Builder

	for _, v := range param {
		build.WriteString(v)

	}
	return build.String()

}

func CustomMarshal(message interface{}) (string, error) {
	/*
		自定义序列化函数，解决 "&"被转译的问题
	*/

	bf := bytes.NewBuffer([]byte{})

	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "    ")

	if err := jsonEncoder.Encode(message); err != nil {
		return buoys.ErrorFlag, err
	}

	return bf.String(), nil
}

// golang读取文件并且返回列表
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// 字符串列表去重函数
func RemoveDuplicateElement(infos []string) []string {
	result := make([]string, 0, len(infos))
	temp := map[string]struct{}{}
	for _, item := range infos {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//获取当前时间的字符串形式
func GenNowTime() string {
	datetimeStr := time.Now().Format("2006-01-02 15:04:05")
	return datetimeStr
}
