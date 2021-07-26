package format

import (
	"fmt"
)

func ToIndentedKeyValue(strs ...string) string {
	formatted := ""
	i := 0

	if len(strs)%2 != 0 {
		return formatted
	}

	for i < len(strs) {
		if i >= 0 && i < len(strs) {
			formatted += fmt.Sprintf("%s: \t %v \n", strs[i], strs[i+1])
			i = i + 2
		}

		if i >= len(strs) {
			formatted += "\n"
			break
		}
	}

	return formatted
}
