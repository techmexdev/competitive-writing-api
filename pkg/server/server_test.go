package server_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	cw "github.com/techmexdev/competitive_writing_api"
	srvMock "github.com/techmexdev/competitive_writing_api/pkg/server/mock"
	storeMock "github.com/techmexdev/competitive_writing_api/pkg/storage/mock"
)

func TestPassagesList(t *testing.T) {
	store := storeMock.New()
	srv := srvMock.New(store)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/passage", nil)
	srv.ServeHTTP(w, r)

	var pp []cw.Passage
	err := json.Unmarshal(w.Body.Bytes(), &pp)
	if err != nil {
		log.Println(err)
	}

	var data bytes.Buffer
	json.Indent(&data, w.Body.Bytes(), "", "  ")
	log.Printf("resp: %s", data.String())
}

func TestAnalysisList(t *testing.T) {
	store := storeMock.New()
	srv := srvMock.New(store)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/analysis", nil)
	srv.ServeHTTP(w, r)

	var aa []cw.Analysis
	err := json.Unmarshal(w.Body.Bytes(), &aa)
	if err != nil {
		log.Println(err)
	}

	var data bytes.Buffer
	json.Indent(&data, w.Body.Bytes(), "", "  ")
	log.Printf("resp: %s", data.String())
}

func TestSelectionList(t *testing.T) {
	store := storeMock.New()
	srv := srvMock.New(store)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/selection", nil)
	srv.ServeHTTP(w, r)

	var ss []cw.Selection
	err := json.Unmarshal(w.Body.Bytes(), &ss)
	if err != nil {
		log.Println(err)
	}

	var ppJSON bytes.Buffer
	json.Indent(&ppJSON, w.Body.Bytes(), "", "  ")
	log.Printf("resp: %s", ppJSON.String())
}
