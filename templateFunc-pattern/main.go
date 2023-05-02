package main

import "fmt"

//定义一个抽象的游戏接口
type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
}

//实现一个游戏抽象基类,用于实现模板方法
type GameBase struct {
	Game
}

func (g *GameBase) Play(){
	g.Initialize()
	g.StartPlay()
	g.EndPlay()
}

//实现具体的游戏类: 篮球
type BasketballGame struct {
	GameBase
}

func(b *BasketballGame)Initialize(){
	fmt.Println("Basketball game initialze")
}

func(b *BasketballGame)StartPlay(){
	fmt.Println("Basketball game start")
}

func(b *BasketballGame)EndPlay(){
	fmt.Println("Basketball game end")
}

//实现具体的游戏类: 足球
type FootballGame struct {
	GameBase
}

func(b *FootballGame)Initialize(){
	fmt.Println("FootballGame game initialze")
}

func(b *FootballGame)StartPlay(){
	fmt.Println("FootballGame game start")
}

func(b *FootballGame)EndPlay(){
	fmt.Println("FootballGame game end")
}

func main(){
	basketballgame := &BasketballGame{}
	basketballgame.Game = basketballgame
	basketballgame.GameBase.Play()

	fmt.Println()

	footballgame := &FootballGame{}
	footballgame.Game = footballgame
	footballgame.GameBase.Play()
}