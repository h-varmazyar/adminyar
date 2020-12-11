package main

import (
	"github.com/mrNobody95/adminyar/bot/handler"
	"github.com/mrNobody95/adminyar/bot/noSql"
	"log"
)

func main() {
	err := noSql.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = handler.StartBotAPI()
	if err != nil {
		log.Fatal(err)
	}
	defer noSql.Close()
}
