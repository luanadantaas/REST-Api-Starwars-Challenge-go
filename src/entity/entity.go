package entity

type Planet struct {
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Climate     string `json:"climate"`
	Terrain     string `json:"land"`
	Appearances int    `json:"appearances"`
}

func (*Planet) TableName() string {
	return "planet"
}
