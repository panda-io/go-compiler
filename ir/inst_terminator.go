package ir

import (
	"fmt"
	"strings"
)

// === [ Terminators ] =========================================================

// Terminator is an LLVM IR terminator instruction (a control flow instruction).
//
// A Terminator has one of the following underlying
//
// Terminators
//
// https://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ir.TermRet           // https://godoc.org/github.com/llir/llvm/ir#TermRet
//    *ir.TermBr            // https://godoc.org/github.com/llir/llvm/ir#TermBr
//    *ir.TermCondBr        // https://godoc.org/github.com/llir/llvm/ir#TermCondBr
//    *ir.TermSwitch        // https://godoc.org/github.com/llir/llvm/ir#TermSwitch
//    *ir.TermIndirectBr    // https://godoc.org/github.com/llir/llvm/ir#TermIndirectBr
//    *ir.TermInvoke        // https://godoc.org/github.com/llir/llvm/ir#TermInvoke
//    *ir.TermCallBr        // https://godoc.org/github.com/llir/llvm/ir#TermCallBr
//    *ir.TermResume        // https://godoc.org/github.com/llir/llvm/ir#TermResume
//    *ir.TermCatchSwitch   // https://godoc.org/github.com/llir/llvm/ir#TermCatchSwitch
//    *ir.TermCatchRet      // https://godoc.org/github.com/llir/llvm/ir#TermCatchRet
//    *ir.TermCleanupRet    // https://godoc.org/github.com/llir/llvm/ir#TermCleanupRet
//    *ir.TermUnreachable   // https://godoc.org/github.com/llir/llvm/ir#TermUnreachable

type Terminator interface {
	isTerminator()
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet is an LLVM IR ret terminator.
type TermRet struct {
	// Return value; or nil if void return.
	X Value
}

// NewRet returns a new ret terminator based on the given return  A nil
// return value indicates a void return.
func NewRet(x Value) *TermRet {
	return &TermRet{X: x}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// Void return instruction.
//
//    'ret' XTyp=VoidType Metadata=(',' MetadataAttachment)+?
//
// Value return instruction.
//
//    'ret' XTyp=ConcreteType X=Value Metadata=(',' MetadataAttachment)+?
func (term *TermRet) LLString() string {
	buf := &strings.Builder{}
	if term.X == nil {
		buf.WriteString("ret void")
	} else {
		fmt.Fprintf(buf, "ret %s", term.X)
	}
	return buf.String()
}

// --- [ br ] ------------------------------------------------------------------

// TermBr is an unconditional LLVM IR br terminator.
type TermBr struct {
	// Target branch.
	Target Value // *ir.Block
}

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func NewBr(target Value) *TermBr {
	return &TermBr{Target: target}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'br' Target=Label Metadata=(',' MetadataAttachment)+?
func (term *TermBr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "br %s", term.Target)
	return buf.String()
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr is a conditional LLVM IR br terminator.
type TermCondBr struct {
	// Branching condition.
	Cond Value
	// True condition target branch.
	TargetTrue Value // *ir.Block
	// False condition target branch.
	TargetFalse Value // *ir.Block
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func NewCondBr(cond Value, targetTrue, targetFalse *Block) *TermCondBr {
	return &TermCondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'br' CondTyp=IntType Cond=Value ',' TargetTrue=Label ',' TargetFalse=Label Metadata=(',' MetadataAttachment)+?
func (term *TermCondBr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "br %s, %s, %s", term.Cond, term.TargetTrue, term.TargetFalse)
	return buf.String()
}

// --- [ switch ] --------------------------------------------------------------

// TermSwitch is an LLVM IR switch terminator.
type TermSwitch struct {
	// Control variable.
	X Value
	// Default target branch.
	TargetDefault Value // *ir.Block
	// Switch cases.
	Cases []*Case
}

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func NewSwitch(x Value, targetDefault *Block, cases ...*Case) *TermSwitch {
	return &TermSwitch{X: x, TargetDefault: targetDefault, Cases: cases}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'switch' X=TypeValue ',' Default=Label '[' Cases=Case* ']' Metadata=(',' MetadataAttachment)+?
func (term *TermSwitch) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "switch %s, %s [\n", term.X, term.TargetDefault)
	for _, c := range term.Cases {
		fmt.Fprintf(buf, "\t\t%s\n", c)
	}
	buf.WriteString("\t]")
	return buf.String()
}

// ~~~ [ Switch case ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Case is a switch case.
type Case struct {
	// Case comparand.
	X Value // Constant (integer constant or integer constant expression)
	// Case target branch.
	Target Value // *ir.Block
}

// NewCase returns a new switch case based on the given case comparand and
// target basic block.
func NewCase(x Constant, target *Block) *Case {
	return &Case{X: x, Target: target}
}

// String returns the string representation of the switch case.
func (c *Case) String() string {
	// X=TypeConst ',' Target=Label
	return fmt.Sprintf("%s, %s", c.X, c.Target)
}

// --- [ indirectbr ] ----------------------------------------------------------

// TermIndirectBr is an LLVM IR indirectbr terminator.
type TermIndirectBr struct {
	// Target address.
	Addr Value // blockaddress
	// Set of valid target basic blocks.
	ValidTargets []Value // slice of *ir.Block
}

// NewIndirectBr returns a new indirectbr terminator based on the given target
// address (derived from a blockaddress constant) and set of valid target basic
// blocks.
func NewIndirectBr(addr Constant, validTargets ...*Block) *TermIndirectBr {
	// convert validTargets slice to []Value.
	var targets []Value
	for _, target := range validTargets {
		targets = append(targets, target)
	}
	return &TermIndirectBr{Addr: addr, ValidTargets: targets}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'indirectbr' Addr=TypeValue ',' '[' ValidTargets=(Label separator ',')* ']' Metadata=(',' MetadataAttachment)+?
func (term *TermIndirectBr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "indirectbr %s, [", term.Addr)
	for i, target := range term.ValidTargets {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(target.String())
	}
	buf.WriteString("]")
	return buf.String()
}

// --- [ invoke ] --------------------------------------------------------------

// TermInvoke is an LLVM IR invoke terminator.
type TermInvoke struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Invokee (callee function).
	// TODO: specify the set of underlying types of Invokee.
	Invokee Value
	// Function arguments.
	//
	// Arg has one of the following underlying types:
	//    Value
	//    TODO: add metadata value?
	Args []Value
	// Normal control flow return point.
	NormalRetTarget Value // *ir.Block
	// Exception control flow return point.
	ExceptionRetTarget Value // *ir.Block

	// Type of result produced by the terminator.
	Typ Type
}

// NewInvoke returns a new invoke terminator based on the given invokee,
// function arguments and control flow return points for normal and exceptional
// execution.
//
// TODO: specify the set of underlying types of invokee.
func NewInvoke(invokee Value, args []Value, normalRetTarget, exceptionRetTarget *Block) *TermInvoke {
	term := &TermInvoke{Invokee: invokee, Args: args, NormalRetTarget: normalRetTarget, ExceptionRetTarget: exceptionRetTarget}
	// Compute type.
	term.Type()
	return term
}

// String returns the LLVM syntax representation of the terminator as a type-
// value pair.
func (term *TermInvoke) String() string {
	return fmt.Sprintf("%s %s", term.Type(), term.Ident())
}

// Type returns the type of the terminator.
func (term *TermInvoke) Type() Type {
	// Cache type if not present.
	if term.Typ == nil {
		sig := term.Sig()
		term.Typ = sig.RetType
	}
	return term.Typ
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'invoke' CallingConvopt ReturnAttrs=ReturnAttribute* AddrSpaceopt Typ=Type Invokee=Value '(' Args ')' FuncAttrs=FuncAttribute* OperandBundles=('[' (OperandBundle separator ',')+ ']')? 'to' NormalRetTarget=Label 'unwind' ExceptionRetTarget=Label Metadata=(',' MetadataAttachment)+?
func (term *TermInvoke) LLString() string {
	buf := &strings.Builder{}
	if !term.Type().Equal(Void) {
		fmt.Fprintf(buf, "%s = ", term.Ident())
	}
	buf.WriteString("invoke")
	// Use function signature instead of return type for variadic functions.
	invokeeType := term.Type()
	if sig := term.Sig(); sig.Variadic {
		invokeeType = sig
	}
	fmt.Fprintf(buf, " %s %s(", invokeeType, term.Invokee.Ident())
	for i, arg := range term.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")
	fmt.Fprintf(buf, "\n\t\tto %s unwind %s", term.NormalRetTarget, term.ExceptionRetTarget)
	return buf.String()
}

// Sig returns the function signature of the invokee.
func (term *TermInvoke) Sig() *FuncType {
	t, ok := term.Invokee.Type().(*PointerType)
	if !ok {
		panic(fmt.Errorf("invalid invokee type; expected *PointerType, got %T", term.Invokee.Type()))
	}
	sig, ok := t.ElemType.(*FuncType)
	if !ok {
		panic(fmt.Errorf("invalid invokee type; expected *FuncType, got %T", t.ElemType))
	}
	return sig
}

// --- [ callbr ] --------------------------------------------------------------

// TermCallBr is an LLVM IR callbr terminator.
type TermCallBr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Callee function.
	// TODO: specify the set of underlying types of Callee.
	Callee Value
	// Function arguments.
	//
	// Arg has one of the following underlying types:
	//    Value
	//    TODO: add metadata value?
	Args []Value
	// Normal control flow return point.
	NormalRetTarget Value // *ir.Block
	// Other control flow return points.
	OtherRetTargets []Value // slice of *ir.Block

	// Type of result produced by the terminator.
	Typ Type
}

// NewCallBr returns a new callbr terminator based on the given callee, function
// arguments and control flow return points for normal and exceptional
// execution.
//
// TODO: specify the set of underlying types of callee.
func NewCallBr(callee Value, args []Value, normalRetTarget *Block, otherRetTargets ...*Block) *TermCallBr {
	// Convert otherRetTargets slice to []Value.
	var otherRets []Value
	for _, otherRetTarget := range otherRetTargets {
		otherRets = append(otherRets, otherRetTarget)
	}
	term := &TermCallBr{Callee: callee, Args: args, NormalRetTarget: normalRetTarget, OtherRetTargets: otherRets}
	// Compute type.
	term.Type()
	return term
}

// String returns the LLVM syntax representation of the terminator as a type-
// value pair.
func (term *TermCallBr) String() string {
	return fmt.Sprintf("%s %s", term.Type(), term.Ident())
}

// Type returns the type of the terminator.
func (term *TermCallBr) Type() Type {
	// Cache type if not present.
	if term.Typ == nil {
		sig := term.Sig()
		term.Typ = sig.RetType
	}
	return term.Typ
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'callbr' CallingConvopt ReturnAttrs=ReturnAttribute* AddrSpaceopt Typ=Type Callee=Value '(' Args ')' FuncAttrs=FuncAttribute* OperandBundles=('[' (OperandBundle separator ',')+ ']')? 'to' NormalRetTarget=Label '[' OtherRetTargets=(Label separator ',')* ']' Metadata=(',' MetadataAttachment)+?
func (term *TermCallBr) LLString() string {
	buf := &strings.Builder{}
	if !term.Type().Equal(Void) {
		fmt.Fprintf(buf, "%s = ", term.Ident())
	}
	buf.WriteString("callbr")
	// Use function signature instead of return type for variadic functions.
	calleeType := term.Type()
	if sig := term.Sig(); sig.Variadic {
		calleeType = sig
	}
	fmt.Fprintf(buf, " %s %s(", calleeType, term.Callee.Ident())
	for i, arg := range term.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")
	fmt.Fprintf(buf, "\n\t\tto %s [", term.NormalRetTarget)
	for i, otherRetTarget := range term.OtherRetTargets {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(otherRetTarget.String())
	}
	buf.WriteString("]")
	return buf.String()
}

// Sig returns the function signature of the callee.
func (term *TermCallBr) Sig() *FuncType {
	t, ok := term.Callee.Type().(*PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *PointerType, got %T", term.Callee.Type()))
	}
	sig, ok := t.ElemType.(*FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee type; expected *FuncType, got %T", t.ElemType))
	}
	return sig
}

// --- [ resume ] --------------------------------------------------------------

// TermResume is an LLVM IR resume terminator.
type TermResume struct {
	// Exception argument to propagate.
	X Value
}

// NewResume returns a new resume terminator based on the given exception
// argument to propagate.
func NewResume(x Value) *TermResume {
	return &TermResume{X: x}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'resume' X=TypeValue Metadata=(',' MetadataAttachment)+?
func (term *TermResume) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "resume %s", term.X)
	return buf.String()
}

// --- [ catchswitch ] ---------------------------------------------------------

// TermCatchSwitch is an LLVM IR catchswitch terminator.
type TermCatchSwitch struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Parent exception pad.
	ParentPad Value // ir.ExceptionPad
	// Exception handlers.
	Handlers []Value // []*ir.Block
	// Optional default target basic block to transfer control flow to; or nil to
	// unwind to caller function.
	DefaultUnwindTarget Value // *ir.Block or nil
}

// NewCatchSwitch returns a new catchswitch terminator based on the given parent
// exception pad, exception handlers and optional default unwind target. If
// defaultUnwindTarget is nil, catchswitch unwinds to caller function.
func NewCatchSwitch(parentPad Value, handlers []*Block, defaultUnwindTarget *Block) *TermCatchSwitch {
	// convert handlers slice to []Value.
	var hs []Value
	for _, handler := range handlers {
		hs = append(hs, handler)
	}
	term := &TermCatchSwitch{ParentPad: parentPad, Handlers: hs}
	if defaultUnwindTarget != nil {
		// Note: since DefaultUnwindTarget is an interface we have to be careful
		// with typed nil values (e.g. `(*ir.Block)(nil)`). This is to ensure that
		// DefaultUnwindTarget is nil and not `{Type: ir.Block, Value: nil}`.
		//
		// ref: https://golang.org/doc/faq#nil_error
		term.DefaultUnwindTarget = defaultUnwindTarget
	}
	return term
}

// String returns the LLVM syntax representation of the terminator as a type-
// value pair.
func (term *TermCatchSwitch) String() string {
	return fmt.Sprintf("%s %s", term.Type(), term.Ident())
}

// Type returns the type of the terminator.
func (term *TermCatchSwitch) Type() Type {
	return Token
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'catchswitch' 'within' ParentPad=ExceptionPad '[' Handlers=Handlers ']' 'unwind' DefaultUnwindTarget=UnwindTarget Metadata=(',' MetadataAttachment)+?
func (term *TermCatchSwitch) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", term.Ident())
	fmt.Fprintf(buf, "catchswitch within %s [", term.ParentPad.Ident())
	for i, handler := range term.Handlers {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(handler.String())
	}
	buf.WriteString("] unwind ")
	if term.DefaultUnwindTarget != nil {
		buf.WriteString(term.DefaultUnwindTarget.String())
	} else {
		buf.WriteString("to caller")
	}
	return buf.String()
}

// --- [ catchret ] ------------------------------------------------------------

// TermCatchRet is an LLVM IR catchret terminator, which catches an in-flight
// exception from CatchPad and returns control flow to normal at Target.
type TermCatchRet struct {
	// Exit catchpad.
	CatchPad Value // *ir.InstCatchPad
	// Target basic block to transfer control flow to.
	Target Value // *ir.Block
}

// NewCatchRet returns a new catchret terminator based on the given exit
// catchpad and target basic block.
func NewCatchRet(catchPad *InstCatchPad, target *Block) *TermCatchRet {
	return &TermCatchRet{CatchPad: catchPad, Target: target}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'catchret' 'from' CatchPad=Value 'to' Target=Label Metadata=(',' MetadataAttachment)+?
func (term *TermCatchRet) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "catchret from %s to %s", term.CatchPad.Ident(), term.Target)
	return buf.String()
}

// --- [ cleanupret ] ----------------------------------------------------------

// TermCleanupRet is an LLVM IR cleanupret terminator, which indicates that the
// personality function of a cleanuppad has finished and transfers control flow
// to an optional target basic block or unwinds to the caller function.
type TermCleanupRet struct {
	// Exit cleanuppad.
	CleanupPad Value // *ir.InstCleanupPad
	// Optional target basic block to transfer control flow to; or nil to unwind
	// to caller function.
	UnwindTarget Value // *ir.Block or nil
}

// NewCleanupRet returns a new cleanupret terminator based on the given exit
// cleanuppad and optional unwind target. If unwindTarget is nil, cleanupret
// unwinds to caller function.
func NewCleanupRet(cleanupPad *InstCleanupPad, unwindTarget *Block) *TermCleanupRet {
	term := &TermCleanupRet{CleanupPad: cleanupPad}
	if unwindTarget != nil {
		// Note: since UnwindTarget is an interface we have to be careful
		// with typed nil values (e.g. `(*ir.Block)(nil)`). This is to ensure that
		// UnwindTarget is nil and not `{Type: ir.Block, Value: nil}`.
		//
		// ref: https://golang.org/doc/faq#nil_error
		term.UnwindTarget = unwindTarget
	}
	return term
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'cleanupret' 'from' CleanupPad=Value 'unwind' UnwindTarget Metadata=(',' MetadataAttachment)+?
func (term *TermCleanupRet) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "cleanupret from %s unwind ", term.CleanupPad.Ident())
	if term.UnwindTarget != nil {
		buf.WriteString(term.UnwindTarget.String())
	} else {
		buf.WriteString("to caller")
	}
	return buf.String()
}

// --- [ unreachable ] ---------------------------------------------------------

// TermUnreachable is an LLVM IR unreachable terminator.
type TermUnreachable struct {
}

// NewUnreachable returns a new unreachable terminator.
func NewUnreachable() *TermUnreachable {
	return &TermUnreachable{}
}

// LLString returns the LLVM syntax representation of the terminator.
//
// 'unreachable' Metadata=(',' MetadataAttachment)+?
func (term *TermUnreachable) LLString() string {
	buf := &strings.Builder{}
	buf.WriteString("unreachable")
	return buf.String()
}

// Terminator instructions
func (*TermRet) isTerminator()         {}
func (*TermBr) isTerminator()          {}
func (*TermCondBr) isTerminator()      {}
func (*TermSwitch) isTerminator()      {}
func (*TermIndirectBr) isTerminator()  {}
func (*TermInvoke) isTerminator()      {}
func (*TermCallBr) isTerminator()      {}
func (*TermResume) isTerminator()      {}
func (*TermCatchSwitch) isTerminator() {}
func (*TermCatchRet) isTerminator()    {}
func (*TermCleanupRet) isTerminator()  {}
func (*TermUnreachable) isTerminator() {}
