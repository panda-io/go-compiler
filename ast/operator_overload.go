package ast

/************************************

- unary operations

    | expression	| translated to |
    |:---:|:---|
    |+a	| a.unary_plus() |
    |-a	| a.unary_minus() |
    |!a	| a.unary_not() |

- increments and decrements

    | expression	| translated to |
    |:---:|:---|
    |a++| a.increment() |
    |a--| a.decrement() |

- binary operations

    | expression	| translated to |
    |:---:|:---|
    |a + b| a.plus(b) |
    |a - b| a.minus(b) |
    |a * b| a.mul(b) |
    |a / b| a.div(b) |
    |a % b| a.mod(b) |

- indexed access

    | expression	| translated to |
    |:---:|:---|
    |a[i]| a.get(i) |
    |a[i] = value| a.set(i, value) |

- invoke operator

    | expression	| translated to |
    |:---:|:---|
    |a()| a.invoke() |
    |a(b, c, d)| a.invoke(b, c, d) |

- augmented assignments

    | expression	| translated to |
    |:---:|:---|
    |a = b| a.assign(b) |
    |a += b| a.plus_assign(b) |
    |a -= b| a.minus_assign(b) |
    |a *= b| a.mul_assign(b) |
    |a /= b| a.div_assign(b) |
    |a %= b| a.mod_assign(b) |

- equality, inequality and comparison

    | expression	| translated to |
    |:---:|:---|
    |a == b| a.compare(b) == 0 |
    |a != b| a.compare(b) != 0 |
    |a > b| a.compare(b) > 0 |
    |a < b| a.compare(b) < 0 |
    |a >= b| a.compare(b) >= 0 |
	|a <= b| a.compare(b) <= 0 |

***************************************/

const (
	OperatorUnaryPlus  = "unary_plus"
	OperatorUnaryMinus = "unary_minus"
	OperatorUnaryNot   = "unary_not"

	OperatorIncreament = "increment"
	OperatorDecreament = "decrement"

	OperatorBinaryPlus  = "plus"
	OperatorBinaryMinus = "minus"
	OperatorBinaryMul   = "mul"
	OperatorBinaryDiv   = "div"
	OperatorBinaryMod   = "mod"

	OperatorGet = "get"
	OperatorSet = "set"

	OperatorInvoke = "invoke"

	OperatorAssign      = "assign"
	OperatorPlusAssign  = "plus_assign"
	OperatorMinusAssign = "minus_assign"
	OperatorMulAssign   = "mul_assign"
	OperatorDivAssign   = "div_assign"
	OperatorRemAssign   = "rem_assign"

	OperatorCompare = "compare"
)
