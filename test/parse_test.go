package main

import (
	"log"
	"os"
	"testing"

	"github.com/orsetii/defuse/cmd/parse"
)

func TestParse(t *testing.T) {
	f, err := os.Open("../demos/natus-vincere-vs-astralis-m1-nuke.dem")
	if err != nil {
		log.Println(err)
	}
	parse.ParseDemo(f, true)

}
