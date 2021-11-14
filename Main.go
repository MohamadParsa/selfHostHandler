package main

import(
	"HTTP/hostHandler"
)

func main()  {
	setting:=hostHandler.Settings{
		RootWebPath:"Views",
		DefaultPageName:"/index.html",
		NotFindePagePath:"/ErrorPages/PageNotFound.html"}
		hostHandler.RunHostHandler("/","8099",setting)

}