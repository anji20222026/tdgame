package controller

import (
	"github.com/aoyako/telegram_2ch_res_bot/logic"
	"github.com/aoyako/telegram_2ch_res_bot/storage"
)

// InfoController is an implementation of controller.Info
type GroupController struct {
	stg *storage.Storage
}

// NewUserController constructor of UserController struct
func NewGroupController(stg *storage.Storage) *GroupController {
	return &GroupController{stg: stg}
}

// NewUserController constructor of UserController struct
func (ucon *GroupController) GroupRegister(chatID int64) error {

	user := &logic.Group{
		Groupid: chatID,
	}
	return ucon.stg.GroupRegister(user)

}

// NewUserController constructor of UserController struct
func (ucon *GroupController) UnGroupregister(chatID int64) error {
	user := &logic.Group{
		Groupid: chatID,
	}
	return ucon.stg.UnGroupRegister(user)
}
