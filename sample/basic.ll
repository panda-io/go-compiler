declare i32 @printf(i8* %0)

declare i8* @malloc(i32 %0)

declare i8* @calloc(i32 %0, i32 %1)

declare i8* @realloc(i8* %0, i32 %1)

declare void @free(i8* %0)

define i32 @add(i32 %0, i32 %1) {
2:
	%3 = add i32 %0, %1
	ret i32 %3
}

define i32 @main() {
0:
	%1 = call i32 @add(i32 1, i32 2)
	ret i32 %1
}
