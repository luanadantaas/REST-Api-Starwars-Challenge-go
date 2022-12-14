package entity

type Planet struct {
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Climate     string `json:"climate"`
	Terrain     string `json:"land"`
	Appearances int    `json:"appearances"`
}

type ApiResponse struct {
	Result []results `json:"results"`
}

type results struct {
	Films []string `json:"films"`
}

func (*Planet) TableName() string {
	return "planet"
}
