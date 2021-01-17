@global.color.red = global i32 0
@global.color.green = global i32 1
@global.color.blue = global i32 2

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define i32 @global.add(i32 %a, i32 %b) {
entry:
	%0 = load i32, i32 %a
	%1 = load i32, i32 %b
	%2 = add i32 %0, %1
	ret i32 %2
}

define i32 @main() {
entry:
	%0 = call i32 @global.add(i32 1, i32 2)
	ret i32 %0
}
