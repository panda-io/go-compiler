@global.a = global i32 zeroinitializer
@string.7ac4a7d2c8bf695dc8a116f6b70a8e1a = constant [9 x i8] c"b = %d \0A\00"
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

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
	%1 = alloca i32
	br label %body


body:
	store i32 9, i32* %1
	%2 = load i32, i32* %1
	%3 = call i32 (i8*, ...) @printf(i8* bitcast ([9 x i8]* @string.7ac4a7d2c8bf695dc8a116f6b70a8e1a to i8*), i32 %2)
	%4 = call i32 @puts(i8* bitcast ([12 x i8]* @string.cb091131e20d7842e7627e8736856b45 to i8*))
	store i32 0, i32* %0
	br label %exit


exit:
	%5 = load i32, i32* %0
	ret i32 %5

}
