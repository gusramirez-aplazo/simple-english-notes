package domain

import (
	"gorm.io/gorm"
	"time"
)

type CornellNote struct {
	ID         uint           `gorm:"type:bigserial;primaryKey;not null;autoIncrement;unique;uniqueIndex" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Categories []Category     `gorm:"many2many:cornell_categories;foreignKey:ID" json:"categories"` // Dictionary, Math, Something else, etc
	Subjects   []Subject      `gorm:"many2many:cornell_subjects;foreignKey:ID" json:"subjects"`     // Noun, Verb, Calculus, etc
	Notes      []Note         `gorm:"not null;foreignKey:ID" json:"notes"`
	Topic      Topic          `gorm:"foreignKey:ID;references:TopicID;not null;unique" json:"topic"` // Stem, Limit, Matrix
	TopicID    uint           `gorm:"type:bigserial;primaryKey;autoIncrement"`
}
