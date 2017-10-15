package main

import (
	"bufio"
	"errors"
	"flag" /*!!!*/
	"fmt"
	"io"
	"os"
)

/*================================= types =========================*/

/*================================= prototypes ====================*/
/*void usage(void)
int main()
void process_args()
void process_input()*/
/*================================= FlagInit() ================*/
var (
	flagSet    = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	start_page = flagSet.Int("s", -1, "read from <s> page")
	end_page   = flagSet.Int("e", -1, "read until <e> page")
	page_len   = flagSet.Int("l", 72, "lines per page(default:72 lines/page)")/*test:set 2*/
	fin        = flagSet.Bool("f", false, "read one page until '\f' ")/*use test.txt to test*/
)
/*================================= process_args() ================*/
func process_args() bool {
	if len(os.Args) <= 2 {
		err := errors.New("command need both start_page:-s=number and  end_page:-e=number")
		fmt.Fprintln(os.Stderr, "warning(command format): ", err)
		return false
	}
	if os.Args[1][0:2] != "-s" {
		err := errors.New("command should be like as -s=number -e=number [options] [filename]")
		fmt.Fprintln(os.Stderr, "warning(command formant): ", err)
		return false
	}
	if *start_page <= 0 {
		err := errors.New("the start_page can not be less than 1")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	if os.Args[2][0:2] != "-e" {
		err := errors.New("command should be like as -s=number -e=number [options] [filename]")
		fmt.Fprintln(os.Stderr, "warning(command formant): ", err)
		return false
	}
	if *end_page <= 0 {
		err := errors.New("the end_page can not be less than 1")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	if *start_page > *end_page {
		err := errors.New("the end_page can not be less than start_page")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	if *page_len <= 0 {
		err := errors.New("the page_line can not be less than 0")
		fmt.Fprintln(os.Stderr, "warning(arguments): ", err)
		return false
	}
	return true
}
/*================================= process_input() ===============*/
func process_putin(Ibuf *bufio.Reader, Obuf *os.File) {
	var count int
	count = *end_page - *start_page + 1
	if !*fin { /*read all the char from the file from the startpage*/
		for i := 1; i < *start_page; i++ {
			for j := 0; j < *page_len; j++ {
				Ibuf.ReadString('\n')
			}
		}
		for i := 0; i < count; i++ {
			for j := 0; j < *page_len; j++ {
				line, err := Ibuf.ReadString('\n')
				if err != nil {
					if err == io.EOF &&
						i != count &&
						j != *page_len {
						err2 := errors.New("the pages in the file is too less to read")
						fmt.Fprintln(os.Stderr, "warning(file reading) ", err2)
						return
					} else {
						fmt.Fprint(os.Stderr, "warning(file reading) ", err.Error())
					}
				}
				if Obuf != nil {
					Obuf.WriteString(line)
				} else {
					fmt.Print(line)
				}
			}
		}
	} else { /*the cut of the page*/
		for i := 1; i < *start_page; i++ {
			Ibuf.ReadString('\f')
		}
		for i := 0; i < count; i++ {
			line, err := Ibuf.ReadString('\f')
			if err != nil {
				if err == io.EOF && i != count {
					err3 := errors.New("the pages in the file is too less to read")
					fmt.Fprintln(os.Stderr, "warning(file reading) ", err3)
					return
				} else {
					fmt.Fprint(os.Stderr, "warning(file reading) ", err.Error())
				}
			}
			if Obuf != nil {
				Obuf.WriteString(line)
			} else {
				fmt.Print(line)
			}
		}
	}
}

/*func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
O_RDONLY：只读模式(read-only)
O_WRONLY：只写模式(write-only)
O_RDWR：读写模式(read-write)
O_APPEND：追加模式(append)
O_CREATE：文件不存在就创建(create a new file if none exists.)
O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
O_TRUNC：打开并清空文件
至于操作权限perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０.虽然go语言给perm权限设定了很多的常量，但是习惯上也可以直接使用数字，如0666(具体含义和Unix系统的一致).
*/
/*  0：八进制
    6:读+写
    4：读
    7：读+写+执行*/
func write(In string, Ou string) {
	var Ibuf *bufio.Reader
	if In != "" {
		inFile, err := os.OpenFile(In, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning(open file) ", err.Error())
		}
		Ibuf = bufio.NewReader(inFile)
	} else {
		Ibuf = bufio.NewReader(os.Stdin)
	}
	var Obuf *os.File
	var err error
	if Ou != "" {
		Obuf, err = os.OpenFile(Ou, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning(open file) ", err.Error())
		}
	} else {
		Obuf = nil
	}
	process_putin(Ibuf, Obuf)
}


/*================================= main()=== =====================*/
func main() {
	flagSet.Parse(os.Args[1:])
	if process_args() {
		var inputFile string
		var outputFile string
		if flagSet.NArg() > 0 {
			inputFile = flagSet.Arg(0)
		}
		if flagSet.NArg() > 1 {
			outputFile = flagSet.Arg(1)
		}
		write(inputFile, outputFile)
	}
}

