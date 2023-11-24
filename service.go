package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetActivity(context.Context) (*Activity, error)
}

type ActivityService struct {
	url string
}

func NewActivityService(url string) Service {
	return &ActivityService{
		url: url,
	}

}

func (s *ActivityService) GetActivity(ctx context.Context) (*Activity, error) {
	res, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	activity := &Activity{}
	if err := json.NewDecoder(res.Body).Decode(activity); err != nil {
		return nil, err
	}
	return activity, nil
}
