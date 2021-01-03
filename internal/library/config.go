package library

import (
	"io/ioutil"
	"os"

	"Caesar/pkg/record"

	"gopkg.in/yaml.v2"
)

type configs struct {
	Silence         bool              `yaml:"Silence"`
	DisplayCode     []int             `yaml:"DisplayCode"`
	UpperRatioBound float64           `yaml:"UpperRatioBound"`
	WafTop          int               `yaml:"WafTop"`
	TimeoutCount    int               `yaml:"TimeoutCount"`
	SuffixConnector []string          `yaml:"SuffixConnector"`
	DirectorySuffix []string          `yaml:"DirectorySuffix"`
	DynamicSuffix   []string          `yaml:"DynamicSuffix"`
	TimeOut         int               `yaml:"TimeOut"`
	UserAgent       []string          `yaml:"UserAgent"`
	Proxy           string            `yaml:"Proxy"`
	Cookie          string            `yaml:"Cookie"`
	Headers         map[string]string `yaml:"Headers"`
}

type Config struct {
	Filename string
}

func (c *Config) LoadConfigFromYaml() *configs {
	/*
		载入配置文件
	*/
	content, err := ioutil.ReadFile(c.Filename)

	if err != nil {
		record.Logger.Fatal("Can not found config")
		os.Exit(1)
	}

	var confBase configs

	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err := yaml.Unmarshal(content, &confBase); err != nil {
		record.Logger.Fatal(err)
		os.Exit(1)
	}
	return &confBase
}

// 构造函数
func NewProfile(fileName string) *Config {
	c := &Config{
		Filename: fileName,
	}
	return c
}
