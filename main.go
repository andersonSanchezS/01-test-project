package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"dce37781-cce1-4408-add9-b4ab65536acd": "Corona",
	"e243c11f-0420-4404-9aa9-4185ceb573d4": "Budweiser",
	"1110e378-6f4b-43d7-8c49-256fef70251f": "Heineken",
}

func main() {
	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("please prompt a command")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		id := beersCmd.String("id", "", "find a beer by id")
		beersCmd.Parse(os.Args[2:])
		if *id != "" {
			fmt.Println(beers[*id])
		} else {
			fmt.Println(beers)
		}
	}
}
