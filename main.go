package main

import (
	"github.com/joho/godotenv"
	"github.com/t4ku/notion-logseq/cmd"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("no .env file")
	}
	cmd.Execute()
}
