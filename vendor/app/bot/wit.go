package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	confidenceThreshold = 0.7
)

type witMessage struct {
	Entities map[string][]struct {
		Value      string  `json:"value"`
		Confidence float64 `json:"confidence"`
	} `json:"entities"`
}

func getEntities(message, token string) map[string]string {
	entities := make(map[string]string)

	witResponse, err := callWitAPI(message, token)
	if err != nil {
		log.Println(err)
		return nil
	}

	for entity, values := range witResponse.Entities {
		var bestValue string
		maxConfidence := float64(0)

		for _, entityResult := range values {
			if (entityResult.Confidence > confidenceThreshold) &&
				(entityResult.Confidence > maxConfidence) {
				bestValue = entityResult.Value
				maxConfidence = entityResult.Confidence
			}
		}

		entities[entity] = bestValue
	}

	return entities
}

func callWitAPI(message, token string) (*witMessage, error) {
	// Build request
	q := url.Values{}
	q.Add("q", message)

	url := url.URL{
		Scheme:   "https",
		Host:     witAPI,
		Path:     "/message",
		RawQuery: q.Encode(),
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		return nil, err
	}

	// Perform request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	defer resp.Body.Close()
	var entities witMessage

	err = json.NewDecoder(resp.Body).Decode(&entities)
	if err != nil {
		return nil, err
	}

	return &entities, nil
}
