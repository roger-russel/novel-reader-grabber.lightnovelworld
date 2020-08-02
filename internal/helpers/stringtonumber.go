package helpers

import (
	"strconv"
	"strings"
)

//StringToNumber get string return number, detect number like 100-101 and return 100
func StringToNumber(s string) (int, error) {

	number, err := strconv.Atoi(s)

	if err != nil {
		i := strings.Index(s, "-")
		if i > 0 {
			number, err = strconv.Atoi(s[:i])
		}
	}

	return number, err

}
