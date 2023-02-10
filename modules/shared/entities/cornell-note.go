package entities

import (
	"gorm.io/gorm"
)

type CornellNote struct {
	gorm.Model

	CornellNoteID uint       `gorm:"type:bigserial;primaryKey;uniqueIndex;autoIncrement" json:"cornellNoteId"`
	Categories    []Category `gorm:"many2many:cornell_categories;foreignKey:CornellNoteID" json:"categories"` // Dictionary, Math, Something else, etc
	Subjects      []Subject  `gorm:"many2many:cornell_subjects;foreignKey:CornellNoteID" json:"subjects"`     // Noun, Verb, Calculus, etc
	Topic         string     `gorm:"not null;unique" json:"topic"`                                            // Stem, Limit, Struct
	Notes         []Note     `gorm:"not null;foreignKey:NoteID" json:"notes"`
}
