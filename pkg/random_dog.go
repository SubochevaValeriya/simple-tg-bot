package randoms

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DogImage struct {
	URL string `json:"url"'`
}

func RandomDog() (string, string, error) {
	resp, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	dogImage := DogImage{URL: ""}

	err = json.NewDecoder(resp.Body).Decode(&dogImage)
	if err != nil {
		return "", "", err
	}
	fmt.Println(dogImage.URL[len(dogImage.URL)-3 : len(dogImage.URL)])
	if dogImage.URL[len(dogImage.URL)-3:len(dogImage.URL)] == "gif" {
		return "", dogImage.URL, nil
	}

	return dogImage.URL, "", nil
}
