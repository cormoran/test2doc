package vars

import (
	"log"
	"net/http"
	"runtime"

	"github.com/cormoran/test2doc/doc/parse"
	"github.com/cormoran/test2doc/test"
	"github.com/labstack/echo"
)

func MakeEchoExtractor(e *echo.Echo) parse.URLVarExtractor {
	return func(req *http.Request) map[string]string {
		ctx := e.AcquireContext()
		defer e.ReleaseContext(ctx)
		pnames := ctx.ParamNames()
		if len(pnames) == 0 {
			return nil
		}

		paramsMap := make(map[string]string, len(pnames))
		for _, name := range pnames {
			paramsMap[name] = ctx.Param(name)
		}
		return paramsMap
	}
}

func MakeEchoGetHandlerInfoFunc(e *echo.Echo) test.GetHandlerInfoFuncType {
	handlerSet := make(map[string]struct{})
	for _, route := range e.Routes() {
		handlerSet[route.Name] = struct{}{}
	}
	return func() test.HandlerInfo {
		i := 1
		max := 15

		var pc uintptr
		var file, fnName string
		var ok bool

		// iterate until we find the top level func in this pkg (the handler)
		for i < max {
			pc, file, _, ok = runtime.Caller(i)
			if !ok {
				log.Println("test2doc: setHandlerInfo: !ok")
				return test.HandlerInfo{
					FileName: "",
					FuncName: "",
				}
			}

			fn := runtime.FuncForPC(pc)
			fnName = fn.Name()

			if _, ok := handlerSet[fnName]; ok {
				break
			}
			i++
		}
		return test.HandlerInfo{
			FileName: file,
			FuncName: fnName,
		}
	}
}
