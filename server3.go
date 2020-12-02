package main 
import "net/http"
type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"inPark"`
	Height       int    `json:"height"`
}
type coasterHandlers struct {
	
	store map[string]Coaster
}

func (h *coastersHandlers)get(w http.ResponseWriter, r *http.Request) {
}
func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{},
	}

}
func main() {
CoasterHandlers := newCoasterHandlers()
http.HandleFunc("/coasters", coastersHandlers.get)
err := http.ListenAndServe(":8080", nil)
if err != nil {
panic(err)
}
}