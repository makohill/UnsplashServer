package main

type Unsplash struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	FullUrl   string    `json:"fullurl"`
}

type Unsplashs []Unsplash
