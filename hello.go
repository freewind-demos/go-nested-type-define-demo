package main

import "fmt"

type Site struct {
	name string
	categories []struct {
		name        string
		description string
	}
}

func main() {
	uglyInit()
	uglyInit2()
}

func uglyInit() {
	fmt.Println("------------- uglyInit --------------")
	site := Site{
		name: "go",
		categories: [] struct {
			name        string
			description string
		}{{
			name:        "go-demos",
			description: "all kinds of go demos",
		}},
	}
	fmt.Println(site)
}

func uglyInit2() {
	fmt.Println("------------- uglyInit2 --------------")
	site := Site{
		name: "go",
	}
	site.categories = make([]struct {
		name        string
		description string
	}, 1)
	site.categories[0].name = "go-demos"
	site.categories[0].description = "all kinds of go demos"
	fmt.Println(site)
}
