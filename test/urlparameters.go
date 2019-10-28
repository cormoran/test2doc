package test

import "github.com/cormoran/test2doc/doc/parse"

func RegisterURLVarExtractor(fn parse.URLVarExtractor) {
	parse.SetURLVarExtractor(&fn)
}
