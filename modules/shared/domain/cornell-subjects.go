package domain

import (
	"gorm.io/gorm"
	"time"
)

type CornellSubjects struct {
	CornellNoteID uint           `gorm:"type:bigserial;primaryKey;" json:"cornell_note_id"`
	SubjectID     uint           `gorm:"type:bigserial;primaryKey;" json:"subject_id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
