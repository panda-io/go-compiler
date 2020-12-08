# Panda compiler

## **AST**

### **declarations**
- variable (var, const)
- function
- enum
- interface
- class

### **statements**
- raw source @"source"
- empty statement ;
- expression statement
- declaration statement
- compound statements {}
- if statement 
  - if (condition) {} else
  - if (initialize;condition) {} else
- switch statement 
  - switch(init)
  - switch(init;operand) {}
- for statement 
  - for {} 
  - for (condition) {}
  - for (init;condition)
  - for (init;condition;loop) {}
  - for (init value : range) {}
  - for (init key; init value : range) {}
- continue statement
- break statement
- return statement
- try statement
- throw statement

### **expressions**
- primary expression
  - literal
  - this, base
  - name
  - .name //qualified name
  
- postfix operators
  - subscript operator	[]
  - function call	()
  - member access operator	.
  - postfix increment operator	++
  - postfix decrement operator	--

- unary operators
  - unary plus operator +
  - unary negation operator -
  - logical negation operator !
  - complement operator ~
  - new operator

- binary operators
  - multiplicative operators *, /, %
  - additive operators +, -
  - shift operators >>, <<
  - relational and equality operators <, >, <=, >=, ==, !=
  - bitwise operators &, ^, |
  - logical operators &&, ||
  - assignment operators =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=

- conditional operator condition ? expression : expression 

- constant expressions
  - constant expressions used in declarations of:
    - array bounds
    - selectors in case statements
    - bit-field length specification
    - enumeration initializers
  - only operands that are legal constant expressions:
    - literals
    - enumeration constants
    - values declared as const

- parentheses experssions ()

- #operator precedence and associativity#
  - group 1. scope resolution (type)	a.b.c
  - group 2. left to right:
    - member selection .
    - subscript  []
    - function call	()
    - postfix increment	++	
    - postfix decrement	--	
    - type name	typeid	
  - group 3. right to left
    - complement	~
    - logical not	!
    - unary negation	-	
    - unary plus	+
    - create object	new
  - group 4. left to right
    - multiplication	*	
    - division	/	
    - modulus	%	
  - group 5. left to right
    - addition	+	
    - subtraction	-	
  - group 6. left to right
    - left shift	<<	
    - right shift	>>	
  - group 7. left to right
    - less than	<	
    - greater than	>	
    - less than or equal to	<=	
    - greater than or equal to	>=	
  - group 8. left to right
    - equality	==	
    - inequality	!=	
  - group 9. left to right
    - bitwise AND	&
  - group 10. left to right
    - bitwise exclusive OR	^
  - group 11. left to right
    - bitwise inclusive OR	|
  - group 12. left to right
    - logical AND	&&
  - group 13. left to right
    - logical OR	||
  - group 14. right to left
    - conditional	? :	
  - group 15. right to left
    - assignment	=	
    - multiplication assignment	*=	
    - division assignment	/=	
    - modulus assignment	%=	
    - addition assignment	+=	
    - subtraction assignment	-=	
    - left-shift assignment	<<=	
    - right-shift assignment	>>=	
    - bitwise AND assignment	&=
    - bitwise inclusive OR assignment	|=	
    - bitwise exclusive OR assignment	^=
 
### **type**
 - builtin type
 - type name
 - type argument
 - type argument list
 - type parameter
 - type parameter list
 - argument
 - argument list
 - parameter
 - parameter list

### **metadata**
- text          
- object (literal)
  
### **modifiers**
- public
- static

### **operator overloading**
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

### **limitations**
- no function overloading
- single inheritance
- variable declaration can only use const expression for initialization

### **roadmap**
- reflection
- serlize
  - binary
-------------------------
- coroutine
- new object(){a = 1, b = 2, c = 3}
