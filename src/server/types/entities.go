package types

// Person represents a person and their attributes.
type Person struct {
	Colors   []string `json:"colors"`
	ID       int      `json:"id"`
	ImgURL   string   `json:"imgUrl"`
	Location string   `json:"location"`
	Name     string   `json:"name"`
}
