# Panda compiler written in golang

> ## **ast**
> ---
> ### declarations
> - variable (var, const)
> - function
> - enum
> - interface
> - class
> ### statements
> - emit
> - declaration
> - empty (;)
> - increase-decrease (++, --)
> - expression
> - if-else
> - for (foreach)
> - switch
> - return
> - branch (break, continue)
> - try-catch-finally
> - block
> - yield
> - throw
> - await
> ### expressions
> - literal
> - identifier
> - unary
> - binary
> - ternary
> - new
> - this-base
> - invocation (call)
> - selector (.)
> - element-access ([])
> - paren
> - ellipsis
> ### type
> - field
> - field-list
> ### metadata
> - text
> - object (literal)
> ### modifiers
> - public
> - static
> - async
















--------------------- 备忘 -------------------------
unresolved can be cached
file has import scope
    
--------------------- 限制 TO-DO -------------------
#前置声明，声明，定义

#serialize metadata

#变量声明时初始化，只能是basic lit, 复杂类型不能在声明时初始化

#interface 不能继承
#class 单继承

#声明，如果没有设置初始值，则自动赋值

#静态变量赋值

#store file info into program