package log

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// New 全局日志
// 六种日志级别：
// logrus.Debug("Useful debugging information.")
// logrus.Info("Something noteworthy happened!")
// logrus.Warn("You should probably take a look at this.")
// logrus.Error("Something failed but I'm not quitting.")
// logrus.Fatal("Bye.")   //log之后会调用os.Exit(1)
// logrus.Panic("I'm bailing.")   //log之后会panic()
// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
func New() *logrus.Logger {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	year, month, day := time.Now().Date()
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(strconv.Itoa(year)+"-"+strconv.Itoa(int(month))+"-"+strconv.Itoa(day)+"-"+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Error("log error Failed to log to file, using default stderr")
	}
	return log
}

/*
// 写入redis
func init() {
	hookConfig := logredis.HookConfig{
		Host:     "127.0.0.1",
		Key:      "my_redis_key",
		Format:   "v0",
		App:      "my_app_name",
		Port:     6379,
		Hostname: "my_app_hostname", // will be sent to field @source_host
		DB:       0,                 // optional
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
	// 静止写入stdout 提高性能
	logrus.SetOutput(ioutil.Discard)
}
*/
func init() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.

}
