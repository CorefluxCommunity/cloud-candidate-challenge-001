package main

type CreateRequest struct {
	DropletName string   `json:"droplet_name"`
	Region      string   `json:"region"`
	Size        string   `json:"size"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
}

type SearchRequest struct {
	TagToFind []string `json:"tag_to_find"`
}

type SortRequest struct {
	Direction string `json:"direction"`
}
