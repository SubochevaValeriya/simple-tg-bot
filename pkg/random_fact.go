package randoms

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Fact struct {
	Fact string `json:"fact"'`
}

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
	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))
	if err != nil {
		return "", err
	}
	var fact []Fact

	err = json.Unmarshal(data, &fact)
	if err != nil {
		return "", err
	}

	log.Println(fact)

	return fact[0].Fact, nil
}
