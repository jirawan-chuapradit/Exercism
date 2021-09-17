package reverse

import (
	"strings"
)

func Reverse(remark string) string {

	var ans string

	temp := strings.Split(remark, "")

	for j := len(temp) - 1; j >= 0; j-- {
		ans = ans + string(temp[j])
	}
	return ans
}
