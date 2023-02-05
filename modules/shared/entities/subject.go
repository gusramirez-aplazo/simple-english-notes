package entities

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	SubjectID   uint   `gorm:"type:bigserial;primaryKey;not null;autoIncrement;uniqueIndex" json:"subjectId"`
	Name        string `gorm:"not null;unique;type:varchar(100)"  json:"name"`
	Description string `json:"description"`
}
