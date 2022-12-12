package finder

import (
	"bufio"
	"io"
)

func LoadMap(reader io.Reader) map[string]bool {
	listMap := map[string]bool{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		listMap[scanner.Text()] = true
	}
	return listMap
}

func LoadList(reader io.Reader) []string {
	list := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}
