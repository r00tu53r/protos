package ftp

import (
	"strconv"
	"strings"

	"github.com/r00tu53r/protos/ftp/fsm"
)

func getCode(word string) (int, error) {
	last := len(word) - 1
	code, err := strconv.Atoi(word)
	if err != nil {
		if word[last] == '-' {
			code, err = strconv.Atoi(word[:last])
			if err != nil {
				return -1, err
			}
			return code, nil
		}
		return -1, err
	}
	return code, nil
}

func isRequest(line string) (string, bool) {
	firstWord := strings.Split(line, " ")[0]
	_, ok := fsm.Commands[strings.ToUpper(firstWord)]
	return firstWord, ok
}

func isResponse(line string) (int, bool) {
	firstWord := strings.Split(line, " ")[0]
	code, err := getCode(firstWord)
	if err != nil {
		return -1, false
	}
	return code, true
}
