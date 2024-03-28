package process

import (
	"flag"
	"fmt"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myFlags arrayFlags

func stringArrayFlag() {
	flag.Var(&myFlags, "list1", "Some description for this param.")
	flag.Parse()

	fmt.Println(myFlags)
}
