package entities

type CommonAtts struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CornellNoteRequest struct {
	Topic      string       `json:"topic"`
	Subjects   []CommonAtts `json:"subjects"`
	Categories []CommonAtts `json:"categories"`
}
