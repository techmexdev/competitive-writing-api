package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/selecting"
	"github.com/techmexdev/competitive_writing_api/pkg/storage/postgres"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
)

type data struct {
	PP []passage.Passage
	UU []user.User
	WW []writing.Writing
	SS []selecting.Selection
}

func main() {
	dsn := os.Getenv("PG_DSN")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("could not open db connection: %s", err)
	}
	defer db.Close()

	store := postgres.New(db)
	d := createData()

	for _, p := range d.PP {
		err := store.CreatePassage(p)
		if err != nil {
			log.Fatalf("failed creating passages: %s", err)
		}
	}
	pp, err := store.ListPassages()
	if err != nil {
		log.Printf("error listing passages = %s\n", err)
	}

	log.Printf("pp = %+v\n", pp)

	for _, u := range d.UU {
		_, err := store.CreateUser(u.AuthID)
		if err != nil {
			log.Fatalf("failed creating user %s: %s", u.AuthID, err)
		}
	}

	for _, w := range d.WW {
		_, err := store.CreateWriting(w)
		if err != nil {
			log.Fatalf("failed creating writing %v: %s", w, err)
		}
	}

	ww, err := store.ListWritingsWithPassage(writing.Passage{ID: pp[0].ID})
	if err != nil {
		log.Fatalf("failed listing writing w passage %s: %s", pp[0].ID, err)
	}

	log.Printf("ww with id %s = %+v\n", pp[0].ID, ww)

	for _, s := range d.SS {
		err := store.CreateSelection(s)
		if err != nil {
			log.Fatalf("failed creating selection %v: %s", s, err)
		}
	}
}

func createData() data {
	uu := []user.User{
		user.User{AuthID: "02938"},
		user.User{AuthID: "87593"},
		user.User{AuthID: "98240"},
		user.User{AuthID: "75892"},
	}
	uID := uu[0].AuthID

	pp := []passage.Passage{
		{ID: "1", Author: "Immanuel Kant", Book: "The Critique of Pure Reason",
			Text: `That all our knowledge begins with experience there can be no doubt. 
		For how is it possible that the faculty of cognition should be 
		awakened into exercise otherwise than by means of objects which affect 
		our senses, and partly of themselves produce representations, partly 
		rouse our powers of understanding into activity, to compare to 
		connect, or to separate these, and so to convert the raw material of 
		our sensuous impressions into a knowledge of objects, which is 
		called experience? In respect of time, therefore, no knowledge of ours 
		is antecedent to experience, but begins with it.`,
		},
		{ID: "2", Author: "Heidegger", Book: "Being and Time",
			Text: `The Necessity for Explicitly Restating the Question of Being 
		This question has today been forgotten. Even though in our time we 
		deem it progressive to give our approval to 'metaphysics* again, it is held 
		that we have been exempted from the exertions of a newly rekindled 
		ytyavrofia^ta n€pl rijs ovalas. Yet the question we are touching upon is not just 
		any question. It is one which provided a stimulus for the researches of 
		Plato and Aristotle, only to subside from then on or a theme for actual 
		investigation, 1 What these two men achieved was to persist through many 
		alterations and 'retouchings* down to the 'logic* of Hegel. And what 
		they wrested with the utmost intellectual effort from the phenomena, 
		fragmentary and incipient though it was, has long since become 
		trivialized.`,
		},
		{ID: "3", Author: "Lewis Carrol", Book: "Alice in Wonderland", Text: "Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, `and what is the use of a book,' thought Alice `without pictures or conversation?'"},
	}
	p := pp[0]

	ww := []writing.Writing{
		{ID: "1", Author: uID, Passage: writing.Passage{ID: p.ID},
			Text: `Knowledge doesn't exist in a vaccum. We can think of knowledge as human encoed information.
		Books don't contain knowledge, they contain information;
		it is when we read, and understand a piece of information,
		that knowledge is born. Knowledge is not raw data, it comes from either previous knowledge, 
		or directly from experience. A book is a collection of pages that form a story,
		a page is a piece of paper with text on it, text is a collection of characters, etc.
		At some point, the recursion has to stop, and our knowledge will come directly from experience.
		Whatever knowledge we are unable to articulate, comes directly from experience.
		To teach you something, I don't edit the connection between your neurons, I explain how something
		relates to something else, and you make the connections yourself.`},
		{ID: "2", Author: uID, Passage: writing.Passage{ID: p.ID},
			Text: `Live love laugh`},
		{ID: "3", Author: uID, Passage: writing.Passage{ID: p.ID},
			Text: `In the end, the love you get is equal to the love you make.`},
	}
	w := ww[0]
	var selWW []selecting.Writing
	for _, w := range ww {
		selWW = append(selWW, selecting.Writing{ID: w.ID})
	}
	ss := []selecting.Selection{
		{ID: "1", Selector: uID, Passage: selecting.Passage{ID: p.ID}, BestWriting: selecting.Writing{ID: w.ID}, WritingChoices: selWW},
	}
	return data{
		PP: pp, UU: uu, WW: ww, SS: ss,
	}
}
