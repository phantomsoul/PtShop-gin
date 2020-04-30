package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

// error logger
var logger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// 格式化时间
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 初始化日志 logger
func InitLog(logPath string, level string) {
	infoLogFile := logPath + "/info.log"
	errorLogFile := logPath + "/error.log"
	if !Exists(logPath) {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			fmt.Println("mkdir logPath err!")
			return
		}
		iFile, iErr := os.Create(infoLogFile)
		eFile, eErr := os.Create(errorLogFile)
		defer iFile.Close()
		defer eFile.Close()
		if iErr != nil {
			fmt.Println("create info log file err!")
			fmt.Println(iErr)
			return
		}
		if eErr != nil {
			fmt.Println("create error log file err!")
			fmt.Println(eErr)
			return
		}
	}

	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	config := zapcore.EncoderConfig{
		TimeKey:       "T",                         // json时时间键
		LevelKey:      "L",                         // json时日志等级键
		NameKey:       "N",                         // json时日志记录器键
		CallerKey:     "C",                         // json时日志文件信息键
		MessageKey:    "M",                         // json时日志消息键
		StacktraceKey: "S",                         // json时堆栈键
		LineEnding:    zapcore.DefaultLineEnding,   // 友好日志换行符
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 友好日志等级名大小写（info INFO）
		EncodeTime:    TimeEncoder,                 // 友好日志时日期格式化
		//EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller: zapcore.ShortCallerEncoder, // 日志文件信息（包/文件.go:行号）
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	jsonEncoder := zapcore.NewJSONEncoder(config)

	// 实现两个判断日志等级的interface  可以自定义级别展示
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer抽象
	infoWriter := getWriter(infoLogFile)
	warnWriter := getWriter(errorLogFile)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		// 将info及以下写入logPath,  warn及以上写入errPath
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(warnWriter), warnLevel),
		//日志都会在console中展示
		zapcore.NewCore(consoleEncoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), getLoggerLevel(level)),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 跳过自身打印逻辑错误
}

//日志文件切割
func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

//查看文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Sugar(msg string, fields ...interface{}) {
	sugar := logger.Sugar()
	sugar.Infof(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}
