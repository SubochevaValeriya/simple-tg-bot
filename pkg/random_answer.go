package randoms

import (
	"encoding/json"
	"net/http"
)

type Answer struct {
	Answer string `json:"answer"'`
	Image  string `json:"image"'`
}

func RandomAnswer() (Answer, error) {
	answer := Answer{}
	resp, err := http.Get("https://yesno.wtf/api")
	defer resp.Body.Close()
	if err != nil {
		return answer, err
	}

	err = json.NewDecoder(resp.Body).Decode(&answer)
	if err != nil {
		return answer, err
	}

	return answer, nil
}
