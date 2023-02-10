package entities

import "gorm.io/gorm"

type Note struct {
	gorm.Model

	NoteID  uint   `gorm:"type:bigserial;primaryKey;not null;autoIncrement;unique;uniqueIndex" json:"id"`
	Cue     string `gorm:"type:varchar(512)" json:"cue"`
	Content string `gorm:"not null" json:"content"`
}
