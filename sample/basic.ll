@global.a = global i32 zeroinitializer
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"
@string.3860584d72c3c1e2f6ebe3a6e165a60f = constant [8 x i8] c"default\00"
@string.43d2b345b462cd06e5f7841d678588a5 = constant [5 x i8] c"loop\00"
@string.e4774a52ffe13545ae9d5ce20dd839b9 = constant [9 x i8] c"a = %d \0A\00"

declare i32 @puts(i8* %text)

declare i32 @printf(i8* %format, ...)

declare i8* @malloc(i32 %size)

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
	br i1 %8, label %13, label %11


9:
	%10 = call i32 @puts(i8* bitcast ([7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*))
	br label %6


11:
	%12 = load i32, i32* @global.a
	switch i32 %12, label %16 [
		i32 0, label %18
		i32 1, label %20
	]


13:
	%14 = call i32 @puts(i8* bitcast ([7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*))
	br label %11


15:
	br label %25


16:
	%17 = call i32 @puts(i8* bitcast ([8 x i8]* @string.3860584d72c3c1e2f6ebe3a6e165a60f to i8*))
	br label %15


18:
	%19 = call i32 @puts(i8* bitcast ([7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*))
	br label %15


20:
	%21 = call i32 @puts(i8* bitcast ([7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*))
	br label %15


22:
	%23 = load i32, i32* @global.a
	%24 = call i32 (i8*, ...) @printf(i8* bitcast ([9 x i8]* @string.e4774a52ffe13545ae9d5ce20dd839b9 to i8*), i32 %23)
	store i32 0, i32* %0
	br label %exit


25:
	%26 = load i32, i32* @global.a
	%27 = icmp slt i32 %26, 5
	br i1 %27, label %31, label %22


28:
	%29 = load i32, i32* @global.a
	%30 = add i32 %29, 1
	store i32 %30, i32* @global.a
	br label %25


31:
	%32 = call i32 @puts(i8* bitcast ([5 x i8]* @string.43d2b345b462cd06e5f7841d678588a5 to i8*))
	br label %28

}
