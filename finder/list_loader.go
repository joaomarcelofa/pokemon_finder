package finder

import (
	"bufio"
	"io"
)

func loadList(reader io.Reader) map[string]bool {
	listMap := map[string]bool{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		listMap[scanner.Text()] = true
	}
	return listMap
}
