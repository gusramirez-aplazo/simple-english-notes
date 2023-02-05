package infra

import (
	"errors"
	"strconv"
)

func ParseID(id string) (uint, error) {
	if len(id) == 0 {
		return 0, errors.New("ID is empty")
	}

	ui64, parseErr := strconv.ParseUint(id, 10, 64)

	if parseErr != nil {
		return 0, parseErr
	}

	return uint(ui64), nil
}
