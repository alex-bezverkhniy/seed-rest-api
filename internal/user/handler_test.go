package user_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"seed-rest-api/internal/infrastructure"
	"seed-rest-api/internal/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockedUser = &user.User{ID: 1, Name: "MockedUser", Address: "TestAddress"}

func TestUserHandler(t *testing.T) {
	mockedBodyBytes, err := json.Marshal(mockedUser)
	if err != nil {
		t.Fatalf(err.Error())
	}

	mockedBody := bytes.NewBuffer(mockedBodyBytes)

	tests := []struct {
		name       string
		method     string
		route      string
		body       io.Reader
		wantErr    bool
		wantStatus string
		wantCode   int
		wantBody   string
	}{
		{
			name:       "Get all users",
			method:     "GET",
			route:      "/api/v1/users",
			body:       nil,
			wantErr:    false,
			wantStatus: "success",
			wantCode:   200,
			wantBody:   `{"status":"success","data":[{"id":1,"name":"MockedUser","address":"TestAddress","created":123,"modified":321}]}`,
		},
		{
			name:       "Get user by ID",
			method:     "GET",
			route:      "/api/v1/users/1",
			body:       nil,
			wantErr:    false,
			wantStatus: "success",
			wantCode:   200,
			wantBody:   `{"status":"success","data":{"id":1,"name":"MockedUser","address":"TestAddress","created":123,"modified":321}}`,
		},
		{
			name:       "Get user by wrong ID",
			method:     "GET",
			route:      "/api/v1/users/WrongID",
			body:       nil,
			wantErr:    false,
			wantStatus: "faild",
			wantCode:   400,
			wantBody:   `{"status":"fail","message":"Please specify a valid user ID"}`,
		},
		{
			name:       "Get non existing user by ID",
			method:     "GET",
			route:      "/api/v1/users/10",
			body:       nil,
			wantErr:    false,
			wantStatus: "faild",
			wantCode:   404,
			wantBody:   `{"status":"fail","message":"User with specified ID is not found"}`,
		},
		{
			name:       "Create new user",
			method:     "POST",
			route:      "/api/v1/users",
			body:       mockedBody,
			wantErr:    false,
			wantStatus: "sucsess",
			wantCode:   202,
			wantBody:   `{"status":"success","message":"User created"}`,
		},
	}

	app := infrastructure.SetupMock()

	for _, tt := range tests {

		// Create request
		req, _ := http.NewRequest(
			tt.method,
			tt.route,
			tt.body,
		)
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		res, err := app.Test(req, -1)

		assert.Equalf(t, tt.wantErr, err != nil, tt.name)

		if tt.wantErr {
			continue
		}

		assert.Equalf(t, tt.wantCode, res.StatusCode, tt.name)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, tt.name)

		regexpCrd, err := regexp.Compile(`(.*\"created\":)([\d]*)(.*)`)
		if err != nil {
			t.Fatal(err)

		}

		gotBody := string(body)
		tt.wantBody = regexpCrd.ReplaceAllString(tt.wantBody, `$1"000"$3`)
		gotBody = regexpCrd.ReplaceAllString(gotBody, `$1"000"$3`)

		regexpMdf, err := regexp.Compile(`(.*\"modified\":)([\d]*)(.*)`)
		if err != nil {
			t.Fatal(err)
		}
		tt.wantBody = regexpMdf.ReplaceAllString(tt.wantBody, `$1"000"$3`)
		gotBody = regexpMdf.ReplaceAllString(gotBody, `$1"000"$3`)

		assert.Equalf(t, tt.wantBody, gotBody, tt.name)

	}
}
