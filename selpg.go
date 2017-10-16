/*================================= includes ======================*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*================================= types =========================*/
type selpgArgs struct {
	startPage  int
	endPage    int
	inFilename string
	pageLen    int
	pageType   rune
	printDest  string
}

/*================================= main()=== =====================*/
func main() {
	sa := new(selpgArgs)
	progname = os.Args[0]

	sa.startPage = -1
	sa.endPage = -1
	sa.inFilename = ""
	sa.pageLen = 5
	sa.pageType = 'l'
	sa.printDest = ""

	processArgs(len(os.Args), os.Args, sa)
	processInput(sa)
}

/*================================= globals =======================*/
var progname string /* program name, for error messages */

func usage() {
	fmt.Printf("Usage of %s:\n\n", progname)
	fmt.Printf("\tselpg -s<Number> -e<Number> [options] [filename]\n\n")
	fmt.Printf("\t-s<Number>\tStart from Page <number>.\n")
	fmt.Printf("\t-e<Number>\tEnd to Page <number>.\n")
	fmt.Printf("\t-l<Number>\t[options]Specify the number of line per page.Default is 5.\n")
	fmt.Printf("\t-f\t\t[options]Specify that the pages are sperated by \\f.\n")
	fmt.Printf("\t[filename]\t[options]Read input from the file.\n\n")
	fmt.Printf("If no file specified, %s will read input from stdin. Control-D to end.\n\n", progname)
}

/*================================= process_args() ================*/
func processArgs(ac int, args []string, sa *selpgArgs) {
	if ac < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		usage()
		os.Exit(1)
	}

	temp := []rune(args[1])
	if temp[0] != '-' || temp[1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstart_page\n", progname)
		usage()
		os.Exit(2)
	}
	page, err := strconv.Atoi(string(temp[2:]))
	if page < 1 || err != nil {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n", progname, string(temp[2:]))
		usage()
		os.Exit(3)
	}
	sa.startPage = page

	temp = []rune(args[2])
	if temp[0] != '-' || temp[1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n", progname)
		usage()
		os.Exit(4)
	}
	page, err = strconv.Atoi(string(temp[2:]))
	if page < 1 || page < sa.startPage || err != nil {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n", progname, string(temp[2:]))
		usage()
		os.Exit(5)
	}
	sa.endPage = page

	endIndex := 3
	for endIndex < ac && []rune(args[endIndex])[0] == '-' {
		temp = []rune(args[endIndex])

		switch temp[1] {
		case 'l':
			lineNum, err := strconv.Atoi(string(temp[2:]))
			if lineNum < 1 || err != nil {
				fmt.Fprintf(os.Stderr, "%s: invalid page length %s\n", progname, string(temp[2:]))
				usage()
				os.Exit(6)
			}
			sa.pageLen = lineNum
			endIndex++

		case 'f':
			if strings.Compare(string(temp), "-f") != 0 {
				fmt.Fprintf(os.Stderr, "%s: option should be \"-f\"\n", progname)
				usage()
				os.Exit(7)
			}
			sa.pageType = 'f'
			endIndex++

		case 'd':
			if len(temp[2:]) < 1 {
				fmt.Fprintf(os.Stderr, "%s: -d option requires a printer destination\n", progname)
				usage()
				os.Exit(8)
			}
			sa.printDest = string(temp[2:])
			endIndex++

		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option %s\n", progname, string(temp))
			usage()
			os.Exit(9)
		}
	}

	if endIndex <= ac-1 {
		sa.inFilename = args[endIndex]
		f, err := os.Open(sa.inFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n", progname, sa.inFilename)
			os.Exit(10)
		}
		f.Close()
	}
}

/*================================= process_input() ===============*/
func processInput(sa *selpgArgs) {
	fin := os.Stdin
	var err error
	if sa.inFilename != "" {
		fin, err = os.Open(sa.inFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", progname, sa.inFilename)
			os.Exit(11)
		}
	}

	fout := os.Stdout
	var cmd *exec.Cmd
	if sa.printDest != "" {
		temp := fmt.Sprintf("./%s", sa.printDest)
		cmd = exec.Command("sh", "-c", temp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n", progname, temp)
			os.Exit(12)
		}
	}

	var line string
	pageCount := 1
	inputReader := bufio.NewReader(fin)
	rst := ""
	if sa.pageType == 'l' {
		lineCount := 0

		for true {
			line, err = inputReader.ReadString('\n')
			if err != nil {
				break
			}
			lineCount++
			if lineCount > sa.pageLen {
				pageCount++
				lineCount = 1
			}
			if pageCount >= sa.startPage && pageCount <= sa.endPage {
				if sa.printDest == "" {
					fmt.Fprintf(fout, "%s", line)
				} else {
					rst += line
				}
			}
		}
	} else {
		for true {
			c, _, erro := inputReader.ReadRune()
			if erro != nil {
				break
			}
			if c == '\f' {
				pageCount++
			}
			if pageCount >= sa.startPage && pageCount <= sa.endPage {
				if sa.printDest == "" {
					fmt.Fprintf(fout, "%c", c)
				} else {
					rst += string(c)
				}
			}
		}
	}

	if sa.printDest != "" {
		cmd.Stdin = strings.NewReader(rst)
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println("print error!")
		}
	}

	/*page type*/

	if pageCount < sa.startPage {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", progname, sa.startPage, pageCount)
	} else {
		if pageCount < sa.endPage {
			fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, sa.endPage, pageCount)
		}
	}

	fin.Close()
	fout.Close()
	fmt.Fprintf(os.Stderr, "test error\n", progname)
}

/*================================= EOF=== =====================*/
