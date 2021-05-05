package main

type Node struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	NodeType string `json:"nodeType"`
}

type Rel struct {
	PID string
	ID  string
}
