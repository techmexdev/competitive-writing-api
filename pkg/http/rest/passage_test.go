package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/techmexdev/competitive_writing_api/pkg/http/rest"
	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/storage/memo"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
)

func TestPassageList(t *testing.T) {
	passStore := memo.NewPassageStorage(&memo.StatefulData{})

	psgs := []passage.Passage{
		{
			ID: "0", Author: "author", Book: "book", Text: "text",
		},
		{
			ID: "1", Author: "author", Book: "book", Text: "text",
		},
		{
			ID: "2", Author: "author", Book: "book", Text: "text",
		},
	}
	for _, p := range psgs {
		err := passStore.CreatePassage(p)
		if err != nil {
			t.Fatalf("failed creating passage: %s", err)
		}
	}

	pass := passage.NewService(passStore)
	handler := rest.New(rest.Config{Psg: pass}, logrus.New())

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/passage", nil)
	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("have response code %v, want %v", w.Code, http.StatusOK)
	}

	var pp []writing.Passage
	err := json.NewDecoder(w.Body).Decode(&pp)
	if err != nil {
		t.Fatalf("failed decoding from json response into passages: %s", err)
	}

	if len(pp) != len(psgs) {
		t.Fatalf("have len %v, want %v", len(pp), len(psgs))
	}

	if !reflect.DeepEqual(pp, psgs) {
		t.Fatalf("have passages: %v, want %v", pp, psgs)
	}
}
