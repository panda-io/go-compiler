@global.a = global i32 zeroinitializer
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"
@string.3860584d72c3c1e2f6ebe3a6e165a60f = constant [8 x i8] c"default\00"
@string.43d2b345b462cd06e5f7841d678588a5 = constant [5 x i8] c"loop\00"

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
	br i1 %8, label %14, label %12


9:
	%10 = bitcast [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*
	%11 = call i32 @puts(i8* %10)
	br label %6


12:
	%13 = load i32, i32* @global.a
	switch i32 %13, label %18 [
		i32 0, label %21
		i32 1, label %24
	]


14:
	%15 = bitcast [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*
	%16 = call i32 @puts(i8* %15)
	br label %12


17:
	br label %28


18:
	%19 = bitcast [8 x i8]* @string.3860584d72c3c1e2f6ebe3a6e165a60f to i8*
	%20 = call i32 @puts(i8* %19)
	br label %17


21:
	%22 = bitcast [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268 to i8*
	%23 = call i32 @puts(i8* %22)
	br label %17


24:
	%25 = bitcast [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc to i8*
	%26 = call i32 @puts(i8* %25)
	br label %17


27:
	store i32 0, i32* %0
	br label %exit


28:
	%29 = load i32, i32* @global.a
	%30 = icmp slt i32 %29, 5
	br i1 %30, label %34, label %27


31:
	%32 = load i32, i32* @global.a
	%33 = add i32 %32, 1
	store i32 %33, i32* @global.a
	br label %28


34:
	%35 = bitcast [5 x i8]* @string.43d2b345b462cd06e5f7841d678588a5 to i8*
	%36 = call i32 @puts(i8* %35)
	br label %31

}

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)
