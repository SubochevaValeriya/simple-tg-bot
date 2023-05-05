package randoms

import (
	"encoding/json"
	"net/http"
)

type DogImage struct {
	URL string `json:"url"'`
}

func RandomDog() (string, error) {
	resp, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	dogImage := DogImage{URL: ""}

	err = json.NewDecoder(resp.Body).Decode(&dogImage)
	if err != nil {
		return "", err
	}

	return dogImage.URL, nil
}
