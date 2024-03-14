package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What is your favorite season?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println("That is cool! your favorite season is " + s)
	fmt.Println("Whats your favorite holiday?")
	on := bufio.NewReader(os.Stdin)
	t, _ := on.ReadString('\n')
	t = strings.TrimSpace(t)
	t = strings.ToUpper(t)
	fmt.Println("Hey " + t + " is cool")

}
