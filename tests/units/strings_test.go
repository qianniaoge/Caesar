package units

import (
	"Caesar/pkg/utils"
	"testing"
)

func TestRandom(t *testing.T) {
	demo := []string{"shadow"}
	println(utils.GetRandomElement(demo))
	println(utils.GenRandString(4))
}

func TestIps(t *testing.T) {
	println(utils.DomainToIP("www.baidu.com"))
	println(utils.DomainToIP("192.168.3.2"))
}
