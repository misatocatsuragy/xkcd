package main

import (
	"flag"
	"log"
	"strings"

	"xkcd/index"
)

func main() {
	indexPath := flag.String("f", "", "Specify file for index")
	isCreate := flag.Bool("c", false, "Create index file specidied with -f option, don't perform search")

	flag.Parse()

	if *isCreate {
		err := index.CreateIndex(*indexPath)
		if err != nil {
			log.Fatalf("xkcd: %v", err)
		}
	} else {
		terms := strings.Join(flag.Args(), " ")
		err := index.SearchComics(terms, *indexPath)
		if err != nil {
			log.Fatalf("xkcd: %v", err)
		}
	}
}
