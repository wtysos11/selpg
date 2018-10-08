
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
import flag "github.com/spf13/pflag"

func main(){
	var startPage = flag.IntP("start","s",-1,"start page")
	var endPage = flag.IntP("end","e",-1,"end page")
	var lineNumber = flag.IntP("line","l",72,"line number in one page")
	var format = flag.BoolP("format","f",false,"whether to find \f and make it as standard")
	var destination = flag.StringP("dest","d","","destination of output")	
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
	if *startPage==-1 || *endPage == -1{
		fmt.Println("Please enter start page and end page")
		os.Exit(1)
	} else if *format == true && *lineNumber !=72{
		fmt.Println("-l and -f can't be used at the same time")
		os.Exit(1)
	}

	var data []byte
	var err error
	switch flag.NArg(){
	case 0:
		data,err = ioutil.ReadAll(os.Stdin)
		if(err!=nil){
			fmt.Println("error:",err)
			os.Exit(1)
		}
	case 1:
		data,err = ioutil.ReadFile(flag.Arg(0))
		if(err!=nil){
			fmt.Println("error:",err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Input must have one argument as file name or no argument.")
		os.Exit(1)
	}
	//拿到了一个[]byte型的数据data
	fmt.Printf("data receive:\n %s",string(data))
	//暂时没有做文件存储
	//行数固定为72行
	if *format == false {
		lineCount := 0
		pageCount := 1
		pointer := 0
		standard := "\n"
		for pointer < len(data) {
			//找到换行符
			for data[pointer] != standard[0] && pointer < len(data) {
				pointer++
				if pointer == len(data){//at the end
					lineCount++ //行数增加
					if lineCount > *lineNumber {
						lineCount = 1
						pageCount++
					}
					if pageCount >= *startPage && pageCount <= *endPage {
						fmt.Printf("%s",string(data[:]))
					}
					break
				}
			}
			fmt.Println(lineCount,pageCount,string(data[:pointer+1]))


			pointer++ //跳过换行符
			lineCount++ //行数增加
			if lineCount > *lineNumber {
				lineCount = 1
				pageCount++
			}

			
			if pageCount >= *startPage && pageCount <= *endPage {
				fmt.Printf("%s",string(data[:pointer]))
			}
			data = data[pointer:]
			pointer = 0
		}
	} else{
		pageCount := 0
		pointer := 0
		standard := "\f"
		for pointer<len(data) {
			for data[pointer]!=standard[0] {
				pointer++
				if pointer == len(data){
					pageCount++
					if pageCount >= *startPage && pageCount <= *endPage {
						fmt.Println(string(data[:]))
					}
					break;
				}
			}

			

			pointer++
			pageCount++

			if pageCount >= *startPage && pageCount <= *endPage {
				fmt.Printf("%s",string(data[:pointer]))
			}
			data = data[pointer:]
			pointer = 0
		}
	}

}