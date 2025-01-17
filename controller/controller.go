package controller

import (
	"tdgames/logic"
	"tdgames/storage"
)

// User interface defines methods for User Controller
type Group interface {
	GroupRegister(chatID int64) error   // Performs user registration
	UnGroupregister(chatID int64) error // Performs user deregistration
}

// User interface defines methods for User Controller
type User interface {
	Register(userid int64, chatID int64) error                             // Performs user registration
	Unregister(chatID int64) error                                         // Performs user deregistration
	GetUsersByPublication(pub *logic.Publication) ([]logic.User, error)    // Returns owner of publication
	Sign(chatID int, chatid int64, sign int) (int64, bool)                 // 签到
	Balance(userid, chatID int64) (*logic.Leaderboard, error)              // 余额
	Transfer(userID string, targetid string, payload int64) (int64, error) //税率 转账
	Deposit(userID int, payload int64) (int64, error)                      //存钱													//存款
	DrawMoney(userID int, payload int64) (int64, error)                    //取款
}

// Subscription interface defines methods for Publication Controller
type Subscription interface {
	AddNew(chatID int64, request string) error // Adds new subscription to user with publication
	Create(chatID int64, request string) error
	Remove(chatID int64, request string) error                 // Removes existing sybscription from user
	Update(chatID int64, request string) error                 // Updates selected subscription
	GetSubsByChatID(chatID int64) ([]logic.Publication, error) // Returns all user's subs
	GetAllSubs() []logic.Publication                           // Returns all publications
	GetAllDefaultSubs() []logic.Publication
	RemoveDefault(chatID int64, request string) error
	Subscribe(chatID int64, request string) error
}

// Info interface definces methods for Info Controller
type Info interface {
	GetLastTimestamp() uint64    // Returns time of the latest post
	SetLastTimestamp(tsp uint64) // Sets time of the latest post
}

// Controller struct is used to access database
type Controller struct {
	User
	Subscription
	Info
	Group
}

// NewController constructor of Controller
func NewController(stg *storage.Storage) *Controller {
	return &Controller{
		User:         NewUserController(stg),
		Subscription: NewSubscriptionController(stg),
		Info:         NewInfoController(stg),
		Group:        NewGroupController(stg),
	}
}
