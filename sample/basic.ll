@global.color.red = global i32 0
@global.color.green = global i32 1
@global.color.blue = global i32 2

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
