package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	textiterator "github.com/joaomarcelofa/pokemon_finder/text_iterator"
	"github.com/stretchr/testify/require"
)

func TestGetPokemon(t *testing.T) {
	handler, err := newPokemonHandler()
	require.NoError(t, err)

	t.Run("testing a valid input", func(t *testing.T) {
		req, err := mountNewRequest("Eu sou um treinador pokémon e tenho um Pikachu, mas também tenho um Gengar")
		require.NoError(t, err)
		res := httptest.NewRecorder()

		expectedResult := []textiterator.Word{
			{
				Text:    "Pikachu",
				StartAt: 39,
				EndAt:   45,
			},
			{
				Text:    "Gengar",
				StartAt: 68,
				EndAt:   73,
			},
		}

		handler.get(res, req)

		result, err := getWords(res.Result().Body)
		require.NoError(t, err)

		require.Equal(t, res.Result().StatusCode, http.StatusOK, "expecting status code 200")
		require.Equal(t, result, expectedResult, "expecting the a specific, another was returned")
	})
}

func mountNewRequest(text string) (*http.Request, error) {
	input := input{text}
	payloadBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(payloadBytes)

	request, err := http.NewRequest(http.MethodGet, "/pokemon", reader)
	return request, err
}

func getWords(body io.ReadCloser) ([]textiterator.Word, error) {
	values := []textiterator.Word{}
	d := json.NewDecoder(body)
	err := d.Decode(&values)
	if err != nil {
		return nil, err
	}
	return values, err
}
