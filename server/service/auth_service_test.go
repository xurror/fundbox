package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"getting-to-go/config"
	"getting-to-go/service/mocks"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUnmarshalJson(t *testing.T) {
	logger := logrus.New()
	config := &config.AppConfig{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	jsonData := `{"key": "value"}`
	expected := map[string]interface{}{"key": "value"}

	result := authService.unmarshalJson([]byte(jsonData))
	assert.Equal(t, expected, result)
}

func TestUnmarshalResponse(t *testing.T) {
	logger := logrus.New()
	config := &config.AppConfig{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	response := httptest.NewRecorder()
	response.Body = bytes.NewBufferString(`{"key": "value"}`)
	expected := map[string]interface{}{"key": "value"}

	result := authService.unmarshalResponse(response.Result())
	assert.Equal(t, expected, result)
}

func TestExecuteRequest(t *testing.T) {
	logger := logrus.New()
	config := &config.AppConfig{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)
	expected := map[string]interface{}{"key": "value"}

	result := authService.executeRequest(req)
	assert.Equal(t, expected, result)
}

func TestAuthorize(t *testing.T) {
	logger := logrus.New()
	config := &config.AppConfig{
		Auth0: config.Auth0Config{
			Domain:   "example.com",
			Audience: "exampleAudience",
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	tokenString := "validTokenString"

	userService.EXPECT().GetUserByAuth0Id(gomock.Any(), "auth0Id").Return(nil, nil)
	userService.EXPECT().CreateUser(gomock.Any(), "auth0Id").Return(nil, nil)

	_, err := authService.Authorize(ctx, tokenString)
	assert.NoError(t, err)
}
func BenchmarkUnmarshalJson(b *testing.B) {
	logger := logrus.New()
	config := &config.AppConfig{}
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	jsonData := []byte(`{"key": "value"}`)

	for i := 0; i < b.N; i++ {
		authService.unmarshalJson(jsonData)
	}
}

func BenchmarkUnmarshalResponse(b *testing.B) {
	logger := logrus.New()
	config := &config.AppConfig{}
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	response := httptest.NewRecorder()
	response.Body = bytes.NewBufferString(`{"key": "value"}`)

	for i := 0; i < b.N; i++ {
		authService.unmarshalResponse(response.Result())
	}
}

func BenchmarkExecuteRequest(b *testing.B) {
	logger := logrus.New()
	config := &config.AppConfig{}
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	authService := NewAuthService(logger, config, userService)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	req, _ := http.NewRequest("GET", server.URL, nil)

	for i := 0; i < b.N; i++ {
		authService.executeRequest(req)
	}
}
