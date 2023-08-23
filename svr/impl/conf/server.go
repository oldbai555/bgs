package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/conf"
	"github.com/oldbai555/lbtool/log"
	"os"
	"path/filepath"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

const defaultServerConfigPath = "/Users/zhangjianjun/work/mypro/bgs/svr/impl/conf/server.json"

func LoadServerConfig(path string) {
	if path == "" {
		path = defaultServerConfigPath
	}
	data, err := os.ReadFile(filepath.ToSlash(defaultServerConfigPath))
	if err != nil {
		log.Errorf("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Errorf("%v", err)
	}

	conf.LogLevel = Server.LogLevel
	conf.LogPath = Server.LogPath
	conf.LogFlag = LogFlag
	conf.ConsolePort = Server.ConsolePort
	conf.ProfilePath = Server.ProfilePath
}
