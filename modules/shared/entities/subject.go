package entities

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	SubjectID uint   `gorm:"type:bigserial;primaryKey;not null;autoIncrement;uniqueIndex" json:"id"`
	Name      string `gorm:"not null;unique;type:varchar(100)"  json:"name"`
}

type SubjectResponse struct {
	SubjectID uint   `json:"id"`
	Name      string `json:"name"`
}
