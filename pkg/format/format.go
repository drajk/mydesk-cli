package format

import (
	"fmt"
)

func ToIndentedKeyValue(strs ...string) (ret string, err error) {
	i := 0
	for {
		if i >= 0 && i < len(strs) {
			ret += fmt.Sprintf("%s: \t %v \n", strs[i], strs[i+1])
			i = i + 2
		} else {
			ret += "\n"
			break
		}
	}

	return
}
