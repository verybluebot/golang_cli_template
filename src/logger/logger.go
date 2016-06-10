package logger

import(
    "github.com/Sirupsen/logrus"
)

func InitLogger() {
    var textFormater = new(logrus.TextFormatter)
    textFormater.TimestampFormat = "2006-01-02 15:04:05"
    textFormater.FullTimestamp = true

    logrus.SetFormatter(textFormater)
}

func Printf(format string, v ...interface{}) {
    logrus.Printf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
    logrus.Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
    logrus.Panicf(format, v...)
}

func Debugf(format string, v ...interface{}) {
    logrus.Debugf(format, v...)
}

func Warnf(format string, v ...interface{}) {
    logrus.Warnf(format, v...)
}

func Warningf(format string, v ...interface{}) {
    logrus.Warningf(format, v...)
}

func Errorf(format string, v ...interface{}) {
    logrus.Errorf(format, v...)
}


func Print(v ...interface{}) {
    logrus.Print(v...)
}

func Fatal( v ...interface{}) {
    logrus.Fatal(v...)
}

func Panic(v ...interface{}) {
    logrus.Panic(v...)
}

func Debug(v ...interface{}) {
    logrus.Debug(v...)
}

func Warn(v ...interface{}) {
    logrus.Warn(v...)
}

func Warning(v ...interface{}) {
    logrus.Warning(v...)
}

func Error(v ...interface{}) {
    logrus.Error(v...)
}


func Println(v ...interface{}) {
    logrus.Println(v...)
}

func Fatalln( v ...interface{}) {
    logrus.Fatalln(v...)
}

func Panicln(v ...interface{}) {
    logrus.Panicln(v...)
}

func Debugln(v ...interface{}) {
    logrus.Debugln(v...)
}

func Warnln(v ...interface{}) {
    logrus.Warnln(v...)
}

func Warningln(v ...interface{}) {
    logrus.Warningln(v...)
}

func Errorln(v ...interface{}) {
    logrus.Errorln(v...)
}
