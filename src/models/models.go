package models

type BabyName struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Rank   int    `json:"rank"`
}

type RatingsFile struct {
	BabyNames []BabyName `json:"names"`
}
