
package main

import (
	"fmt"
)
import flag "github.com/spf13/pflag"

var startPage = flag.IntP("start","s",-1,"start page")
var endPage = flag.IntP("end","e",-1,"end page")
var lineNumber = flag.IntP("line","l",72,"line number in one page")
var format = flag.BoolP("format","f",false,"whether to find \f and make it as standard")
var destination = flag.StringP("dest","d","","destination of output")

func main(){

	flag.Parse()
//debug
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("startPage=",*startPage)
	fmt.Println("endPage=",*endPage)
	fmt.Println("line=",*lineNumber)
	fmt.Println("format=",*format)
	fmt.Println("desination=",*destination)
//legal check
	//-s and -e must use

	//-l and -f can't be use at the same time

}
