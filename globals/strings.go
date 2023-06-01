package globals

import (
	"errors"
	"strings"
)

func SafeSplit(s string, sep1 string, sep2 string) (string, error) {
	splits := strings.Split(s, sep1)
	if len(splits) <= 1 {
		return "", errors.New("split length on sep1 equals 1")
	}

	splits = strings.Split(splits[1], sep2)
	if len(splits) <= 1 {
		return "", errors.New("split length on sep2 equals 1")
	}

	return splits[0], nil
}
