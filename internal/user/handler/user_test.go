package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/handler"

	"github.com/gorilla/mux"
)

func newRequestWithVars(method, target string, body []byte, vars map[string]string) *http.Request {
	req := httptest.NewRequest(method, target, bytes.NewReader(body))
	if vars != nil {
		req = mux.SetURLVars(req, vars)
	}
	return req
}

func TestHandler_CreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			CreateUserFn: func(req dto.CreateUserRequest) (*entity.User, error) {
				return &entity.User{ID: "u-1", Username: req.Username}, nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CreateUserRequest{Username: "jdoe"})
		rec := httptest.NewRecorder()
		h.CreateUser(rec, newRequestWithVars(http.MethodPost, "/users", body, nil))

		if rec.Code != http.StatusCreated {
			t.Fatalf("CreateUser() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
		}

		var got dto.UserResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if got.ID != "u-1" || got.Username != "jdoe" {
			t.Fatalf("CreateUser() response = %+v", got)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.CreateUser(rec, newRequestWithVars(http.MethodPost, "/users", []byte("{"), nil))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("CreateUser() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			CreateUserFn: func(req dto.CreateUserRequest) (*entity.User, error) {
				return nil, errors.New("boom")
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CreateUserRequest{})
		rec := httptest.NewRecorder()
		h.CreateUser(rec, newRequestWithVars(http.MethodPost, "/users", body, nil))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("CreateUser() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_GetUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			GetUserFn: func(userUUID string) (*entity.User, error) {
				if userUUID != "u-1" {
					t.Fatalf("GetUser() userUUID = %q, want %q", userUUID, "u-1")
				}
				return &entity.User{ID: "u-1", Username: "jdoe"}, nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.GetUser(rec, newRequestWithVars(http.MethodGet, "/users/u-1", nil, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("GetUser() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
		}
	})

	t.Run("not found", func(t *testing.T) {
		uc := &fakeUsecase{
			GetUserFn: func(userUUID string) (*entity.User, error) { return nil, errors.New("not found") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.GetUser(rec, newRequestWithVars(http.MethodGet, "/users/missing", nil, map[string]string{"userUUID": "missing"}))

		if rec.Code != http.StatusNotFound {
			t.Fatalf("GetUser() status = %d, want %d", rec.Code, http.StatusNotFound)
		}
	})
}

func TestHandler_GetUserByUsername(t *testing.T) {
	uc := &fakeUsecase{
		GetUserByUsernameFn: func(username string) (*entity.User, error) {
			if username != "jdoe" {
				t.Fatalf("GetUserByUsername() username = %q, want %q", username, "jdoe")
			}
			return &entity.User{ID: "u-1", Username: "jdoe"}, nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.GetUserByUsername(rec, newRequestWithVars(http.MethodGet, "/users/username/jdoe", nil, map[string]string{"username": "jdoe"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("GetUserByUsername() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
}

func TestHandler_ListUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			ListUsersFn: func() ([]*entity.User, error) {
				return []*entity.User{{ID: "u-1"}, {ID: "u-2"}}, nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.ListUsers(rec, newRequestWithVars(http.MethodGet, "/users", nil, nil))

		if rec.Code != http.StatusOK {
			t.Fatalf("ListUsers() status = %d, want %d", rec.Code, http.StatusOK)
		}

		var got []dto.UserResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if len(got) != 2 {
			t.Fatalf("ListUsers() response len = %d, want 2", len(got))
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			ListUsersFn: func() ([]*entity.User, error) { return nil, errors.New("boom") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.ListUsers(rec, newRequestWithVars(http.MethodGet, "/users", nil, nil))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("ListUsers() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			UpdateUserFn: func(userUUID string, req dto.UpdateUserRequest) (*entity.User, error) {
				if userUUID != "u-1" {
					t.Fatalf("UpdateUser() userUUID = %q, want %q", userUUID, "u-1")
				}
				return &entity.User{ID: userUUID, Username: req.Username}, nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.UpdateUserRequest{Username: "new-name"})
		rec := httptest.NewRecorder()
		h.UpdateUser(rec, newRequestWithVars(http.MethodPut, "/users/u-1", body, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("UpdateUser() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.UpdateUser(rec, newRequestWithVars(http.MethodPut, "/users/u-1", []byte("{"), map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("UpdateUser() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestHandler_DeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		uc := &fakeUsecase{
			DeleteUserFn: func(userUUID string) error {
				got = userUUID
				return nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteUser(rec, newRequestWithVars(http.MethodDelete, "/users/u-1", nil, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusNoContent {
			t.Fatalf("DeleteUser() status = %d, want %d", rec.Code, http.StatusNoContent)
		}
		if got != "u-1" {
			t.Fatalf("DeleteUser() userUUID = %q, want %q", got, "u-1")
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			DeleteUserFn: func(userUUID string) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteUser(rec, newRequestWithVars(http.MethodDelete, "/users/u-1", nil, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("DeleteUser() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}
