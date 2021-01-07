@global.color.red = global i32 0
@global.color.green = global i32 1
@global.color.blue = global i32 2

declare i32 @printf(i8* %0)

declare i8* @malloc(i32 %0)

declare i8* @realloc(i8* %0, i32 %1)

declare void @free(i8* %0)

declare i32 @memcmp(i8* %0, i8* %1, i32 %2)

declare void @memcpy(i8* %0, i8* %1, i32 %2)

declare void @memset(i8* %0, i32 %1, i32 %2)

define i32 @global.add(i32 %0, i32 %1) {
2:
	%3 = add i32 %0, %1
	ret i32 %3
}

define i32 @main() {
0:
	%1 = call i32 @global.add(i32 1, i32 2)
	ret i32 %1
}
