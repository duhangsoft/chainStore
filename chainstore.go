package chainStore

import (
	"net"
	"net/http"
	"os"
)

func Run(webPort string, cashPort string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", safeHander(chainHandler))
	mux.HandleFunc("/public/", safeHander(publicHandler))
	//mux.HandleFunc("/cash"),safeHander(cashHandler)
	go func() {
		err := http.ListenAndServe(":"+webPort, mux)
		if err != nil {
			logger.Errorln(err)
			return
		}
		return
	}()
	lis, err := net.Listen("tcp", cashPort)
	defer lis.Close()

	if err != nil {
		logger.Errorln("Error when listen: ", cashPort)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			logger.Errorln("Error accepting client: ", err.Error())
		}
		go cashHandler(conn)

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
func publicHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "." + r.URL.Path
	f, err := os.Open(fileName)
	if err != nil {

	}
	defer f.Close()
	buf := make([]byte, 1024)

	f.Read(buf)
	w.Write(buf)
}
func cashHandler(con net.Conn) {
	logger.Debugln("New connection:" + con.RemoteAddr().String())
	var data = make([]byte, 1024)
	var res string
	for {
		length, err := con.Read(data)
		if err != nil {
			logger.Debugf("Client %v quit.\n", con.RemoteAddr().String())
			con.Close()
			return
		}
		res = string(data[0:length])
		logger.Debugln(con.RemoteAddr().String() + "send:" + res)
		res = "I am server ,you is " + con.RemoteAddr().String() + ", You said :" + res
		con.Write([]byte(res))

	}
}
