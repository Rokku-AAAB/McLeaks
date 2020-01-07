package mcleaks

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// McLeaksResult struct
type McLeaksResult struct {
        McName  string `json:"mcname,omitempty"`
        Session string `json:"session,omitempty"`
}

// McLeaksResponse struct
type McLeaksResponse struct {
        Success bool `json:"success,omitempty"`
        Result  McLeaksResult `json:"result,omitempty"`
}

// McLeaksEmitToken sends a given Token to the mcleaks API and returns its Response decoded to MacLeaksResponse
func McLeaksEmitToken(token []byte) (McLeaksResponse, error) {
	resp, err1 := http.Post("https://auth.mcleaks.net/v1/redeem", "application/json", bytes.NewBuffer(token))

	if err1 != nil {
		return McLeaksResponse{}, err1
	}

	respData, err2 := ioutil.ReadAll(resp.Body)

        if err2 != nil {
                return McLeaksResponse{}, err2
        }

	var mcResponse McLeaksResponse
	err3 := json.Unmarshal(respData, &mcResponse)

        if err3 != nil {
                return McLeaksResponse{}, err3
        }

	if !mcResponse.Success {
		return McLeaksResponse{}, errors.New(string(respData))
	}

	return mcResponse, nil
}
