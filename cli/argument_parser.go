package cli

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Dir   string
	Port  int
	Token bool
	Serve bool
}

func ParseFlags() Config {
	var config Config

	if len(os.Args) < 2 || os.Args[1] != "serve" {
		showHelp()
		os.Exit(1)
	}

	config.Serve = true

	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	fs.StringVar(&config.Dir, "d", "~/datagate", "Directory to serve files from")
	fs.IntVar(&config.Port, "p", 8080, "Port to run the server on")
	fs.BoolVar(&config.Token, "t", false, "Enable token-based authentication")

	if err := fs.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing flags:", err)
		showHelp()
		os.Exit(1)
	}

	return config
}

func showHelp() {
	fmt.Println("Usage: datagate serve [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  -d string   Directory to serve files from (default \"./\")")
	fmt.Println("  -p int      Port to run the server on (default 8080)")
	fmt.Println("  -t          Enable token-based authentication")
}
