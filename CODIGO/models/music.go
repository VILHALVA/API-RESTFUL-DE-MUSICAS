package models

type Music struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Artist   string `json:"artist"`
    Album    string `json:"album"`
    Year     int    `json:"year"`
}
