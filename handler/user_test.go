package handler_test

import (
	"bytes"
	"encoding/json"
	"kuwa72/sample-gorm-txdb-testing/handler"
	"kuwa72/sample-gorm-txdb-testing/testutil"
	"kuwa72/sample-gorm-txdb-testing/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler_CreateUser(t *testing.T) {
	db, _ := testutil.NewTestDB("TestUserHandler_CreateUser")
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	db.Create(
		&usecase.User{
			Name:     "testname2",
			Email:    "test2@example.com",
			Password: "testpass2",
		},
	)

	h := &handler.UserHandler{
		DB: db,
	}

	router := handler.SetupRouter()
	router = h.CreateUser(router)

	tests := []struct {
		name   string
		req    handler.CreateUserRequest
		status int
		want   usecase.User
	}{
		{
			name: "create success",
			req: handler.CreateUserRequest{
				Name:     "testname1",
				Email:    "test1@example.com",
				Password: "testpass1",
			},
			status: http.StatusOK,
			want: usecase.User{
				ID:       1,
				Name:     "testname1",
				Email:    "test1@example.com",
				Password: "testpass1",
			},
		},
		{
			name: "exists",
			req: handler.CreateUserRequest{
				Name:     "testname2",
				Email:    "test2@example.com",
				Password: "testpass2",
			},
			status: http.StatusOK,
			want:   usecase.User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			w := httptest.NewRecorder()
			reqJson, _ := json.Marshal(tt.req)
			req, _ := http.NewRequest("POST", "/user/add", bytes.NewBuffer(reqJson))
			router.ServeHTTP(w, req)

			assert.Equal(tt.status, w.Code)
			if tt.status == http.StatusOK {
				var got usecase.User
				json.Unmarshal(w.Body.Bytes(), &got)
				assert.Equal(tt.want.Name, got.Name)
				assert.Equal(tt.want.Email, got.Email)
				assert.Equal(tt.want.Password, got.Password)

				assert.NoError(db.Where("email = ?", tt.req.Email).First(&got).Error)
				assert.Equal(tt.req.Name, got.Name)
				assert.Equal(tt.req.Email, got.Email)
				assert.Equal(tt.req.Password, got.Password)
			}
		})
	}
}
