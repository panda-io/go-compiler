// The basic block API of the ir package was heavily inspired by
// https://github.com/mrbenshef/goory.

package ir

import (
	"fmt"
	"strings"
)

// === [ Basic blocks ] ========================================================

// Block is an LLVM IR basic block; a sequence of non-branching instructions
// terminated by a control flow instruction (e.g. br or ret).
type Block struct {
	// Name of local variable associated with the basic block.
	LocalIdent
	// Instructions of the basic block.
	Insts      []Instruction
	Terminated bool
}

// NewBlock returns a new basic block based on the given label name. An empty
// label name indicates an unnamed basic block.
func NewBlock(name string) *Block {
	block := &Block{}
	block.SetName(name)
	return block
}

func (block *Block) AddInstruction(inst Instruction) {
	if block.Terminated {
		panic("block already terminated")
	}
	block.Insts = append(block.Insts, inst)
	if _, ok := inst.(Terminator); ok {
		block.Terminated = true
	}
}

func (block *Block) InsertAlloca(inst *InstAlloca) {
	if block.Terminated {
		block.Insts = append(block.Insts, block.Insts[len(block.Insts)-1])
		block.Insts[len(block.Insts)-2] = inst
	} else {
		block.Insts = append(block.Insts, inst)
	}
}

// String returns the LLVM syntax representation of the basic block as a
// type-value pair.
func (block *Block) String() string {
	return fmt.Sprintf("%s %s", block.Type(), block.Ident())
}

// Type returns the type of the basic block.
func (block *Block) Type() Type {
	return Label
}

// LLString returns the LLVM syntax representation of the basic block
// definition.
//
// Name=LabelIdentopt Insts=Instruction* Term=Terminator
func (block *Block) LLString() string {
	buf := &strings.Builder{}
	if block.IsUnnamed() {
		//fmt.Fprintf(buf, "; <label>:%d\n", block.LocalID)
		// Explicitly print basic block label to conform with Clang 9.0, and
		// because it's the sane thing to do.
		fmt.Fprintf(buf, "%s\n", LabelID(block.LocalID))
	} else {
		fmt.Fprintf(buf, "%s\n", LabelName(block.LocalName))
	}
	for _, inst := range block.Insts {
		fmt.Fprintf(buf, "\t%s\n", inst.LLString())
	}
	return buf.String()
}
