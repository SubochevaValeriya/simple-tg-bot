package randoms

import (
	"image"
	"net/http"
)

func RandomCat() (image.Image, error) {
	resp, err := http.Get("https://cataas.com/cat")
	if err != nil {
		return nil, err
	}
	cat, _, err := image.Decode(resp.Body)

	return cat, err
}
