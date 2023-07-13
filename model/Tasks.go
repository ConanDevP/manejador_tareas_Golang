package model

type Tasks struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}
