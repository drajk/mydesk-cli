package format

import (
	"fmt"
)

func ToIndentedKeyValue(strs ...string) string {
	fomatted := ""
	i := 0

	for {
		if i >= 0 && (i+1) < len(strs) {
			fomatted += fmt.Sprintf("%s: \t %v \n", strs[i], strs[i+1])
			i = i + 2
		}

		if i == 0 {
			break
		}

		if i >= len(strs) {
			fomatted += "\n"
			break
		}
	}

	return fomatted
}
