package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[float32]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

//mux := pat.New()
//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
//mux.Get("/about", http.HandlerFunc(handlers.Repo.Home))
