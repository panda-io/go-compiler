@global.a = global i32 zeroinitializer
@global.b = global i32 1
@global.c = global i32 2
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define i32 @main() {
entry:
	%0 = alloca i32
	br label %body

body:
	%1 = load i32, i32* @global.a
	%2 = add i32 %1, 1
	store i32 %2, i32* @global.a
	%3 = load i32, i32* @global.a
	%4 = icmp eq i32 %3, 0
	br i1 %4, label %9, label %6

exit:
	%5 = load i32, i32* %0
	ret i32 %5

6:
	%7 = load i32, i32* @global.a
	%8 = icmp eq i32 %7, 1
	br i1 %8, label %13, label %12

9:
	%10 = bitcast [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*
	%11 = call i32 @puts(i8* %10)
	br label %6

12:
	store i32 0, i32* %0
	br label %exit

13:
	%14 = bitcast [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*
	%15 = call i32 @puts(i8* %14)
	br label %12
}
