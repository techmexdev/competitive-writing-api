package memo_test

import (
	"testing"

	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/storage/memo"
)

func TestPassageCreate(t *testing.T) {
	store := memo.NewPassageStorage(&memo.StatefulData{})
	psg := passage.Passage{
		ID: "0", Author: "author", Book: "book", Text: "text",
	}
	err := store.CreatePassage(psg)
	if err != nil {
		t.Fatalf("failed creating passage: %s", err)
	}
}

func TestPassageRead(t *testing.T) {
	store := memo.NewPassageStorage(&memo.StatefulData{})
	psg := passage.Passage{
		ID: "0", Author: "author", Book: "book", Text: "text",
	}
	err := store.CreatePassage(psg)
	if err != nil {
		t.Fatalf("failed creating passage: %s", err)
	}

	p, err := store.ReadPassage(psg.ID)
	if err != nil {
		t.Fatalf("failed reading passage: %s", err)
	}

	if p.ID != psg.ID {
		t.Fatalf("have passage ID %s, want %s", p.ID, psg.ID)
	}
}

func TestPassageList(t *testing.T) {
	var store passage.Storage
	store = memo.NewPassageStorage(&memo.StatefulData{})
	psgs := []passage.Passage{
		{
			ID: "0", Author: "author", Book: "book", Text: "text",
		},
		{
			ID: "1", Author: "author", Book: "book", Text: "text",
		},
	}
	for _, p := range psgs {
		err := store.CreatePassage(p)
		if err != nil {
			t.Fatalf("failed creating passage: %s", err)
		}
	}

	pp, err := store.ListPassages()
	if err != nil {
		t.Fatalf("failed listing passages: %s", err)
	}
	if len(pp) != len(psgs) {
		t.Fatalf("have len %v, want %v", len(pp), len(psgs))
	}

	// TODO: compare pp and psgs contents
}
