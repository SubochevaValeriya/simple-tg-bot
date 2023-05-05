package randoms

import (
	"io"
	"net/http"
)

func RandomCat() ([]byte, error) {
	resp, err := http.Get("https://cataas.com/cat")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, err
}
