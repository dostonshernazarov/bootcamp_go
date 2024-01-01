package main

import (
	"fmt"
)

type Game struct {
	On    bool
	Ammo  int
	Power int
}
func (g *Game) Shoot() bool {
	
	if g.On == true &&  g.Ammo > 0 {
		g.Ammo --
		return true
    } else {
        return false
    }
}

func (g *Game) RideBike() bool {
	
	if g.On == true && g.Power > 0 {
		g.Power --
		return true
    } else {
        return false
    }
}

func main(){

	fmt.Println("Hello")

	game := Game{
		On: true,
		Ammo: 2,
		Power: 2,
	}

	for i:=0;i<5;i++{

		

		fmt.Println(game.RideBike())
	}


}