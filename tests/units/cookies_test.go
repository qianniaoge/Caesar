package units

import (
	"Caesar/pkg/utils"
	"testing"
)

/*
用来测试cookie的正则表达式
*/
func TestCookies(t *testing.T) {
	mat := "username=admin; userid=1; PHPSESSID=9d1q9o4927a42p2thki1ql82p7"
	println(utils.MatchCookie(mat))
}
