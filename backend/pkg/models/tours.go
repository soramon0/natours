package models

type Tour struct {
	Id              int      `bson:"_id,omitempty" json:"id"`
	Name            string   `bson:"name,omitempty" json:"name,omitempty"`
	Duration        int      `json:"duration"`
	MaxGroupSize    int      `json:"maxGroupSize"`
	Difficulty      string   `json:"difficulty"`
	RatingsAverage  float64  `json:"ratingsAverage"`
	RatingsQuantity int      `json:"ratingsQuantity"`
	Price           int      `json:"price"`
	Summary         string   `json:"summary"`
	Description     string   `json:"description"`
	ImageCover      string   `json:"imageCover"`
	Images          []string `json:"images"`
	StartDates      []string `json:"startDates"`
}
