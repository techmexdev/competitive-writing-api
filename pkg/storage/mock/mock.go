package mock

import cw "github.com/techmexdev/competitive_writing_api"

// New creates Storage with hard-coded data stored.
func New() cw.Storage {
	uu := []cw.User{
		{Username: "username", Email: "email@example.com", Password: "Password123"},
		{Username: "username1", Email: "username1@example.com", Password: "Password123"},
		{Username: "username2", Email: "username2@example.com", Password: "Password123"},
	}
	usn := uu[0].Username

	pp := []cw.Passage{
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

	aa := []cw.Analysis{
		{ID: "1", Username: usn, Passage: p,
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
		{ID: "2", Username: uu[1].Username, Passage: p,
			Text: `Live love laugh`},
		{ID: "3", Username: uu[1].Username, Passage: p,
			Text: `In the end, the love you get is equal to the love you make.`},
	}
	a := aa[0]
	ss := []cw.Selection{
		{ID: "1", Username: usn, Passage: p, BestAnalysis: a, AnalysisChoices: aa},
	}
	return cw.Storage{
		User:      &userStorage{users: uu},
		Passage:   &passageStorage{passages: pp},
		Analysis:  &analysisStorage{analyses: aa},
		Selection: &selectionStorage{selections: ss},
	}
}

type userStorage struct {
	users []cw.User
}

func (s *userStorage) Create(usr cw.User) (err error) {
	return nil
}

func (s *userStorage) Read(username string) (usr cw.User, err error) {
	return s.users[0], nil
}

func (s *userStorage) List() (usr []cw.User, err error) {
	return s.users, nil
}

type passageStorage struct {
	passages []cw.Passage
}

func (s *passageStorage) Create(par cw.Passage) (err error) {
	return nil
}

func (s *passageStorage) Read(id string) (par cw.Passage, err error) {
	return s.passages[0], nil
}

func (s *passageStorage) List() (par []cw.Passage, err error) {
	return s.passages, nil
}

type analysisStorage struct {
	analyses []cw.Analysis
}

func (s *analysisStorage) Create(par cw.Analysis) (err error) {
	return nil
}

func (s *analysisStorage) Read(id string) (par cw.Analysis, err error) {
	return s.analyses[0], nil
}

func (s *analysisStorage) List() (par []cw.Analysis, err error) {
	return s.analyses, nil
}

type selectionStorage struct {
	selections []cw.Selection
}

func (s *selectionStorage) Create(par cw.Selection) (err error) {
	return nil
}

func (s *selectionStorage) Read(id string) (par cw.Selection, err error) {
	return s.selections[0], nil
}

func (s *selectionStorage) List() (par []cw.Selection, err error) {
	return s.selections, nil
}
