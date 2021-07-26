package file

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSON(filePath string, data interface{}) (err error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileBytes, data)
	if err != nil {
		return
	}

	return
}
