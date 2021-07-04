package zaloai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ZaLoAIRequestData ...
type ZaLoAIRequestData struct {
	URL       string
	Key       string
	Input     string
	SpeakerID string // 1,2,3
}

// AudioResponseZaloAI ...
type AudioResponseZaloAI struct {
	Error     error                   `json:"error"`
	ErrorCode int                     `json:"error_code"`
	Msg       string                  `json:"error_message"`
	RequestID string                  `json:"request_id"`
	Data      AudioResponseZaloAIData `json:"data"`
}

// AudioResponseZaloAIData ...
type AudioResponseZaloAIData struct {
	URL string `json:"url"`
}

// CheckValidData ...
func (req ZaLoAIRequestData) CheckValidData() bool {
	if req.Key == "" || req.URL == "" || req.SpeakerID == "" || req.Input == "" {
		return false
	}
	return true
}

// GenerateAudioByContent ...
func GenerateAudioByContent(data ZaLoAIRequestData) AudioResponseZaloAI {
	// Check empty data request
	if !data.CheckValidData() {
		return AudioResponseZaloAI{
			Error: errors.New(fmt.Sprintf("%s", "Error input Data")),
		}
	}

	// Get body
	body := url.Values{}
	body.Set("speaker_id", data.SpeakerID)
	body.Set("input", data.Input)
	strings.NewReader(body.Encode())

	// Request New
	r, err := http.NewRequest("POST", data.URL, strings.NewReader(body.Encode())) // URL-encoded payload
	if err != nil {
		return AudioResponseZaloAI{
			Error: err,
		}
	}

	// Set header
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))
	r.Header.Set("apikey", data.Key)

	// Client
	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	resp, err := client.Do(r)
	if err != nil {
		return AudioResponseZaloAI{
			Error: err,
		}
	}
	defer resp.Body.Close()

	by, err := ioutil.ReadAll(resp.Body)
	var (
		res AudioResponseZaloAI
	)
	json.Unmarshal(by, &res)

	return res

}
