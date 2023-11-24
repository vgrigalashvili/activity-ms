package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next: next}

}
func (s *LoggingService) GetActivity(ctx context.Context) (activity *Activity, err error) {
	defer func(start time.Time) {
		fmt.Printf("- Activity: %s\n- error: %v\n- duration: %v\n", activity.Activity, err, time.Since(start))
	}(time.Now())
	return s.next.GetActivity(ctx)
}
