//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L *zap.Logger

type Logger = zap.Logger

type Field = zap.Field

// Level log level
type Level = zapcore.Level

// all log level
const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// Any constructs a field with the given key and value
func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}

func Error(err error) Field {
	return zap.Error(err)
}

func String(key string, val string) Field {
	return zap.String(key, val)
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func With(fields ...Field) *Logger {
	return zap.L().With(fields...)
}

func init() {
	// Print log when start
	L = New(Config{Level: "debug"})
}

// NewEncoderConfig creates logger config for debug mode
func NewEncoderConfig() zapcore.EncoderConfig {
	conf := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return conf
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// ParseLevel parses string to zap level
func ParseLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	default:
		return zap.WarnLevel, fmt.Errorf("level %s not exist", level)
	}
}

// New create a new Sugared logger
func New(c Config, fields ...Field) *zap.Logger {
	var (
		format zapcore.Encoder
		write  zapcore.WriteSyncer
	)
	logLevel, err := ParseLevel(c.Level)
	if err != nil {
		L.Warn("failed to parse log level, use default level (info)", String("level", c.Level))
	}

	if c.Format == "json" {
		format = zapcore.NewJSONEncoder(NewEncoderConfig())
	} else {
		format = zapcore.NewConsoleEncoder(NewEncoderConfig())
	}

	if c.Encoding == "json" {
		write = zapcore.AddSync(&lumberjack.Logger{
			Filename:   c.Path,
			MaxAge:     c.MaxAge,  //days
			MaxSize:    c.MaxSize, // megabytes
			MaxBackups: c.MaxBackups,
		})
	} else {
		write = os.Stdout
	}
	core := zapcore.NewCore(
		format,
		write,
		logLevel,
	)
	var options []zap.Option
	if len(fields) > 0 {
		options = append(options, zap.Fields(fields...))
	}
	if logLevel == zap.DebugLevel {
		options = append(options, zap.AddCaller())
	}
	return zap.New(core, options...)
}
