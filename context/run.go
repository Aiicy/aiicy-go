package context

import (
	"flag"
	"github.com/aiicy/aiicy-go/logger"
	"os"
	"runtime/debug"
)

// Run service
func Run(handle func(Context) error) {
	var h bool
	var c string
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&c, "c", DefaultConfFile, "the configuration file")
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	ctx := NewContext(c)
	defer func() {
		if r := recover(); r != nil {
			ctx.Log().Error("service is stopped with panic", logger.String("panic", string(debug.Stack())))
		}
	}()

	pwd, _ := os.Getwd()
	ctx.Log().Info("service starting", logger.Any("args", os.Args), logger.String("pwd", pwd))
	err := handle(ctx)
	if err != nil {
		ctx.Log().Error("service has stopped with error", logger.Error(err))
	} else {
		ctx.Log().Info("service has stopped")
	}
}
