package animal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EvgenyiK/animals/datastore"
)

type AnimalHandler struct {
	datastore datastore.Animal
}

func New(animal datastore.Animal) AnimalHandler {
	return AnimalHandler{datastore: animal}
}

//обработчик
func (a AnimalHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.get(w, r)
	case http.MethodPost:
		a.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (a AnimalHandler) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write([]byte("invalid parametr id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp,err:= a.datastore.Get(i)
	if err != nil {
		_, _ = w.Write([]byte("could not retrieve email"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body,_:=json.Marshal(resp)
	_, _= w.Write(body)
}
