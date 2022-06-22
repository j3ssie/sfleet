package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var logger = logrus.New()

// InitLog init log
func InitLog(logLevel string) {
	// options.LogFile = path.Join(options.Scan.TmpOutput, fmt.Sprintf("metabigor-%s.log", GetTS()))
	// f, err := os.OpenFile(options.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	//     logger.Error("error opening file: %v", err)
	// }

	// mwr := io.MultiWriter(os.Stderr, nil)

	logger.SetLevel(logrus.InfoLevel)
	logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &prefixed.TextFormatter{
			ForceColors:     true,
			ForceFormatting: true,
		},
	}

	if logLevel == "debug" {
		// logger.SetOutput(mwr)
		logger.SetLevel(logrus.DebugLevel)
	} else if logLevel == "verbose" {
		// logger.SetOutput(mwr)
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetOutput(ioutil.Discard)
	}
}

// PrintLine print seperate line
func PrintLine() {
	dash := color.HiWhiteString("-")
	fmt.Println(strings.Repeat(dash, 40))
}

// GoodF print good message
func GoodF(format string, args ...interface{}) {
	good := color.HiGreenString("[+]")
	fmt.Printf("%s %s\n", good, fmt.Sprintf(format, args...))
}

// BannerF print info message
func BannerF(format string, data string) {
	banner := fmt.Sprintf("%v%v%v ", color.WhiteString("["), color.BlueString(format), color.WhiteString("]"))
	fmt.Printf("%v%v\n", banner, color.HiGreenString(data))
}

// BlockF print info message
func BlockF(name string, data string) {
	banner := fmt.Sprintf("%v%v%v ", color.WhiteString("["), color.GreenString(name), color.WhiteString("]"))
	fmt.Printf(fmt.Sprintf("%v%v\n", banner, data))
}

// BadBlockF print info message
func BadBlockF(name string, data string) {
	banner := fmt.Sprintf("%v%v%v ", color.WhiteString("["), color.RedString(name), color.WhiteString("]"))
	fmt.Printf(fmt.Sprintf("%v%v\n", banner, data))
}

// InforF print info message
func InforF(format string, args ...interface{}) {
	logger.Info(fmt.Sprintf(format, args...))
}

// ErrorF print good message
func ErrorF(format string, args ...interface{}) {
	logger.Error(fmt.Sprintf(format, args...))
}

// WarnF print good message
func WarnF(format string, args ...interface{}) {
	logger.Warning(fmt.Sprintf(format, args...))
}

// TraceF print good message
func TraceF(format string, args ...interface{}) {
	logger.Trace(fmt.Sprintf(format, args...))
}

// DebugF print debug message
func DebugF(format string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(format, args...))
}

// Emojif print good message
func Emojif(e string, format string, args ...interface{}) string {
	emj := strings.TrimSpace(emoji.Sprint(e))
	return fmt.Sprintf("%1s %s", emj, fmt.Sprintf(format, args...))
}
