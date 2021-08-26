package main

import (
	"fmt"
	"regexp"
)

func main() {
	location := "https://s3.sa-east-1.amazonaws.com/learning.golang.dev.water.cloud/teste/25/60eeae6b46bc1b00066ef4c1.zip"

	regex, _ := regexp.Compile("(.*water.cloud/)(.*)")

	matchs := regex.FindSubmatch([]byte(location))

	for match := range matchs {
		fmt.Println(string(matchs[match]))
	}
}
