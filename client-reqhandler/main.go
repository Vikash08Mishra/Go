package main

import (
	"fmt"
)

func main() {
	fmt.Println("Setting up routes path and starting web server.")

	// run as separate go routine if any setup/initialization post this.
	// last set up should run on main go routine to block main from exiting
	SetupAndStartServer()

}
