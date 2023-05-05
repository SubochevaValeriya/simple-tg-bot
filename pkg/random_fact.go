package randoms

import (
	"io"
	"net/http"
)

func RandomFact() (string, error) {
	resp, err := http.Get("https://api.api-ninjas.com/v1/facts?limit=1")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
