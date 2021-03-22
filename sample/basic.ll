@string.5bddf146c13b387514280200e83cf08b = constant [15 x i8] c"value #1: %d \0A\00"
@string.aee0f3d368512408f6bd5274bf51a219 = constant [15 x i8] c"value #2: %d \0A\00"

declare i32 @puts(i8* %text)

declare i32 @printf(i8* %format, ...)

declare i8* @malloc(i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define void @global.print_number(i32* %value) {
entry:
	br label %body


body:
	%0 = load i32, i32* %value
	%1 = call i32 (i8*, ...) @printf(i8* bitcast ([15 x i8]* @string.5bddf146c13b387514280200e83cf08b to i8*), i32 %0)
	%2 = load i32, i32* %value
	%3 = call i32 (i8*, ...) @printf(i8* bitcast ([15 x i8]* @string.aee0f3d368512408f6bd5274bf51a219 to i8*), i32 %2)
	br label %exit


exit:
	ret void

}
