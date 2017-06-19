package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	ospath "path"
	"strings"

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
	path := r.FormValue("path")
	if path == "" || strings.Contains(path, "..") {
		fmt.Fprintf(w, "path is empty or include two dot\n")
		return
	}
	if !strings.Contains(path, "./") {
		if *base != "" {
			path = ospath.Join(*base, path)
		}
	}
	args := r.FormValue("args")

	out, err := Run(path, []string{args})
	if err != nil {
		fmt.Fprintf(w, "%v\nerror: %v\n", out, err)
	} else {
		fmt.Fprintf(w, "%v", out)
	}
	var outline string
	fmt.Sscanln(out, &outline)
	log.Printf("cmd: %v, args: %v, out: %v, err: %v\n", path, args, outline, err)
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
