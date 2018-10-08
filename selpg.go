
package main

import (
	"fmt"
)
import flag "github.com/spf13/pflag"

var startPage = flag.IntP("start","s",0,"start page")
var endPage = flag.IntP("end","e",0,"end page")

func main(){

	flag.Parse()

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("startPage=%d",*startPage)
	fmt.Println("endPage=%d",*endPage)
}