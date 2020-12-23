define i32 @main() {
0:
	%1 = call i32 @add(i32 1, i32 2)
	ret i32 %1
}

define i32 @add(i32 %0, i32 %1) {
2:
	%3 = add i32 %0, %1
	ret i32 %3
}