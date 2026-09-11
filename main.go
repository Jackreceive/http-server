package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var methods = map[string]bool{
	"GET": true, "POST": true, "PUT": true, "DELETE": true,
	"HEAD": true, "OPTIONS": true, "PATCH": true,
}

func isVersion(s string) bool {
	// TODO: must look like "HTTP/<digit>.<digit>"
	return strings.HasPrefix(s, "HTTP/")
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimRight(sc.Text(), "\r")
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		// TODO: 3 parts, valid method, path begins with "/", valid version
		if len(parts) != 3 || !methods[parts[0]] || !strings.HasPrefix(parts[1], "/") {
			fmt.Println("INVALID")
			continue
		}
		fmt.Printf("METHOD=%s PATH=%s VERSION=%s\n", parts[0], parts[1], parts[2])
	}
}
