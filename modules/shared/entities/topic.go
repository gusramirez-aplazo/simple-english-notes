package entities

import "gorm.io/gorm"

type Topic struct {
	gorm.Model

	TopicID uint   `gorm:"type:bigserial;primaryKey;not null;autoIncrement;unique;uniqueIndex" json:"id"`
	Name    string `gorm:"not null;unique;type:varchar(100)"  json:"name"`
}