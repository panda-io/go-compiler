package ir

import (
	"fmt"
	"testing"
)

// This example produces LLVM IR code equivalent to the following C code, which
// implements a pseudo-random number generator.
//
//    int abs(int x);
//    int seed = 0;
//    // ref: https://en.wikipedia.org/wiki/Linear_congruential_generator
//    //    a = 0x15A4E35
//    //    c = 1
//    int rand(void) {
//       seed = seed*0x15A4E35 + 1;
//       return abs(seed);
//    }
//    int main(void) {
//       print(rand());
//		 return 0;
//	  }

func TestIR(t *testing.T) {
	// Create convenience types and constants.
	zero := NewInt(I32, 0)
	a := NewInt(I32, 0x15A4E35)
	c := NewInt(I32, 1)

	// Create a new LLVM IR module.
	m := NewModule()

	// Create an external function declaration and append it to the module.
	// int abs(int x);
	abs := m.NewFunc("abs", I32, NewParam(I32))
	seed := m.NewGlobalDef("seed", zero)

	rand := m.NewFunc("rand", I32)
	entry := rand.NewBlock("")
	tmp1 := NewLoad(I32, seed)
	entry.AddInstruction(tmp1)
	tmp2 := NewMul(tmp1, a)
	entry.AddInstruction(tmp2)
	tmp3 := NewAdd(tmp2, c)
	entry.AddInstruction(tmp3)
	tmp4 := NewStore(tmp3, seed)
	entry.AddInstruction(tmp4)
	tmp5 := NewCall(abs, tmp3)
	entry.AddInstruction(tmp5)
	entry.Term = NewRet(tmp5)

	fmt.Println(m)

	t.Fail()
}

/*******
@seed = global i32 0

declare i32 @abs(i32 %x)

define i32 @rand() {
0:
	%1 = load i32, i32* @seed
	%2 = mul i32 %1, 22695477
	%3 = add i32 %2, 1
	store i32 %3, i32* @seed
	%4 = call i32 @abs(i32 %3)
	ret i32 %4
}
**************/

func TestAdd(t *testing.T) {
	// Create a new LLVM IR module.
	m := NewModule()

	funcAdd := m.NewFunc("add", I32,
		NewParam(I32),
		NewParam(I32),
	)
	addBlock := funcAdd.NewBlock("")
	tmp1 := NewAdd(funcAdd.Params[0], funcAdd.Params[1])
	addBlock.AddInstruction(tmp1)
	addBlock.Term = NewRet(tmp1)

	funcMain := m.NewFunc(
		"main",
		I32,
	) // omit parameters
	mainBlock := funcMain.NewBlock("")
	tmp2 := NewCall(funcAdd, NewInt(I32, 1), NewInt(I32, 2))
	mainBlock.AddInstruction(tmp2)
	mainBlock.Term = NewRet(tmp2)

	fmt.Println(m)

	t.Fail()
}
