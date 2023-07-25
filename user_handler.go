package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/wimspaargaren/echo/api"
)

type user struct {
	Email     string
	FirstName string
	ID        uuid.UUID
	LastName  string
}

func userFromAPIUser(u api.User) *user {
	return &user{
		Email:     u.Email,
		FirstName: u.FirstName,
		ID:        u.ID,
		LastName:  u.LastName,
	}
}

func userToAPIUser(u *user) api.User {
	return api.User{
		Email:     u.Email,
		FirstName: u.FirstName,
		ID:        u.ID,
		LastName:  u.LastName,
	}
}

type userHandler struct {
	userDB map[uuid.UUID]*user
}

func newUserHandler() *userHandler {
	return &userHandler{
		userDB: map[uuid.UUID]*user{},
	}
}

func (t *userHandler) PostUser(ctx echo.Context) error {
	apiUser := api.User{}

	err := json.NewDecoder(ctx.Request().Body).Decode(&apiUser)
	if err != nil {
		fmt.Println("unable to decode create user request", err)
		return ctx.JSON(http.StatusBadRequest, api.ErrMsg{
			Message: "Unable to decode request body",
		})
	}

	user := userFromAPIUser(apiUser)
	user.ID = uuid.New()
	t.userDB[user.ID] = user

	return ctx.JSON(http.StatusCreated, userToAPIUser(user))
}

func (t *userHandler) GetUserId(ctx echo.Context, id uuid.UUID) error {
	user, ok := t.userDB[id]
	if !ok {
		return ctx.JSON(http.StatusBadRequest, api.ErrMsg{
			Message: "User not found",
		})
	}
	return ctx.JSON(http.StatusOK, userToAPIUser(user))
}

func (t *userHandler) PutUserId(ctx echo.Context, id uuid.UUID) error {
	apiUser := api.User{}

	err := json.NewDecoder(ctx.Request().Body).Decode(&apiUser)
	if err != nil {
		fmt.Println("unable to decode update user request", err)
		return ctx.JSON(http.StatusBadRequest, api.ErrMsg{
			Message: "Unable to decode request body",
		})
	}

	user := userFromAPIUser(apiUser)
	userDB, ok := t.userDB[id]
	if !ok {
		return ctx.JSON(http.StatusBadRequest, api.ErrMsg{
			Message: "User not found",
		})
	}
	userDB.Email = user.Email
	userDB.FirstName = user.FirstName
	userDB.LastName = user.LastName
	t.userDB[id] = userDB

	return ctx.JSON(http.StatusOK, userToAPIUser(userDB))
}

func (t *userHandler) GetUser(ctx echo.Context) error {
	res := []api.User{}

	for _, v := range t.userDB {
		res = append(res, userToAPIUser(v))
	}

	return ctx.JSON(http.StatusOK, res)
}

func (t *userHandler) DeleteUserId(ctx echo.Context, id uuid.UUID) error {
	delete(t.userDB, id)
	return ctx.JSON(http.StatusOK, struct{}{})
}
