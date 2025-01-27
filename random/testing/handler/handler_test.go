package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	users := map[string]User{
		"u1": {ID: "u1", FirstName: "John", LastName: "Doe"},
	}

	tests := []struct {
		name     string
		endpoint string
		want     struct {
			status   int
			response string
		}
	}{
		{
			name:     "Status 200",
			endpoint: "/users?user_id=u1",
			want: struct {
				status   int
				response string
			}{
				status:   200,
				response: `{"ID":"u1","FirstName":"John","LastName":"Doe"}`,
			},
		},
		{
			name:     "Status 404",
			endpoint: "/users?user_id=u3",
			want: struct {
				status   int
				response string
			}{
				status:   404,
				response: "user not found\n",
			},
		},
		{
			name:     "Status 400",
			endpoint: "/users",
			want: struct {
				status   int
				response string
			}{
				status:   400,
				response: "user_id is empty\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.endpoint, nil)
			w := httptest.NewRecorder()
			UserViewHandler(users)(w, request)
			res := w.Result()

			assert.Equal(t, res.StatusCode, tt.want.status)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, tt.want.response, string(resBody))
		})
	}
}
