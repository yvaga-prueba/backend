package model

import "time"

type SizeGuide struct {
	ID        int64     `json:"id" db:"id"`
	Category  string    `json:"category" db:"category"`
	Size      string    `json:"size" db:"size"`
	MinWeight float64   `json:"min_weight" db:"min_weight"`
	MaxWeight float64   `json:"max_weight" db:"max_weight"`
	MinHeight float64   `json:"min_height" db:"min_height"`
	MaxHeight float64   `json:"max_height" db:"max_height"`
	FitType   string    `json:"fit_type" db:"fit_type"`
	ChestCm   float64   `json:"chest_cm" db:"chest_cm"`
	WaistCm   float64   `json:"waist_cm" db:"waist_cm"`
	HipCm     float64   `json:"hip_cm" db:"hip_cm"`
	LengthCm  float64   `json:"length_cm" db:"length_cm"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
