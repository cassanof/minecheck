package main

import (
	tm "github.com/buger/goterm"

	"github.com/elleven11/minecheck/cryptonote"
	"github.com/elleven11/minecheck/draw"
	"github.com/elleven11/minecheck/twominers"
)

func main() {
	cn, err := cryptonote.GetStats("put username") // username goes here
	twm, err := twominers.GetStats("put wallet")     // wallet goes here
	if err != nil {
		panic(err)
	}

	boxes := []*tm.Box{
		draw.MakeCryptonoteBox(cn),
		draw.MakeTwominersBox(twm),
	}

	draw.DrawBoxes(boxes)
}
