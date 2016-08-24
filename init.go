package chainStore

import (
	"os"
)

func initPage() []byte {
	f, err := os.Open("../views/init.html")
	if err != nil {
		return errorPage()
	}
	defer f.Close()
	buf := make([]byte, 10240)
	_, err = f.Read(buf)
	if err != nil {
		return errorPage()
	}
	return buf
}
func errorPage() []byte {
	return []byte("error page")
}
