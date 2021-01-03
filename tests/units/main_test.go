package units

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMainJson(t *testing.T) {
	var info []map[int]string

	for i := 0; i < 10; i++ {
		var num = make(map[int]string)
		num[i] = "&lsl"
		info = append(info, num)
	}

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "    ")
	jsonEncoder.Encode(info)
	fmt.Println(bf.String())

}
