@global.a = global i32 zeroinitializer
@global.b = global i32 1
@global.c = global i32 2

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
	store i32 0, i32* %0
	br label %exit

exit:
	%1 = load i32, i32* %0
	ret i32 %1
}
