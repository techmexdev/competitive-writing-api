package rest_test

// func TestUserSignup(t *testing.T) {
// 	passStore := &memo.PassageStorage{}
//
// 	wri := writing.New(passStore)
// 	ath := auth0.New(auth0.Config{})
// 	handler := rest.New(rest.Config{Writing: wri, Auth: ath}, logrus.New())
//
// 	creds := auth.Creds{Email: "email@example.com", Pwd: "password123"}
// 	body, err := json.MarshalIndent(creds, "", " ")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
//
// 	handler.ServeHTTP(w, r)
//
// 	if w.Code != http.StatusCreated {
// 		t.Fatalf("have response code %v, want %v", w.Code, http.StatusCreated)
// 	}
//
// 	respDump, err := httputil.DumpResponse(w.Result(), true)
// 	if err != nil {
// 		log.Fatalf("failed creating resp dump: %s", err)
// 	}
//
// 	type resp struct {
// 		Token string `json:"token"`
// 	}
//
// 	var res resp
// 	err = json.NewDecoder(w.Body).Decode(&res)
// 	if err != nil {
// 		t.Fatalf("failed decoding from json response into token: %s", err)
// 	}
//
// 	if len(res.Token) == 0 {
// 		t.Fatalf("want non-empty token. have response: %s", respDump)
// 	}
// }
