package main

type Activity struct {
	Activity      string `json:"activity"`
	Type          string `json:"type"`
	Participants  int    `json:"participants"`
	Price         int    `json:"price"`
	Link          string `json:"link"`
	Key           string `json:"key"`
	Accessibility int    `json:"accessibility"`
}
