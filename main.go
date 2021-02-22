package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"path/filepath"
)

var (
	helpFlag = flag.Bool("h", false, "Show this help")
	inFile = flag.String("i", "blackList.txt", "File to read blacklist entries from")
	stdinFlag = flag.Bool("si", false, "Read blacklist entries from stdin")
	stdoutFlag = flag.Bool("so", false, "Write serial numbers to stdout")
	outFile = flag.String("o", "authList.txt", "File to write new serial numbers to")
	appendFlag = flag.Bool("a", false, "Overwrite entries if file exists")
	err error
)

const usage = "`stripSn` [options]"

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}

	var serNumList []string
	if *stdinFlag {
		reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
		for {
			line := readFromStdin(reader)
			if line != "" {
				str := strings.Fields(line)
				// shortcut method, anticipating copy-paste from cli blacklist
				for _, s := range str {
					if len(s) == 12 {
							serNumList = append(serNumList, s)
					}
				}
			} else {
				fmt.Println("Received empty line, continuing...")
				break
			}
		}
	} else {
		fp, err := filepath.Abs(*inFile)
		if err != nil {
			fmt.Println(err); return
		}
		f, err := os.Open(fp)
		if err != nil {
			fmt.Println(err); return
		}
		defer f.Close()
		serNumList = readFromFile(f)
		if len(serNumList) < 1 {
			fmt.Printf("Did not find any serial numbers in %s\n", f.Name()); return
		} else {
			fmt.Printf("Read %d serial numbers from %s\n", len(serNumList), f.Name())
		}
	}
	if *stdoutFlag {
		for _, sn := range serNumList {
			fmt.Printf("%s\n", sn)
		}
		fmt.Printf("Wrote %d serial numbers to stdout\n", len(serNumList))
	} else {
		var of *os.File
		ofp, err := filepath.Abs(*outFile)
		if err != nil {
			fmt.Println(err); return
		}
		if *appendFlag {
			_, err = os.Stat(ofp)
			if err == nil {
				err = os.Remove(ofp)
				if err != nil {
					fmt.Println("Cannot remove existing file at %s, error: %v\nDumping contents to stdout\n\n", ofp, err)
					for _, af := range serNumList {
						fmt.Printf("%s\n", af)
					}
					return
				}
			}
		}
		of, err = os.OpenFile(ofp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err); return
		}
		defer of.Close()
				
		for _, sn := range serNumList {
			fmt.Fprintf(of, "%s\n", sn)
		}
		fmt.Printf("Wrote %d serial numbers to %s\n", len(serNumList), of.Name())
	}
	
	
}

func readFromFile(f *os.File) []string {
	s := bufio.NewScanner(f)
	var st []string
	for s.Scan() {
		str := strings.Split(strings.TrimSpace(s.Text()), " ")
		for _, l := range str {
			if len(l) == 12 {
				st = append(st, l)
			}
		}
	}
	return st
}

func readFromStdin(r *bufio.Reader) string {
	a, _, err := r.ReadLine()
	if err == io.EOF {
		return ""
	} else if err != nil {
		panic(err)
	}
	
	return strings.TrimRight(string(a), "\r\n")
}
