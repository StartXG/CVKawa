package JTPAction

import (
	"JsonToPdf/common"
	"encoding/json"
	"os"
)

func LoadData(filePath *string) {
	file, err := os.ReadFile(*filePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &common.JsonData)
	if err != nil {
		return
	}
}
