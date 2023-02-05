package entities

import (
	"gorm.io/gorm"
)

type CornellNote struct {
	gorm.Model

	CornellNoteID uint       `gorm:"type:bigserial;primaryKey;uniqueIndex;autoIncrement" json:"cornellNoteId"`
	Categories    []Category `gorm:"many2many:cornell_categories;foreignKey:CornellNoteID" json:"categories"` // Dictionary, Math, Something else, etc
	Subjects      []Subject  `gorm:"many2many:cornell_topics;foreignKey:CornellNoteID" json:"subjects"`       // Noun, Verb, Calculus, etc
	Topic         string     `gorm:"not null;unique" json:"topic"`                                            // Stem, Limit, Struct
	//CategoryID    uint       `gorm:"type:bigserial;not null;autoIncrement;uniqueIndex" json:"categoryId"`
	//SubjectID       uint       `gorm:"type:bigserial;not null;autoIncrement;uniqueIndex" json:"topicId"`
	//Question      []RelevantQuestion `json:"questions"`                                        // Is there a way to indicate that something is...
	//Prompt        []RecallPrompt     `json:"prompts"`                                          // Limit in Math is we want to know ...
}
