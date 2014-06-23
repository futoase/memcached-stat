package main

import (
	"flag"
	"fmt"
	mc "github.com/futoase/memcached-stat/libs"
	"os"
)

func main() {
	address := flag.String("address", "localhost:11211", "Address of memcached")
	flag.Parse()
	fmt.Printf("Server: %s\n", *address)

	con, err := mc.Connect(*address)
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}
	defer con.Close()

	stat, _ := con.Stats("")
	fmt.Printf("%s", stat)
}
