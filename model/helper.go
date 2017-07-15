package model

import (
	"fmt"
	"strconv"
)

func ConvertIdIntoInt64(ID interface{}) (int64, error) {
	strID := fmt.Sprintf("%v", ID)
	int64ID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return 0, err
	}
	if len(strID) != 16 {
		return 0, fmt.Errorf("ID: %v is wrong format ^[0-9]{16}$", int64ID)
	}
	return int64ID, nil
}
