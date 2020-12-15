package ir

// ClauseType specifies the clause type of a landingpad clause.
type ClauseType uint8

// Clause types.
const (
	ClauseTypeCatch  ClauseType = iota + 1 // catch
	ClauseTypeFilter                       // filter
)

// AtomicOrdering is an atomic ordering attribute.
type AtomicOrdering uint8

// AtomicOp is an AtomicRMW binary operation.
type AtomicOp uint8

// AtomicRMW binary operations.
const (
	AtomicOpAdd  AtomicOp = iota + 1 // add
	AtomicOpAnd                      // and
	AtomicOpFAdd                     // fadd
	AtomicOpFSub                     // fsub
	AtomicOpMax                      // max
	AtomicOpMin                      // min
	AtomicOpNAnd                     // nand
	AtomicOpOr                       // or
	AtomicOpSub                      // sub
	AtomicOpUMax                     // umax
	AtomicOpUMin                     // umin
	AtomicOpXChg                     // xchg
	AtomicOpXor                      // xor
)

// Atomic ordering attributes.
const (
	AtomicOrderingNone      AtomicOrdering = iota // none
	AtomicOrderingAcqRel                          // acq_rel
	AtomicOrderingAcquire                         // acquire
	AtomicOrderingMonotonic                       // monotonic
	AtomicOrderingRelease                         // release
	AtomicOrderingSeqCst                          // seq_cst
	AtomicOrderingUnordered                       // unordered
)

// OverflowFlag is an integer overflow flag.
type OverflowFlag uint8

// Overflow flags.
const (
	OverflowFlagNSW OverflowFlag = iota // nsw
	OverflowFlagNUW                     // nuw
)

// IPred is an integer comparison predicate.
type IPred uint8

// Integer predicates.
const (
	IPredEQ  IPred = iota // eq
	IPredNE               // ne
	IPredSGE              // sge
	IPredSGT              // sgt
	IPredSLE              // sle
	IPredSLT              // slt
	IPredUGE              // uge
	IPredUGT              // ugt
	IPredULE              // ule
	IPredULT              // ult
)

// FPred is a floating-point comparison predicate.
type FPred uint8

// Floating-point predicates.
const (
	FPredFalse FPred = iota // false
	FPredOEQ                // oeq
	FPredOGE                // oge
	FPredOGT                // ogt
	FPredOLE                // ole
	FPredOLT                // olt
	FPredONE                // one
	FPredORD                // ord
	FPredTrue               // true
	FPredUEQ                // ueq
	FPredUGE                // uge
	FPredUGT                // ugt
	FPredULE                // ule
	FPredULT                // ult
	FPredUNE                // une
	FPredUNO                // uno
)
