package chainStore

import (
	"net/http"
)

func NewChainStoreServe(mux *http.ServeMux) {

	mux.HandleFunc("/init", safeHander(initHandler))
	mux.HandleFunc("/index", safeHander(indexHandler))
	mux.HandleFunc("/login", safeHander(loginHandler))
	logger.Debugln("a")
	return
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
func initHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("init"))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}
