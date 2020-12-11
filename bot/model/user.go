package model

import (
	"errors"
	"fmt"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mitchellh/mapstructure"
	"github.com/mrNobody95/adminyar/bot/noSql"
)

type User struct {
	botAPI.User
	ChatId int64
	Status UserStatus
}

type UserStatus int

const (
	Free           = 0
	AddChannelId   = 1
	AddChannelName = 2
)

func CheckUser(user *botAPI.User) error {
	query := fmt.Sprintf("MERGE (u:user{id:%d, first_name:'%s', last_name:'%s', username:'%s'}) RETURN COUNT(u)",
		user.ID, user.FirstName, user.LastName, user.UserName)
	fmt.Println("query is:", query)
	session, err := noSql.GetSession()
	if err != nil {
		return err
	}
	fmt.Println("session got")
	result, err := session.Run(query, nil)
	if err != nil {
		return err
	}
	if result.Err() != nil {
		return result.Err()
	}
	fmt.Println("check next")
	if result.Next() {
		u := 0
		err = mapstructure.Decode(Parse(result.Record()), &u)
		if err != nil {
			return err
		}
		if u == 1 {
			return nil
		}
	}
	return errors.New("user not found")
}

func ChangeUserStatus(userId int, status UserStatus) error {

}

func LoadUser(userId int) (User, error) {

}
