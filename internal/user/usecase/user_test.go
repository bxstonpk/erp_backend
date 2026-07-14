package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/usecase"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var captured *entity.User
		repo := &fakeRepository{
			CreateUserFn: func(user *entity.User) error {
				captured = user
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.CreateUserRequest{
			Username: "jdoe",
			Email:    "jdoe@example.com",
			Password: "already-hashed",
			FullName: "John Doe",
			Phone:    "0800000000",
			Status:   "active",
		}

		got, err := uc.CreateUser(req)
		if err != nil {
			t.Fatalf("CreateUser() error = %v", err)
		}
		if got.ID == "" {
			t.Fatal("CreateUser() did not generate an ID")
		}
		if captured != got {
			t.Fatal("CreateUser() did not persist the returned user")
		}
		if got.Username != req.Username || got.Email != req.Email || got.Password != req.Password ||
			got.FullName != req.FullName || got.Phone != req.Phone || got.Status != req.Status {
			t.Fatalf("CreateUser() mismatched fields: %+v", got)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			CreateUserFn: func(user *entity.User) error { return wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.CreateUser(dto.CreateUserRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("CreateUser() error = %v, want %v", err, wantErr)
		}
	})
}

func TestUserUsecase_GetUser(t *testing.T) {
	want := &entity.User{ID: "u-1", Username: "jdoe"}
	repo := &fakeRepository{
		GetUserByUUIDFn: func(userUUID string) (*entity.User, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetUserByUUID() userUUID = %q, want %q", userUUID, "u-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetUser("u-1")
	if err != nil {
		t.Fatalf("GetUser() error = %v", err)
	}
	if got != want {
		t.Fatalf("GetUser() = %v, want %v", got, want)
	}
}

func TestUserUsecase_GetUserByUsername(t *testing.T) {
	want := &entity.User{ID: "u-1", Username: "jdoe"}
	repo := &fakeRepository{
		GetUserByUsernameFn: func(username string) (*entity.User, error) {
			if username != "jdoe" {
				t.Fatalf("GetUserByUsername() username = %q, want %q", username, "jdoe")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetUserByUsername("jdoe")
	if err != nil {
		t.Fatalf("GetUserByUsername() error = %v", err)
	}
	if got != want {
		t.Fatalf("GetUserByUsername() = %v, want %v", got, want)
	}
}

func TestUserUsecase_ListUsers(t *testing.T) {
	want := []*entity.User{{ID: "u-1"}, {ID: "u-2"}}
	repo := &fakeRepository{
		GetAllUsersFn: func() ([]*entity.User, error) { return want, nil },
	}
	uc := usecase.New(repo)

	got, err := uc.ListUsers()
	if err != nil {
		t.Fatalf("ListUsers() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("ListUsers() len = %d, want %d", len(got), len(want))
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	t.Run("keeps password when none supplied", func(t *testing.T) {
		existing := &entity.User{
			ID: "u-1", Username: "old", Email: "old@example.com",
			Password: "kept", FullName: "Old Name", Phone: "111", Status: "inactive",
		}
		var saved *entity.User
		repo := &fakeRepository{
			GetUserByUUIDFn: func(userUUID string) (*entity.User, error) { return existing, nil },
			UpdateUserFn: func(user *entity.User) error {
				saved = user
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.UpdateUserRequest{Username: "new", Email: "new@example.com", FullName: "New Name", Phone: "222", Status: "active"}
		got, err := uc.UpdateUser("u-1", req)
		if err != nil {
			t.Fatalf("UpdateUser() error = %v", err)
		}
		if got.Password != "kept" {
			t.Fatalf("UpdateUser() overwrote password without a new one: %q", got.Password)
		}
		if got.Username != "new" || got.Email != "new@example.com" || got.FullName != "New Name" ||
			got.Phone != "222" || got.Status != "active" {
			t.Fatalf("UpdateUser() mismatched fields: %+v", got)
		}
		if saved != got {
			t.Fatal("UpdateUser() did not persist the updated user")
		}
	})

	t.Run("replaces password when supplied", func(t *testing.T) {
		existing := &entity.User{ID: "u-1", Password: "old-hash"}
		repo := &fakeRepository{
			GetUserByUUIDFn: func(userUUID string) (*entity.User, error) { return existing, nil },
			UpdateUserFn:    func(user *entity.User) error { return nil },
		}
		uc := usecase.New(repo)

		got, err := uc.UpdateUser("u-1", dto.UpdateUserRequest{Password: "new-hash"})
		if err != nil {
			t.Fatalf("UpdateUser() error = %v", err)
		}
		if got.Password != "new-hash" {
			t.Fatalf("UpdateUser() Password = %q, want %q", got.Password, "new-hash")
		}
	})

	t.Run("get error", func(t *testing.T) {
		wantErr := errors.New("not found")
		repo := &fakeRepository{
			GetUserByUUIDFn: func(userUUID string) (*entity.User, error) { return nil, wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.UpdateUser("missing", dto.UpdateUserRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("UpdateUser() error = %v, want %v", err, wantErr)
		}
	})

	t.Run("update error", func(t *testing.T) {
		wantErr := errors.New("db down")
		existing := &entity.User{ID: "u-1"}
		repo := &fakeRepository{
			GetUserByUUIDFn: func(userUUID string) (*entity.User, error) { return existing, nil },
			UpdateUserFn:    func(user *entity.User) error { return wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.UpdateUser("u-1", dto.UpdateUserRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("UpdateUser() error = %v, want %v", err, wantErr)
		}
	})
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeleteUserFn: func(userUUID string) error {
				got = userUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUser("u-1"); err != nil {
			t.Fatalf("DeleteUser() error = %v", err)
		}
		if got != "u-1" {
			t.Fatalf("DeleteUser() userUUID = %q, want %q", got, "u-1")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			DeleteUserFn: func(userUUID string) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUser("u-1"); !errors.Is(err, wantErr) {
			t.Fatalf("DeleteUser() error = %v, want %v", err, wantErr)
		}
	})
}
