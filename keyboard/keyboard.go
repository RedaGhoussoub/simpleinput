package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInt() (int, bool) {
	if input, ok := read(); ok {
		value, err := strconv.Atoi(input)
		if err != nil {
			return 0, false
		}
		return value, true
	}
	return 0, false
}

func ReadFloat() (float64, bool) {
	if input, ok := read(); ok {
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	}
	return 0, false
}

func ReadString() (string, bool) {
	return read()
}

func read() (string, bool) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", false
	}
	input = strings.TrimSpace(input)
	return input, true
}
