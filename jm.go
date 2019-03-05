package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/guonaihong/coreutils/utils"
	"github.com/guonaihong/flag"
	"os"
	"strings"
)

type jm struct {
	data        *[]string
	object      *bool
	array       *bool
	delimiter   *string
	combination *bool //支持变量表示法
}

func (j *jm) main(fd *os.File) {

	br := bufio.NewReader(fd)

	isPrint := false
	for {

		l, e := br.ReadString('\n')
		if e != nil && len(l) == 0 {
			break
		}

		ls := strings.FieldsFunc(l, func(r rune) bool {
			return r == rune((*j.delimiter)[0])
		})

		title := *j.data
		fields := append([]string{}, append(title, ls...)...)

		if *j.object {
			isPrint = true
			marshalObject(fields)
		}

		if *j.array {
			isPrint = true
			marshalArray(fields)
		}
	}

	if !isPrint {
		if *j.object {
			marshalObject(*j.data)
		}

		if *j.array {
			marshalArray(*j.data)
		}
	}
}

func die(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func marshalObject(object []string) {
	m := map[string]interface{}{}

	if len(object)%2 != 0 {
		die("jm:The number of parameters must be even:%d", len(object))
	}

	for i := 0; i < len(object); i += 2 {
		m[object[i]] = object[i+1]
	}

	all, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)
}

func marshalArray(array []string) {
	var a []interface{}
	for _, v := range array {
		a = append(a, v)
	}

	all, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)
}

func main() {
	command := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	j := jm{}

	j.data = command.Opt("data", "json data").
		Flags(flag.GreedyMode).
		NewStringSlice([]string{})

	j.object = command.Opt("o, object", "json object").
		Flags(flag.PosixShort).
		NewBool(false)

	j.array = command.Opt("a, array", "json array").
		Flags(flag.PosixShort).
		NewBool(false)

	j.combination = command.Opt("c, combination", "Combine command line and file data").
		Flags(flag.PosixShort).
		NewBool(false)

	j.delimiter = command.Opt("d, delimiter", "use DELIM instead of TAB for field delimiter").
		Flags(flag.PosixShort).
		NewString("\t")

	command.Parse(os.Args[1:])

	args := command.Args()
	if len(args) == 0 {
		args = append(args, "-")
	}

	for _, v := range args {
		fd, err := utils.OpenInputFd(v)
		if err != nil {
			die("jm: %s\n", err)
		}

		j.main(fd)

		utils.CloseInputFd(fd)
	}
}
