package randoms

import (
	"io"
	"net/http"
	"os"
)

func RandomFact() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/facts?limit=1", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Api-Key", os.Getenv("API_FACT_KEY"))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
