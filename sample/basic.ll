@global.a = global i32 zeroinitializer
@global.b = global i32 1
@global.c = global i32 2
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"

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
	br i1 %4, label %6, label %9

exit:
	%5 = load i32, i32* %0
	ret i32 %5

6:
	%7 = bitcast [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*
	%8 = call i32 @puts(i8* %7)
	br label %9

9:
	%10 = load i32, i32* @global.a
	%11 = icmp eq i32 %10, 1
	br i1 %11, label %12, label %15

12:
	%13 = bitcast [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*
	%14 = call i32 @puts(i8* %13)
	br label %15

15:
	store i32 0, i32* %0
	br label %exit
}

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)
