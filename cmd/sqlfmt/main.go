package main

import (
	"fmt"
	"os"

	"github.com/Astra-Sattelite/sqlfmt/internal/formatter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: sqlfmt <file.sql>")
		os.Exit(1)
	}

	path := os.Args[1]

	sql, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlfmt: %v\n", err)
		os.Exit(1)
	}

	formatted := formatter.Format(string(sql))

	if err := os.WriteFile(path, []byte(formatted), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "sqlfmt: %v\n", err)
		os.Exit(1)
	}
}