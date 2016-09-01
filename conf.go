package chainStore

import (
	"database/sql"
	"encoding/json"
	"os"
)

const (
	ConfigFile  = "conf.ini"
	CONFIGERROR = 0
	FIRST       = 1
	MANAGER     = 2
	DATEBASE    = 3
	CONFIGOK    = 100
)

type config struct {
	dbName string `json:"dbName"`
	dbUser string `json:"dbUser"`
	dbPwd  string `json:dbPwd`
	dbPort int32  `json:dbPort`
	amin   string `json:"admin"`
	pwd    string `json:"pwd"`
	db     *sql.DB
}

func newConfig() *config {
	return &config{}
}
func (c *config) read() int {
	f, err := os.OpenFile(ConfigFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		logger.Errorln(err)
		return CONFIGERROR
	}
	defer f.Close()
	var buf []byte

	n, err := f.Read(buf)

	if err != nil {
		logger.Errorln(err)
		return CONFIGERROR
	}
	if n == 0 {
		return FIRST
	}
	err = json.Unmarshal(buf, c)
	if err != nil {
		logger.Errorln(err)
		return CONFIGERROR
	}
	if c.amin == "" {
		return MANAGER
	}
	if c.dbName == "" {
		return DATEBASE
	}
	err = c.db.Ping()
	if err != nil {
		c.db, err = sql.Open("mysql", c.dbUser+":"+c.dbPwd+"@/"+c.dbName+"?charset=utf8")
	}
	return CONFIGOK
}
func (c *config) write() error {
	f, err := os.OpenFile(ConfigFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return err
	}
	defer f.Close()
	buf, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = f.Write(buf)
	return err
}
