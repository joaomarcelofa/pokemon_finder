package finder

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadList(t *testing.T) {
	r := strings.NewReader("bulbasaur\nivysaur\n")
	want := map[string]bool{
		"bulbasaur": true,
		"ivysaur":   true,
	}
	result := loadList(r)
	require.Equal(t, want, result, "list loaded a different config")
}
