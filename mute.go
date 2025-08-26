package twitterscraper

import (
	"bytes"
	"io"
)

func (s *Scraper) MuteUser(userId string) error {
	req, err := s.newRequest("POST", "https://x.com/i/api/1.1/mutes/users/create.json")
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	strBody := "user_id="+userId

	req.Body = io.NopCloser(bytes.NewReader([]byte(strBody)))

	var response struct {
	}

	err = s.RequestAPI(req, &response)
	if err != nil {
		return err
	}

	return nil
}

