package entity

// Movie struct
type Cast struct {
	Movie_id  int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Awards    string `json:"awards"`
}
