package models

// Choice : Model
type Choice struct {
	id         int64
	questionID int64
	choice     string
	votesCount int32
}

// SaveChoice : Method
func SaveChoice() {}

// GetChoice : Method
func GetChoice() {}

// DeleteChoice : Method
func DeleteChoice() {}

// UpdateChoice : Method
func UpdateChoice() {}
