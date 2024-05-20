package Core

type CppHeader struct {
	Defines  []*DefineDef  `parser:"@@"`
	Includes []*IncludeDef `parser:"@@"`

	Classes []*ClassDef  `parser:"@@"`
	Structs []*StructDef `parser:"@@"`

	UnknownDerivatives []*string
}

type DefineDef struct {
	Key   string `parser:"'#define @'"`
	Value string
}

type IncludeDef struct {
	Path string `parser:"'#include (<)@(>)'"`
}

type ClassDef struct {
	DllMod string `parser:"'class @'"`
	Name   string
	Body   *ObjectBody
}

type StructDef struct {
	DllMod string `parser:"'struct @'"`
	Name   string
	Body   *ObjectBody
}

type ObjectBody struct {
	Variables []*VariableDef
	Functions []*FunctionDef
}

type VariableDef struct {
	Type LinkedType
	Name string
}

type LinkedType struct {
	IsConst    bool `parser:"const"`
	IsRef      bool
	PointerNum uint8
	IsRVal     bool
	TypeName   string
	Type       *LinkedType
}

type FunctionDef struct {
	IsStatic   bool           `parser:"'static'"`
	IsVirtual  bool           `parser:"| 'virtual'"`
	ReturnType *LinkedType    `parser:"@@"`
	Name       string         `parser:"@Ident"`
	Arguments  []*VariableDef `parser:"'(' @@* '[,)]'"`
	IsConst    bool           `parser:"'const'"`
	IsOverride bool           `parser:"'override'"`
}
