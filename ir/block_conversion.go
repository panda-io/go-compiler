package ir

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewTrunc appends a new trunc instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewTrunc(from Value, to Type) *InstTrunc {
	inst := NewTrunc(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewZExt appends a new zext instruction to the basic block based on the given
// source value and target type.
func (block *Block) NewZExt(from Value, to Type) *InstZExt {
	inst := NewZExt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSExt appends a new sext instruction to the basic block based on the given
// source value and target type.
func (block *Block) NewSExt(from Value, to Type) *InstSExt {
	inst := NewSExt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPTrunc appends a new fptrunc instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewFPTrunc(from Value, to Type) *InstFPTrunc {
	inst := NewFPTrunc(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPExt appends a new fpext instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewFPExt(from Value, to Type) *InstFPExt {
	inst := NewFPExt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPToUI appends a new fptoui instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewFPToUI(from Value, to Type) *InstFPToUI {
	inst := NewFPToUI(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPToSI appends a new fptosi instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewFPToSI(from Value, to Type) *InstFPToSI {
	inst := NewFPToSI(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewUIToFP appends a new uitofp instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewUIToFP(from Value, to Type) *InstUIToFP {
	inst := NewUIToFP(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSIToFP appends a new sitofp instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewSIToFP(from Value, to Type) *InstSIToFP {
	inst := NewSIToFP(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewPtrToInt appends a new ptrtoint instruction to the basic block based on
// the given source value and target type.
func (block *Block) NewPtrToInt(from Value, to Type) *InstPtrToInt {
	inst := NewPtrToInt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewIntToPtr appends a new inttoptr instruction to the basic block based on
// the given source value and target type.
func (block *Block) NewIntToPtr(from Value, to Type) *InstIntToPtr {
	inst := NewIntToPtr(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewBitCast appends a new bitcast instruction to the basic block based on the
// given source value and target type.
func (block *Block) NewBitCast(from Value, to Type) *InstBitCast {
	inst := NewBitCast(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAddrSpaceCast appends a new addrspacecast instruction to the basic block
// based on the given source value and target type.
func (block *Block) NewAddrSpaceCast(from Value, to Type) *InstAddrSpaceCast {
	inst := NewAddrSpaceCast(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}
