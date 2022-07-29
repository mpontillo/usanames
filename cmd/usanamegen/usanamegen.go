package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alecthomas/kong"

	"github.com/mpontillo/usanames"
)

var cli struct{
	Naive naiveCmd `kong:"cmd,help='Naively generate a random name'"`
}

type naiveCmd struct {
	Count int `kong:"default='1'"`
}

func (r *naiveCmd) Run() error {
	if r.Count < 1 {
		return fmt.Errorf("count must be equal to or greater than 1")
	}
	for i := 1 ; i <= r.Count ; i++ {
		fmt.Println(usanames.NaiveRandomName())
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	kongContext := kong.Parse(&cli)
	err := kongContext.Run()
	kongContext.FatalIfErrorf(err)
}
