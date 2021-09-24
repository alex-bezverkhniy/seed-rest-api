package user_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"seed-rest-api/internal/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {
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
			wantBody:   "{\"status\":\"success\",\"message\":\"\",\"data\":[{\"id\":1,\"name\":\"MockedUser\",\"address\":\"TestAddress\",\"created\":123,\"modified\":321}]}",
		},
		{
			name:       "Get user by ID",
			method:     "GET",
			route:      "/api/v1/users/1",
			body:       nil,
			wantErr:    false,
			wantStatus: "success",
			wantCode:   200,
			wantBody:   "{\"status\":\"success\",\"message\":\"\",\"data\":{\"id\":1,\"name\":\"MockedUser\",\"address\":\"TestAddress\",\"created\":123,\"modified\":321}}",
		},
		{
			name:       "Get user by wrong ID",
			method:     "GET",
			route:      "/api/v1/users/WrongID",
			body:       nil,
			wantErr:    false,
			wantStatus: "faild",
			wantCode:   400,
			wantBody:   "{\"status\":\"fail\",\"message\":\"Please specify a valid user ID\",\"data\":null}",
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

		res, err := app.Test(req, -1)

		assert.Equalf(t, tt.wantErr, err != nil, tt.name)

		if tt.wantErr {
			continue
		}

		assert.Equalf(t, tt.wantCode, res.StatusCode, tt.name)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, tt.name)

		assert.Equalf(t, tt.wantBody, string(body), tt.name)

	}
}