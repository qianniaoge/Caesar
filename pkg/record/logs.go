package record

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

/*
标准的日志记录器
*/

type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var Logger = logrus.New()

func Logs(logPath string, silence bool) {

	// 为当前logrus实例设置消息输出格式为json格式。
	Logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableHTMLEscape: true,
	}
	// 是否记录日志位置
	Logger.SetReportCaller(false)
	Logger.SetLevel(logrus.TraceLevel)

	file, err := os.Create(logPath)
	// os.OpenFile("logPath", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err == nil {

		Logger.SetOutput(ioutil.Discard)
		Logger.AddHook(&WriterHook{ // Send logs with level higher than warning to stderr
			Writer:    file,
			LogLevels: logrus.AllLevels,
		})
		if !silence {
			Logger.AddHook(&WriterHook{ // Send info and debug logs to stdout
				Writer: os.Stdout,
				LogLevels: []logrus.Level{
					logrus.PanicLevel,
					logrus.FatalLevel,
					logrus.WarnLevel,
					logrus.DebugLevel,
				},
			})
		}

	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}

}
