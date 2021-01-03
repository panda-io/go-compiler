package ir

import (
	"fmt"
)

func Size(t Type) int {
	switch t := t.(type) {
	case *ArrayType, *VoidType, *FuncType, *LabelType, *PointerType:
		return 8

	case *StructType:
		return StructSize(t)

	case *IntType:
		return IntSize(t)

	case *FloatType:
		return FloatSize(t)

	case *VectorType:
		return VectorSize(t)

	default:
		panic(fmt.Errorf("support for type %T not yet implemented", t))
	}
}

func StructSize(t *StructType) int {
	var size int
	for _, ty := range t.Fields {
		size += Size(ty)
	}
	return size
}

func IntSize(t *IntType) int {
	return int(t.BitSize) / 8
}

func FloatSize(t *FloatType) int {
	switch t.Kind {
	case FloatKindFloat:
		return 4

	case FloatKindDouble:
		return 8

	default:
		panic(fmt.Errorf("support for floating-point kind %q not yet implemented", t.Kind))
	}
}

func VectorSize(t *VectorType) int {
	return Size(t.ElemType) * int(t.Len)
}
