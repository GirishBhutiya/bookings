package models

// TemplateDate holds data sent from handlers to template
type TemplateDate struct {
	StringMap  map[string]string
	IntegerMap map[string]int
	FloatMap   map[string]float32
	DataMap    map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
}
