package models

import (
	"encoding/json"
	"time"
)

type Transcription struct {
	Text      string
	CreatedAt time.Time
}

func (t Transcription) ToJson() string {
	transcriptionJson := map[string]string{
		"transcription": t.Text,
	}

	json, _ := json.Marshal(transcriptionJson)
	return string(json)
}
