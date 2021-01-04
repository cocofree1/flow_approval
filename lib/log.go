package lib

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/astaxie/beego/config"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	DEBUG = "debug"
	TRACE = "trace"
	INFO = "info"
	WARN = "warn"
	ERROR = "error"
	CRITICAL = "critical"
)

var (
	Log l4g.Logger = make(l4g.Logger)
)

func init() {
	// 获取配置信息
	conf, err := config.NewConfig("ini", "conf/log.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	console,_ := conf.Bool("log_console")
	printLevel := conf.String("log_print_level")
	filename := conf.String("log_filename")
	rotateDaily,_ := conf.Bool("log_rotate_daily")
	maxSize := conf.String("log_maxsize")
	maxLine := conf.String("log_rotate_lines")
	maxBakeup,_ := conf.Int("log_maxbakeup")

	format := "[%D %T] [%L] [%S] (message: %M)"
	// 打印到控制台
	if console {
		flwConsole := l4g.NewConsoleLogWriter()
		flwConsole.SetFormat(format)
		Log.AddFilter("stdout", GetLevelByString(printLevel), flwConsole)
	}


	// 创建日志文件
	dir := filepath.Dir(filename)
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0766)
		if nil != err {
			fmt.Printf("get mkdir =(%+v) err=(%+v)", dir, err)
			os.Exit(1)
		}
	}

	// 将日志写入文件
	flw := l4g.NewFileLogWriter(filename, false)
	flw.SetFormat(format)

	// 日志按天分割
	if rotateDaily {
		flw.SetRotate(true)
		flw.SetRotateDaily(rotateDaily)
		fmt.Println("log_rotate_daily=", rotateDaily)
	}

	// 日志按大小分割
	var iMaxSize int
	if maxSize != "" {
		iMaxSize, err = GetIntDataByKMG(maxSize)
		if err != nil {
			fmt.Printf("Fail to parse log_maxsize=(%+v)\n", maxSize)
			os.Exit(1)
		}
		flw.SetRotate(true)
		flw.SetRotateSize(iMaxSize)
		fmt.Println("log_maxsize=", iMaxSize)
	}

	// 按日志行数进行分割
	var iMaxLine int
	if maxLine != "" {
		iMaxLine, err = GetIntDataByKMG(maxLine)
		if err != nil{
			fmt.Printf("Fail to parse log_rotate_lines=(%+v)\n", maxSize)
			os.Exit(1)
		}
		flw.SetRotate(true)
		flw.SetRotateLines(iMaxLine)
		fmt.Printf("log_rotate_lines=(%+v)\n", iMaxLine)
	}

	// 备份的日志文件数量
	if maxBakeup != 0 {
		flw.SetRotateMaxBackup(maxBakeup)
		flw.SetRotate(true)
	}
	Log.AddFilter("file", GetLevelByString(printLevel), flw)
}

// 根据等级返回等级
func GetLevelByString(strLevel string) l4g.Level {
	level := strings.ToLower(strLevel)
	if level == DEBUG {
		return l4g.DEBUG
	} else if level == INFO {
		return l4g.INFO
	} else if level == WARN {
		return l4g.WARNING
	} else if level == ERROR {
		return l4g.ERROR
	} else if level == CRITICAL {
		return l4g.CRITICAL
	} else {
		return l4g.DEBUG
	}
}

func GetIntDataByKMG(data string) (value int, err error) {
	data = strings.ToUpper(data)
	base := 1
	var suffix string
	if strings.HasSuffix(data, "K") {
		base = 1024
		suffix = "K"
	} else if strings.HasSuffix(data, "M") {
		base = 1024 * 1024
		suffix = "M"
	} else if strings.HasSuffix(data, "G") {
		base = 1024 * 1024 * 1024
		suffix = "G"
	}

	if suffix != "" {
		data = strings.TrimSuffix(data, suffix)
	}
	num, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	value = num * base
	return value, err
}