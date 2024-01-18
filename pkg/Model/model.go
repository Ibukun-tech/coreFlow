package model

// TemplateData hold data set from handler to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFtoken string
	Flash     string
	Warning   string
	Error     string
}
