package Core

type CppHeader struct {
	UnknownDerivatives []*string
	Defines            []*DefineDef
	Includes           []*IncludeDef

	Classes []*ClassDef
	Structs []*StructDef
}

type DefineDef struct {
	Key   string
	Value string
}

type IncludeDef struct {
	Path string
}

type ClassDef struct {
	DllMod string
	Name   string
	Body   *ObjectBody
}

type StructDef struct {
	DllMod string
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
