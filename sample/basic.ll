@global.a = global i32 zeroinitializer
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"
@string.3860584d72c3c1e2f6ebe3a6e165a60f = constant [8 x i8] c"default\00"

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
	%4 = add i32 %3, 1
	store i32 %4, i32* @global.a
	%5 = load i32, i32* @global.a
	%6 = icmp eq i32 %5, 0
	br i1 %6, label %11, label %8


exit:
	%7 = load i32, i32* %0
	ret i32 %7


8:
	%9 = load i32, i32* @global.a
	%10 = icmp eq i32 %9, 1
	br i1 %10, label %16, label %14


11:
	%12 = bitcast [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*
	%13 = call i32 @puts(i8* %12)
	br label %8


14:
	%15 = load i32, i32* @global.a
	switch i32 %15, label %20 [
		i32 0, label %23
		i32 1, label %26
	]


16:
	%17 = bitcast [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*
	%18 = call i32 @puts(i8* %17)
	br label %14


19:
	store i32 0, i32* %0
	br label %exit


20:
	%21 = bitcast [8 x i8]* @string.3860584d72c3c1e2f6ebe3a6e165a60f to i8*
	%22 = call i32 @puts(i8* %21)
	br label %19


23:
	%24 = bitcast [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*
	%25 = call i32 @puts(i8* %24)
	br label %19


26:
	%27 = bitcast [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*
	%28 = call i32 @puts(i8* %27)
	br label %19

}
