@global.color.red = global i32 0
@global.color.green = global i32 1
@global.color.blue = global i32 2
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define i32 @global.add(i32 %a, i32 %b) {
entry:
	%0 = alloca i32
	store i32 %a, i32* %0
	%1 = alloca i32
	store i32 %b, i32* %1
	%2 = load i32, i32* %0
	%3 = load i32, i32* %1
	%4 = add i32 %2, %3
	ret i32 %4
}

define void @global.nothing(i32 %a, i32 %b) {
entry:
	%0 = alloca i32
	store i32 %a, i32* %0
	%1 = alloca i32
	store i32 %b, i32* %1
	ret void
}

define i32 @main() {
entry:
	%0 = getelementptr [12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i32 0, i32 0
	%1 = call i32 @puts(i8* %0)
	%2 = call i32 @global.add(i32 1, i32 2)
	ret i32 %2
}
