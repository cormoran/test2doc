package widgets

import (
	"testing"

	"github.com/adams-sarah/prettytest"
	"github.com/cormoran/test2doc/test"
	"github.com/cormoran/test2doc/vars"
	"github.com/labstack/echo"
)

var server *test.Server

type mainSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	var err error
	e := echo.New()
	AddRoutes(e)

	test.RegisterURLVarExtractor(vars.MakeEchoExtractor(e))
	test.SetHandlerInfoFunc(vars.MakeEchoGetHandlerInfoFunc(e))

	server, err = test.NewServer(e, &test.Config{
		PackageDir: ".",
		OutputDir:  ".",
	})
	if err != nil {
		panic(err.Error())
	}
	defer server.Finish()

	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(mainSuite),
	)
}
