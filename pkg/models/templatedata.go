package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // we dont know what else could be sent, so we call it an interface i guess? it's just... whatever it recieves. duck type to the max.
	CSRFToken string                 // For the sake of form handling, "cross site request forgery token"
	Flash     string                 // Process response message
	Warning   string
	Error     string
}
