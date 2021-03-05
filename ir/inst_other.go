package ir

import (
	"fmt"
	"strings"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstICmp is an LLVM IR icmp instruction.
type InstICmp struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Integer comparison predicate.
	Pred IPred
	// Integer scalar or vector operands.
	X, Y Value // integer scalar, pointer, integer vector or pointer vector.

	// extra.

	// Type of result produced by the instruction.
	Typ Type // boolean or boolean vector
}

// NewICmp returns a new icmp instruction based on the given integer comparison
// predicate and integer scalar or vector operands.
func NewICmp(pred IPred, x, y Value) *InstICmp {
	inst := &InstICmp{Pred: pred, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstICmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction. The result type is either boolean
// type or vector of booleans type.
func (inst *InstICmp) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *IntType, *PointerType:
			inst.Typ = I1
		case *VectorType:
			inst.Typ = NewVectorType(xType.Len, I1)
		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *IntType, *PointerType or *VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'icmp' Pred=IPred X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstICmp) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "icmp %s %s, %s", inst.Pred, inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFCmp is an LLVM IR fcmp instruction.
type InstFCmp struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Floating-point comparison predicate.
	Pred FPred
	// Floating-point scalar or vector operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type // boolean or boolean vector
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// comparison predicate and floating-point scalar or vector operands.
func NewFCmp(pred FPred, x, y Value) *InstFCmp {
	inst := &InstFCmp{Pred: pred, X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction. The result type is either boolean
// type or vector of booleans type.
func (inst *InstFCmp) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *FloatType:
			inst.Typ = I1
		case *VectorType:
			inst.Typ = NewVectorType(xType.Len, I1)
		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *FloatType or *VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fcmp' FastMathFlags=FastMathFlag* Pred=FPred X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFCmp) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fcmp")
	fmt.Fprintf(buf, " %s %s, %s", inst.Pred, inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPhi is an LLVM IR phi instruction.
type InstPhi struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Incoming values.
	Incs []*Incoming

	// extra.

	// Type of result produced by the instruction.
	Typ Type // type of incoming value
}

// NewPhi returns a new phi instruction based on the given incoming values.
func NewPhi(incs ...*Incoming) *InstPhi {
	inst := &InstPhi{Incs: incs}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstPhi) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction. The result type is the type of the
// incoming
func (inst *InstPhi) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.Incs[0].X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'phi' Typ=Type Incs=(Inc separator ',')+ Metadata=(',' MetadataAttachment)+?
func (inst *InstPhi) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("phi ")
	buf.WriteString(inst.Typ.String())
	buf.WriteString(" ")
	for i, inc := range inst.Incs {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(inc.String())
	}
	return buf.String()
}

// ___ [ Incoming value ] ______________________________________________________

// Incoming is an incoming value of a phi instruction.
type Incoming struct {
	// Incoming
	X Value
	// Predecessor basic block of the incoming
	Pred Value // *ir.Block
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x Value, pred Value) *Incoming {
	return &Incoming{X: x, Pred: pred}
}

// String returns the string representation of the incoming
func (inc *Incoming) String() string {
	// '[' X=Value ',' Pred=LocalIdent ']'
	return fmt.Sprintf("[ %s, %s ]", inc.X.Ident(), inc.Pred.Ident())
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSelect is an LLVM IR select instruction.
type InstSelect struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Selection condition.
	Cond Value // boolean or boolean vector
	// True condition
	ValueTrue Value
	// False condition
	ValueFalse Value

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewSelect returns a new select instruction based on the given selection
// condition and true and false condition values.
func NewSelect(cond, valueTrue, valueFalse Value) *InstSelect {
	inst := &InstSelect{Cond: cond, ValueTrue: valueTrue, ValueFalse: valueFalse}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSelect) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSelect) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.ValueTrue.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'select' FastMathFlags=FastMathFlag* Cond=TypeValue ',' ValueTrue=TypeValue ',' ValueFalse=TypeValue Metadata=(',' MetadataAttachment)+?
func (inst *InstSelect) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("select")
	fmt.Fprintf(buf, " %s, %s, %s", inst.Cond, inst.ValueTrue, inst.ValueFalse)
	return buf.String()
}

// ~~~ [ freeze ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFreeze is an LLVM IR freeze instruction.
type InstFreeze struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operand.
	X Value
	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewInstFreeze returns a new freeze instruction based on the given
// operand.
func NewInstFreeze(x Value) *InstFreeze {
	inst := &InstFreeze{X: x}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFreeze) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFreeze) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'freeze' Type Value
func (inst *InstFreeze) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "freeze %s", inst.X)
	return buf.String()
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCall is an LLVM IR call instruction.
type InstCall struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Callee.
	Callee Value
	// Function arguments.
	Args []Value
	// Type of result produced by the instruction.
	Typ Type
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
//
// TODO: specify the set of underlying types of callee.
func NewCall(callee Value, args ...Value) *InstCall {
	inst := &InstCall{Callee: callee, Args: args}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCall) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCall) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		sig := inst.Sig()
		inst.Typ = sig.RetType
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// Tailopt 'call' FastMathFlags=FastMathFlag* CallingConvopt ReturnAttrs=ReturnAttribute* AddrSpaceopt Typ=Type Callee=Value '(' Args ')' FuncAttrs=FuncAttribute* OperandBundles=('[' (OperandBundle separator ',')+ ']')? Metadata=(',' MetadataAttachment)+?
func (inst *InstCall) LLString() string {
	buf := &strings.Builder{}
	if !inst.Type().Equal(Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	buf.WriteString("call")
	// Use function signature instead of return type for variadic functions.
	calleeType := inst.Type()
	if sig := inst.Sig(); sig.Variadic {
		calleeType = sig
	}
	fmt.Fprintf(buf, " %s %s(", calleeType, inst.Callee.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")
	return buf.String()
}

// Sig returns the function signature of the callee.
func (inst *InstCall) Sig() *FuncType {
	t, ok := inst.Callee.Type().(*PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *PointerType, got %T", inst.Callee.Type()))
	}
	sig, ok := t.ElemType.(*FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *FuncType, got %T", t.ElemType))
	}
	return sig
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstVAArg is an LLVM IR va_arg instruction.
type InstVAArg struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Variable argument list.
	ArgList Value
	// Argument type.
	ArgType Type

	// extra.
}

// NewVAArg returns a new va_arg instruction based on the given variable
// argument list and argument type.
func NewVAArg(argList Value, argType Type) *InstVAArg {
	return &InstVAArg{ArgList: argList, ArgType: argType}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstVAArg) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstVAArg) Type() Type {
	return inst.ArgType
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'va_arg' ArgList=TypeValue ',' ArgType=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstVAArg) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "va_arg %s, %s", inst.ArgList, inst.ArgType)
	return buf.String()
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLandingPad is an LLVM IR landingpad instruction.
type InstLandingPad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Result type.
	ResultType Type
	// Filter and catch clauses; zero or more if Cleanup is true, otherwise one
	// or more.
	Clauses []*Clause

	// extra.
}

// NewLandingPad returns a new landingpad instruction based on the given result
// type and filter/catch clauses.
func NewLandingPad(resultType Type, clauses ...*Clause) *InstLandingPad {
	return &InstLandingPad{ResultType: resultType, Clauses: clauses}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLandingPad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLandingPad) Type() Type {
	return inst.ResultType
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'landingpad' ResultType=Type Cleanupopt Clauses=Clause* Metadata=(',' MetadataAttachment)+?
func (inst *InstLandingPad) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "landingpad %s", inst.ResultType)
	for _, clause := range inst.Clauses {
		fmt.Fprintf(buf, "\n\t\t%s", clause)
	}
	return buf.String()
}

// ___ [ Landingpad clause ] ___________________________________________________

// Clause is a landingpad catch or filter clause.
type Clause struct {
	// Clause type (catch or filter).
	Type ClauseType
	// Operand.
	X Value
}

// NewClause returns a new landingpad clause based on the given clause type and
// operand.
func NewClause(clauseType ClauseType, x Value) *Clause {
	return &Clause{Type: clauseType, X: x}
}

// String returns the string representation of the landingpad clause.
func (clause *Clause) String() string {
	return fmt.Sprintf("%d %s", clause.Type, clause.X)
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCatchPad is an LLVM IR catchpad instruction.
type InstCatchPad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Parent catchswitch terminator.
	CatchSwitch Value // *ir.TermCatchSwitch
	// Exception arguments.
	//
	// Arg has one of the following underlying types:
	//    Value
	//    TODO: add metadata value?
	Args []Value

	// extra.
}

// NewCatchPad returns a new catchpad instruction based on the given parent
// catchswitch terminator and exception arguments.
func NewCatchPad(catchSwitch Value, args ...Value) *InstCatchPad {
	return &InstCatchPad{CatchSwitch: catchSwitch, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCatchPad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCatchPad) Type() Type {
	return Token
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'catchpad' 'within' CatchSwitch=LocalIdent '[' Args=(ExceptionArg separator ',')* ']' Metadata=(',' MetadataAttachment)+?
func (inst *InstCatchPad) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "catchpad within %s [", inst.CatchSwitch.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString("]")
	return buf.String()
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCleanupPad is an LLVM IR cleanuppad instruction.
type InstCleanupPad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Parent exception pad.
	ParentPad Value // ir.ExceptionPad
	// Exception arguments.
	//
	// Arg has one of the following underlying types:
	//    Value
	//    TODO: add metadata value?
	Args []Value

	// extra.
}

// NewCleanupPad returns a new cleanuppad instruction based on the given
// parent exception pad and exception arguments.
func NewCleanupPad(parentPad Value, args ...Value) *InstCleanupPad {
	return &InstCleanupPad{ParentPad: parentPad, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCleanupPad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCleanupPad) Type() Type {
	return Token
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'cleanuppad' 'within' ParentPad=ExceptionPad '[' Args=(ExceptionArg separator ',')* ']' Metadata=(',' MetadataAttachment)+?
func (inst *InstCleanupPad) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "cleanuppad within %s [", inst.ParentPad.Ident())
	for i, arg := range inst.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString("]")
	return buf.String()
}
