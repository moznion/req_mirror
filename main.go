package reqMirror

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	ver string
	rev string
)

type req struct {
	Method           string
	URL              *url.URL
	Proto            string
	Header           http.Header
	Body             string
	TransferEncoding []string
	Host             string
}

func writeError(w http.ResponseWriter, err error) {
	log.Printf(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, `{"message":"Internal server error"}`)
}

func Run(args []string) {
	var port uint
	flag.UintVar(&port, "port", 22222, "Port to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeError(w, err)
			return
		}

		b, err := json.Marshal(&req{
			Method:           r.Method,
			URL:              r.URL,
			Proto:            r.Proto,
			Header:           r.Header,
			Body:             string(body),
			TransferEncoding: r.TransferEncoding,
			Host:             r.Host,
		})
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(b))
		return
	})

	log.Printf("Listening 127.0.0.1:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
