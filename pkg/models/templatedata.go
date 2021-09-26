package models

// TemplateData holds data sent from Handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int64
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}