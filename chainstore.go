package chainStore

import (
	"net/http"
)

func Run(url string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", safeHander(chainHandler))
	err := http.ListenAndServe(url, mux)
	if err != nil {
		logger.Errorln(err)
		return
	}
}
func safeHander(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				logger.Errorln(e.Error())
			}
		}()
		fn(w, r)
	}
}
func chainHandler(w http.ResponseWriter, r *http.Request) {
	//do something
	conf := newConfig()

	switch conf.read() {
	case CONFIGERROR:
		w.Write(errorPage())
	case FIRST:
		w.Write(initPage())
	default:
		w.Write([]byte("error"))
	}
}
