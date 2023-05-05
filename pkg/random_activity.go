package randoms

import (
	"fmt"
	"io"
	"net/http"
)

func RandomActivity(minParticipants, maxParticipants int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://www.boredapi.com/api/activity?minparticipants=%v&maxparticipants=%v", minParticipants, maxParticipants))
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
