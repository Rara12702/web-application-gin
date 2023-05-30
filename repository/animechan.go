package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func GetAnimeQuote() (*Quote, error) {
	resp, err := http.Get("https://animechan.vercel.app/api/random/anime?title=naruto")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch anime quote")
	}

	body, err := ioutil.ReadAll(resp.Body)

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}
