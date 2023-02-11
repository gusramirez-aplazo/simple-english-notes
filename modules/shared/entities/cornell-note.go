package entities

import (
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"gorm.io/gorm"
)

type CornellNote struct {
	gorm.Model

	CornellNoteID uint              `gorm:"type:bigserial;primaryKey;uniqueIndex;autoIncrement" json:"cornellNoteId"`
	Categories    []domain.Category `gorm:"many2many:cornell_categories;foreignKey:CornellNoteID" json:"categories"` // Dictionary, Math, Something else, etc
	Subjects      []domain.Subject  `gorm:"many2many:cornell_subjects;foreignKey:CornellNoteID" json:"subjects"`     // Noun, Verb, Calculus, etc
	Notes         []Note            `gorm:"not null;foreignKey:NoteID" json:"notes"`
	Topic         Topic             `gorm:"foreignKey:TopicID;references:TopicID;not null;unique" json:"topic"` // Stem, Limit, Matrix
	TopicID       uint              `gorm:"type:bigserial;primaryKey;autoIncrement"`
}
