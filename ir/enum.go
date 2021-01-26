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
type IPred string

// Integer predicates.
const (
	IPredEQ  IPred = "eq"  // eq
	IPredNE  IPred = "ne"  // ne
	IPredSGE IPred = "sge" // sge
	IPredSGT IPred = "sgt" // sgt
	IPredSLE IPred = "sle" // sle
	IPredSLT IPred = "slt" // slt
	IPredUGE IPred = "uge" // uge
	IPredUGT IPred = "ugt" // ugt
	IPredULE IPred = "ule" // ule
	IPredULT IPred = "ult" // ult
)

// FPred is a floating-point comparison predicate.
type FPred string

// Floating-point predicates.
const (
	FPredFalse FPred = "false" // false
	FPredOEQ   FPred = "oeq"   // oeq
	FPredOGE   FPred = "oge"   // oge
	FPredOGT   FPred = "ogt"   // ogt
	FPredOLE   FPred = "ole"   // ole
	FPredOLT   FPred = "olt"   // olt
	FPredONE   FPred = "one"   // one
	FPredORD   FPred = "ord"   // ord
	FPredTrue  FPred = "true"  // true
	FPredUEQ   FPred = "ueq"   // ueq
	FPredUGE   FPred = "uge"   // uge
	FPredUGT   FPred = "ugt"   // ugt
	FPredULE   FPred = "ule"   // ule
	FPredULT   FPred = "ult"   // ult
	FPredUNE   FPred = "une"   // une
	FPredUNO   FPred = "uno"   // uno
)
