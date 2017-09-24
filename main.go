package main

import (
	"./anka"
	"./bird"
)

type Profile struct {
	Id string
	User string
	Path string
	Type string
	QueryList map[string][]string
}

func HelloWorld() {
	bird.Message("Merhaba Dünya")
}

func ProfilePage() {

	Path := bird.Path
	Type := bird.Type

	ID := ""
	Name := ""

	if Type == "GET" && len(bird.QueryList) > 0 {
		ID = bird.Query("id")
		Name = bird.Query("name")
	}

	

	p := Profile{Id: ID, User: Name, Path: Path, Type: Type, QueryList: bird.QueryList}

	bird.Template("profile", p)
}

func main() {
	
	anka.StaticDir("static")

	bird.NotFound("Page not found!", "error/404.html", "yes")

	bird.Crow("/", HelloWorld)
	bird.Crow("/profile", ProfilePage)

	bird.Serve()

}