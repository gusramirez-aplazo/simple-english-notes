package entities

type SomeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CornellNoteRequest struct {
	Topic      string        `json:"topic"`
	Subjects   []SomeRequest `json:"subjects"`
	Categories []SomeRequest `json:"categories"`
}
