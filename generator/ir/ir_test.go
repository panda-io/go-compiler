package ir

import (
	"fmt"
	"testing"
)

// This example produces LLVM IR code equivalent to the following C code, which
// implements a pseudo-random number generator.
//
//    int abs(int x);
//
//    int seed = 0;
//
//    // ref: https://en.wikipedia.org/wiki/Linear_congruential_generator
//    //    a = 0x15A4E35
//    //    c = 1
//    int rand(void) {
//       seed = seed*0x15A4E35 + 1;
//       return abs(seed);
//    }

func TestIR(t *testing.T) {
	// Create convenience types and constants.
	zero := NewInt(I32, 0)
	a := NewInt(I32, 0x15A4E35)
	c := NewInt(I32, 1)

	// Create a new LLVM IR module.
	m := NewModule()

	// Create an external function declaration and append it to the module.
	//
	//    int abs(int x);
	abs := m.NewFunc("abs", I32, NewParam(I32))

	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seed := m.NewGlobalDef("seed", zero)

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	rand := m.NewFunc("rand", I32)

	// Create an unnamed entry basic block and append it to the `rand` function.
	entry := rand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	tmp1 := entry.NewLoad(I32, seed)
	tmp2 := entry.NewMul(tmp1, a)
	tmp3 := entry.NewAdd(tmp2, c)
	entry.NewStore(tmp3, seed)
	tmp4 := entry.NewCall(abs, tmp3)
	entry.NewRet(tmp4)

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)

	t.Fail()
}
