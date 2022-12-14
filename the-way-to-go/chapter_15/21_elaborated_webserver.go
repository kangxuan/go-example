package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var helloRequests = expvar.NewInt("hello-requests")

var webroot = flag.String("root", "/home/user", "web root directory")

// Counter 计数器
type Counter struct {
	n int
}

type Chan chan int

func main() {
	// 监听命令行输入
	flag.Parse()
	// 根目录处理
	http.Handle("/", http.HandlerFunc(Logger))
	// go/hello
	http.Handle("/go/hello", http.HandlerFunc(HelloHandler))
	// 计数器
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)

	http.Handle("/go/", http.StripPrefix("go", http.FileServer(http.Dir(*webroot))))

	http.Handle("/flags", http.HandlerFunc(FlagHandler))

	http.Handle("/args", http.HandlerFunc(ArgsHandler))

	http.Handle("/chan", ChanCreate())

	http.Handle("/date", http.HandlerFunc(DataHandler))

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panicln("ListenAndServe: ", err)
	}
}

func Logger(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.String())
	// 输出头404
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello, world!\n")
}

func (ctr *Counter) String() string {
	return fmt.Sprintf("%d", ctr.n)
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctr.n++
	case "POST":
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprintf(w, "counter reset\n")
		}
	}

	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func FlagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Flag:\n")
	// VisitAll() 函数迭代所有的标签
	flag.VisitAll(func(i *flag.Flag) {
		if i.Value.String() != i.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", i.Name, i.Value.String(), i.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", i.Name, i.Value.String())
		}
	})
}

func ArgsHandler(w http.ResponseWriter, r *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)

	return c
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Channel send #%d\n", <-ch))
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_r, _w, err := os.Pipe()
	if err != nil {
		_, err2 := fmt.Fprintf(w, "Pipe: %s\n", err)
		if err2 != nil {
			return
		}
		return
	}

	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, _w, _w}})
	defer func(_r *os.File) {
		err := _r.Close()
		if err != nil {
			return
		}
	}(_r)

	err = _w.Close()
	if err != nil {
		return
	}
	if err != nil {
		_, err2 := fmt.Fprintf(w, "fork/exec:%s\n", err)
		if err2 != nil {
			return
		}
		return
	}

	defer func(p *os.Process) {
		err := p.Release()
		if err != nil {
			return
		}
	}(p)

	_, err = io.Copy(w, _r)
	if err != nil {
		return
	}
	wait, err := p.Wait()
	if err != nil {
		_, err := fmt.Fprintf(w, "wait: %s\n", err)
		if err != nil {
			return
		}
		return
	}
	if !wait.Exited() {
		_, err2 := fmt.Fprintf(w, "date: %v\n", wait)
		if err2 != nil {
			return
		}
		return
	}
}
