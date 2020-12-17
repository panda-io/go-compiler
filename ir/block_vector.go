package ir

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewExtractElement appends a new extractelement instruction to the basic block
// based on the given vector and element index.
func (block *Block) NewExtractElement(x, index Value) *InstExtractElement {
	inst := NewExtractElement(x, index)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInsertElement appends a new insertelement instruction to the basic block
// based on the given vector, element and element index.
func (block *Block) NewInsertElement(x, elem, index Value) *InstInsertElement {
	inst := NewInsertElement(x, elem, index)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShuffleVector appends a new shufflevector instruction to the basic block
// based on the given vectors and shuffle mask.
func (block *Block) NewShuffleVector(x, y, mask Value) *InstShuffleVector {
	inst := NewShuffleVector(x, y, mask)
	block.Insts = append(block.Insts, inst)
	return inst
}
