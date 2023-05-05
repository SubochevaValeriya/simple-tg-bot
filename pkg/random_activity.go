package randoms

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Activity struct {
	Activity string `json:"activity"'`
}

func RandomActivity(minParticipants, maxParticipants int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://www.boredapi.com/api/activity?minparticipants=%v&maxparticipants=%v", minParticipants, maxParticipants))
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	activity := Activity{Activity: ""}

	err = json.NewDecoder(resp.Body).Decode(&activity)
	if err != nil {
		return "", err
	}

	return activity.Activity, nil
}
