package api

// func TestGetAccountAPI(t *testing.T) {
// 	account := createRandomAccount()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	store := mockdb.NewMockStore(ctrl)

// 	//build Stub
// 	store.EXPECT().
// 		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
// 		Times(1).
// 		Return(account, nil)

// 	//start test server and send req
// 	server := newTestServer(t, store)
// 	recorder := httptest.NewRecorder()

// 	url := fmt.Sprintf("/account/%d", account.ID)
// 	request, err := http.NewRequest(http.MethodGet, url, nil)

// 	require.NoError(t, err)

// 	server.router.ServeHTTP(recorder, request)
// 	//check response
// 	require.Equal(t, http.StatusOK, recorder.Code)
// 	requireBodyMatchAccount(t, recorder.Body, account)
// }

// func createRandomAccount() db.Account {
// 	return db.Account{
// 		ID:       utils.RandomInt(1, 1000),
// 		Owner:    utils.RandomOwner(),
// 		Balance:  utils.RandomBalance(),
// 		Currency: utils.RandomCurrency(),
// 	}
// }

// func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
// 	data, err := ioutil.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotAccount db.Account
// 	err = json.Unmarshal(data, &gotAccount)
// 	require.NoError(t, err)
// 	require.Equal(t, account, gotAccount)
// }
