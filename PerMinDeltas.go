package go_league_of_legends

type PerMinDeltas struct {
	First  float32 `json:"0-10"`
	Second float32 `json:"10-20"`
	Third  float32 `json:"20-30"`
	Final  float32 `json:"30-end"`
}
