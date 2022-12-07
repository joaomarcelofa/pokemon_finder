package finder

import (
	"bufio"
	"io"
)

func loadMap(reader io.Reader) map[string]bool {
	listMap := map[string]bool{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		listMap[scanner.Text()] = true
	}
	return listMap
}

func loadList(reader io.Reader) []string {
	list := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}
