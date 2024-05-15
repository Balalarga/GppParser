package main

import (
	"GppParser/Core"
	"github.com/alecthomas/participle/v2"
)

func main() {
	codeExample := "#include \"variant\"" +
		"#include <vector>"
	parser := participle.Build[Core.CppHeader]()

}
