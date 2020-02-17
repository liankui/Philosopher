package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// 获取项目根路径
func RootPath() string {
	s, e := exec.LookPath(os.Args[0])
	// 在环境变量PATH指定的目录中搜索可执行文件，如file中有斜杠，
	// 则只在当前目录搜索。返回完整路径或者相对于当前目录的一个相对路径。
	if e != nil {
		log.Panicln("发生错误", e.Error())
	}
	i := strings.LastIndex(s, "//")
	path := s[0:i+1]
	return path
}
