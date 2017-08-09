package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	ospath "path"
	"strings"
	"unicode"

	"github.com/chinglinwen/log"
	"github.com/natefinch/lumberjack"
)

var (
	port = flag.String("p", "9001", "listening port")
	base = flag.String("base", "", "base dir")
)

func main() {
	log.Printf("starting... base: %v", *base)
	http.HandleFunc("/", cmdHandler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("run")
	fmt.Printf("%#v\n", s)

	cmd := strings.FieldsFunc(s, fieldsFunc())
	var path string
	var args []string
	n := len(cmd)
	if n == 0 {
		fmt.Fprintf(w, "path is empty\n")
		return
	}
	path = cmd[0]
	if path == "" || strings.Contains(path, "..") {
		fmt.Fprintf(w, "path is empty or include two dot\n")
		return
	}
	if !strings.Contains(path, "./") {
		if *base != "" {
			path = ospath.Join(*base, path)
		}
	}
	if n > 1 {
		args = append(args, cmd[1:]...)
	}
	for i, v := range args {
		args[i] = strings.Trim(v, `'"`)
	}

	out, err := Run(path, args)
	if err != nil {
		fmt.Fprintf(w, "%v\nerror: %v\n", out, err)
	} else {
		fmt.Fprintf(w, "%v", out)
	}
	var outline string
	fmt.Sscanln(out, &outline)
	log.Printf("cmd: %v, args: %v, out: %v, err: %v\n", path, args, outline, err)
}

func fieldsFunc() func(c rune) bool {
	lastQuote := rune(0)
	return func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return false
		default:
			return unicode.IsSpace(c)

		}
	}
}

func Run(path string, args []string) (string, error) {
	cmd := exec.Command(path, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func init() {
	flag.Parse()
	if *base == "./" {
		*base = os.Getenv("PWD")
	}
	log.SetOutput(&lumberjack.Logger{
		Filename:   "h2s.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
}
