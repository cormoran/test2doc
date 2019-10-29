package test

import (
	"log"
	"runtime"

	"github.com/cormoran/test2doc/doc/parse"
)

type GetHandlerInfoFuncType func() HandlerInfo

var customHandlerInfoFunc GetHandlerInfoFuncType

func SetHandlerInfoFunc(f GetHandlerInfoFuncType) {
	customHandlerInfoFunc = f
}

func GetHandlerInfo() HandlerInfo {
	if customHandlerInfoFunc != nil {
		return customHandlerInfoFunc()
	}
	return DefaultGetHandlerInfoFunc()
}

func DefaultGetHandlerInfoFunc() HandlerInfo {
	i := 1
	max := 15

	var pc uintptr
	var file, fnName string
	var ok, fnInPkg, sawPkg bool

	// iterate until we find the top level func in this pkg (the handler)
	for i < max {
		pc, file, _, ok = runtime.Caller(i)
		if !ok {
			log.Println("test2doc: setHandlerInfo: !ok")
			return HandlerInfo{
				FileName: "",
				FuncName: "",
			}
		}

		fn := runtime.FuncForPC(pc)
		fnName = fn.Name()

		fnInPkg = parse.IsFuncInPkg(fnName)
		if sawPkg && !fnInPkg {
			pc, file, _, ok = runtime.Caller(i - 1)
			fn := runtime.FuncForPC(pc)
			fnName = fn.Name()
			break
		}

		sawPkg = fnInPkg
		i++
	}

	return HandlerInfo{
		FileName: file,
		FuncName: fnName,
	}
}
