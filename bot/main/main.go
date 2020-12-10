package main

import (
	"github.com/mrNobody95/adminyar/bot/handler"
	"github.com/mrNobody95/adminyar/bot/noSql"
)

func main() {
	noSql.Run()
	handler.StartBotAPI()

	defer noSql.Close()
}
