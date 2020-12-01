package types

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/panda-foundation/go-compiler/token"
)

func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if isNil(a) && isNil(b) {
		return
	}
	if a != b {
		panic(fmt.Sprintf("expected %v, but got %v", b, a))
	}
}

func TestTypeEqual(t *testing.T) {
	t0 := &BuitinType{Token: token.Int32}
	t1 := &BuitinType{Token: token.Int32}
	t2 := &TypeName{}
	t2.QualifiedName = []string{"name"}
	t3 := &TypeName{}
	t3.QualifiedName = []string{"name1"}
	t4 := &TypeName{}
	assertEqual(t, t0.Equal(t1), true)
	assertEqual(t, t0.Equal(t2), false)
	assertEqual(t, t2.Equal(t0), false)
	assertEqual(t, t2.Equal(t2), true)
	assertEqual(t, t2.Equal(t3), false)
	assertEqual(t, t2.Equal(t4), false)

	t5 := &TypeName{}
	t5.TypeArguments = &TypeArguments{Ellipsis: -1}
	t6 := &TypeName{}
	t6.TypeArguments = &TypeArguments{Ellipsis: 1}
	t7 := &TypeName{}
	assertEqual(t, t5.Equal(t6), false)
	assertEqual(t, t5.Equal(t7), false)

	t8 := &TypeArguments{Arguments: []Type{&BuitinType{Token: token.Int32}}}
	t9 := &TypeArguments{}
	t10 := &TypeArguments{Arguments: []Type{&BuitinType{Token: token.Int64}}}
	assertEqual(t, t8.Equal(t9), false)
	assertEqual(t, t8.Equal(t10), false)
	assertEqual(t, t8.Equal(t8), true)
	assertEqual(t, t8.Equal(t0), false)

	t11 := &TypeParameters{Ellipsis: true}
	t12 := &TypeParameters{}
	t13 := &TypeParameters{Parameters: []*TypeParameter{&TypeParameter{Name: "n", Type: &BuitinType{Token: token.Int32}}}}
	t14 := &TypeParameters{Parameters: []*TypeParameter{&TypeParameter{Name: "n", Type: &BuitinType{Token: token.Int64}}}}
	assertEqual(t, t11.Equal(t12), false)
	assertEqual(t, t13.Equal(t13), true)
	assertEqual(t, t13.Equal(t14), false)
	assertEqual(t, t12.Equal(t10), false)
	assertEqual(t, t12.Equal(t13), false)

	t15 := &TypeParameter{Name: "n", Type: &BuitinType{Token: token.Int32}}
	t16 := &TypeParameter{Name: "n", Type: &BuitinType{Token: token.Int64}}
	t17 := &TypeParameter{Name: "m", Type: &BuitinType{Token: token.Int32}}
	t18 := &TypeParameter{Name: "m"}
	assertEqual(t, t15.Equal(t16), false)
	assertEqual(t, t15.Equal(t15), true)
	assertEqual(t, t15.Equal(t14), false)
	assertEqual(t, t15.Equal(t17), false)
	assertEqual(t, t18.Equal(t17), false)
	assertEqual(t, t18.Equal(t18), true)

	assertEqual(t, (&Arguments{}).Equal(&Arguments{}), false)

	t20 := &Parameter{Name: "n", Type: &BuitinType{Token: token.String}}
	t21 := &Parameter{Name: "m"}
	t22 := &Parameter{Name: "n", Type: &BuitinType{Token: token.Int16}}
	assertEqual(t, t20.Equal(t20), true)
	assertEqual(t, t20.Equal(t21), false)
	assertEqual(t, t20.Equal(t22), false)
	assertEqual(t, t20.Equal(t18), false)

	t23 := &Base{}
	t24 := &Base{}
	assertEqual(t, t23.Equal(t24), false)

	t26 := &Parameters{Ellipsis: true}
	t27 := &Parameters{}
	t28 := &Parameters{Parameters: []*Parameter{&Parameter{Name: "m", Type: &BuitinType{Token: token.Int16}}}}
	t29 := &Parameters{Parameters: []*Parameter{&Parameter{Name: "n", Type: &BuitinType{Token: token.Int16}}}}
	assertEqual(t, t26.Equal(t27), false)
	assertEqual(t, t26.Equal(t26), true)
	assertEqual(t, t28.Equal(t28), true)
	assertEqual(t, t28.Equal(t29), false)
	assertEqual(t, t27.Equal(t28), false)
	assertEqual(t, t27.Equal(t22), false)
}
