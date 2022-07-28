package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	value := strings.TrimSpace(input)
	return value
}
