package main

import (
	"GppParser/Core"
	"github.com/alecthomas/participle/v2"
	"log"
)

func main() {
	codeExample := "#include \"variant\"" +
		"#include <vector>"
	parser := participle.MustBuild[Core.CppHeader]()
	ast, err := parser.ParseString("", codeExample)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(ast)
}
