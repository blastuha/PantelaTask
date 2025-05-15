package utils

import (
	"fmt"
	"strconv"
)

var ErrInvalidID = fmt.Errorf("invalid ID format")

func ParseUintID(s string) (uint64, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("ParseUintID: %w", ErrInvalidID)
	}

	return uid, nil
}
