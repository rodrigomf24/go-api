package main

import "time"

type Item struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Active		bool		`json:"active"`
	CreatedDate	time.Time	`json:"created_date"`
}

type Items []Item