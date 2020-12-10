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
}

func CheckUser(user *botAPI.User) error {
	query := fmt.Sprintf("MERGE (u:user{id:%d, first_name:'%s', last_name:'%s', username:'%s'}) RETURN COUNT(u)",
		user.ID, user.FirstName, user.LastName, user.UserName)
	session, err := noSql.GetSession()
	if err != nil {
		return err
	}
	result, err := session.Run(query, nil)
	if err != nil {
		return err
	}
	if result.Err() != nil {
		return result.Err()
	}
	if result.Next() {
		u := 0
		err = mapstructure.Decode(Parse(result.Record()), &u)
		if u == 1 {
			return nil
		}
	}
	return errors.New("user not found")
}
