package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/starkandwayne/concourse-resource-rss/tree/develop/in"
)

func main() {
	indata, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	returndata, _ := in.Execute()

}
