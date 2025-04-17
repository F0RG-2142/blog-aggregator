package main

import (
	"bufio"
	"os"

	"github.com/F0RG-2142/blog-aggregator/internal/config"
)

func main() {
	body, _ := config.Read()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	currentUser := string(scanner.Text()[0])
	if body.CurrentUserName != currentUser {
		config.SetUser(currentUser)
	}
}
