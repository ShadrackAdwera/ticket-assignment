package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/ShadrackAdwera/ticket-assignment/db/mocks"
	db "github.com/ShadrackAdwera/ticket-assignment/db/sqlc"
	"github.com/ShadrackAdwera/ticket-assignment/token"
	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func addAuthorization(t *testing.T,
	request *http.Request,
	tokenMaker token.TokenMaker,
	authorizationType string,
	username string,
	id int64,
	duration time.Duration) {
	token, err := tokenMaker.CreateToken(username, id, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func generateAgent() db.Agent {
	return db.Agent{
		ID:        rand.Int63n(100),
		Name:      utils.GenerateRandomString(),
		Status:    utils.GetAgentStatus(),
		CreatedAt: time.Now(),
	}
}

func compareAgentCreatedWithAgentRecorded(t *testing.T, body *bytes.Buffer, agent db.Agent) {
	readData, err := io.ReadAll(body)
	require.NoError(t, err)

	var resData db.Agent

	err = json.Unmarshal(readData, &resData)

	require.NoError(t, err)
	require.Equal(t, agent.ID, resData.ID)
	require.Equal(t, agent.Status, resData.Status)
	require.WithinDuration(t, agent.CreatedAt, resData.CreatedAt, time.Duration(time.Second))
}

func TestGetAgentEndpoint(t *testing.T) {
	agent := generateAgent()

	testCases := []struct {
		name          string
		status        string
		id            int64
		setUpAuth     func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker)
		buildStubs    func(mockdb *mockdb.MockTxStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			status: "ACTIVE",
			id:     agent.ID,
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				addAuthorization(t, request, tokenMaker, authTypeBearer, utils.RandomString(6), utils.RandomInteger(1, 100), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().GetAgent(gomock.Any(), gomock.Eq(agent.ID)).Times(1).Return(agent, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareAgentCreatedWithAgentRecorded(t, recorder.Body, agent)
			},
		},
		{
			name:   "AgentNotFound",
			status: "ACTIVE",
			id:     agent.ID,
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				addAuthorization(t, request, tokenMaker, authTypeBearer, utils.RandomString(6), utils.RandomInteger(1, 100), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().GetAgent(gomock.Any(), gomock.Eq(agent.ID)).Times(1).Return(db.Agent{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "ErrInternalServerError",
			status: "ACTIVE",
			id:     agent.ID,
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				addAuthorization(t, request, tokenMaker, authTypeBearer, utils.RandomString(6), utils.RandomInteger(1, 100), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().GetAgent(gomock.Any(), gomock.Eq(agent.ID)).Times(1).Return(db.Agent{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "ErrBadRequest",
			status: "ACTIVE",
			id:     0,
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				addAuthorization(t, request, tokenMaker, authTypeBearer, utils.RandomString(6), utils.RandomInteger(1, 100), time.Minute)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().GetAgent(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)
			defer ctlr.Finish()

			store := mockdb.NewMockTxStore(ctlr)
			testCase.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/agents/%d", testCase.id)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			testCase.setUpAuth(t, req, server.tokenMaker)

			server.router.ServeHTTP(recorder, req)
			testCase.checkResponse(t, recorder)
		})

	}

}
