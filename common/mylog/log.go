package mylog

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
	"xpopenapi/common/config"
	"xpopenapi/common/util"
)

var log *logrus.Logger

var Atag = "xtx_xp"

var LogFilePath string //日志文件前缀

func init() {
	LogFilePath = "/data/logs/" + Atag
	log = logrus.New()
	//设置日志级别为warn以上
	log.SetLevel(logrus.TraceLevel)
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	// 设置日志格式为json格式
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.AddHook(newLfsHook())
}

func Info(args ...interface{}) {
	arr := []string{}
	for _, arg := range args {
		s := util.ToStr(arg)
		arr = append(arr, s)
	}
	myText := strings.Join(arr, " ")
	log.WithFields(logrus.Fields{
		"Line":    "",
		"Method":  "",
		"Service": Atag,
	}).Info(myText)
}

func Error(args ...interface{}) {
	arr := []string{}
	for _, arg := range args {
		s := util.ToStr(arg)
		arr = append(arr, s)
	}
	myText := strings.Join(arr, " ")
	Line := ""
	Method := ""
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	Line = fmt.Sprintf("%s:%d", file, line)
	Method = f.Name()
	log.WithFields(logrus.Fields{
		"Line":    Line,
		"Method":  Method,
		"Service": Atag,
	}).Error(myText)
}

func Debug(args ...interface{}) {
	arr := []string{}
	for _, arg := range args {
		s := util.ToStr(arg)
		arr = append(arr, s)
	}
	myText := strings.Join(arr, " ")
	Line := ""
	Method := ""
	if config.Base.Env == "dev" {
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		Line = fmt.Sprintf("%s:%d", file, line)

		Method = f.Name()
	}

	log.WithFields(logrus.Fields{
		"Line":    Line,
		"Method":  Method,
		"Service": Atag,
	}).Debug(myText)
}

//DBG 调试日志
func DBG(text string, args ...interface{}) {
	myText := fmt.Sprintf(text, args...)
	Line := ""
	Method := ""
	if config.Base.Env == "dev" {
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		Line = fmt.Sprintf("%s:%d", file, line)

		Method = f.Name()
	}

	log.WithFields(logrus.Fields{
		"Line":    Line,
		"Method":  Method,
		"Service": Atag,
	}).Debug(myText)
}

// LogInfo 运行日志
func LogInfo(text string, args ...interface{}) {
	myText := fmt.Sprintf(text, args...)
	log.WithFields(logrus.Fields{
		"Line":    "",
		"Method":  "",
		"Service": Atag,
	}).Info(myText)
}

// LogWarn 警告日志
func LogWarn(text string, args ...interface{}) {
	myText := fmt.Sprintf(text, args...)
	log.WithFields(logrus.Fields{
		"Line":    "",
		"Method":  "",
		"Service": Atag,
	}).Warn(myText)
}

// LogError 错误日志
func LogError(text string, args ...interface{}) {
	myText := fmt.Sprintf(text, args...)
	Line := ""
	Method := ""
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	Line = fmt.Sprintf("%s:%d", file, line)
	Method = f.Name()
	log.WithFields(logrus.Fields{
		"Line":    Line,
		"Method":  Method,
		"Service": Atag,
	}).Error(myText)
}

type DefaultFieldsHook struct {
	IsCaller bool
	Line     string
	Method   string
}

func (df *DefaultFieldsHook) Fire(entry *logrus.Entry) error {
	if df.IsCaller && df.Line != "" && df.Method != "" {
		entry.Data["Method"] = df.Method
		entry.Data["Line"] = df.Line
	} else {
		entry.Data["Method"] = ""
		entry.Data["Line"] = ""
	}
	return nil
}

func (df *DefaultFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func newLfsHook() logrus.Hook {
	writer, err := rotatelogs.New(
		LogFilePath+"_%Y%m%d"+".log",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(LogFilePath),

		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		rotatelogs.WithMaxAge(72*time.Hour),
		//rotatelogs.WithRotationCount(maxRemainCnt),
		//rotatelogs.WithMaxAge(time.Minute), // 文件最大保存时间

	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	//level, ok := logLevels[*logLevel]
	//
	//if ok {
	//	logrus.SetLevel(level)
	//} else {
	//	logrus.SetLevel(logrus.WarnLevel)aliyun-log-go-sdk
	//}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return lfsHook
}
