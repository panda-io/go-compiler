%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*, i1)* }
%global.counter.vtable.type = type { %global.counter* ()*, void (i8*, i1)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (i8*, i1)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.a = global i32 zeroinitializer
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

define %global.counter* @global.counter.create() {
entry:
	%0 = alloca %global.counter*
	%1 = getelementptr %global.counter, %global.counter* null, i32 1
	%2 = ptrtoint %global.counter* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	call void @memset(i8* %3, i32 0, i32 %2)
	%4 = bitcast i8* %3 to %global.counter*
	%5 = getelementptr %global.counter, %global.counter* %4, i32 0, i32 0
	store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %5
	store %global.counter* %4, %global.counter** %0
	br label %body


body:
	br label %exit


exit:
	%6 = load %global.counter*, %global.counter** %0
	ret %global.counter* %6

}

define void @global.counter.destroy(i8* %this, i1 %free) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	%1 = alloca i1
	store i1 %free, i1* %1
	br label %body


body:
	br label %exit


exit:
	%2 = bitcast i8* %this to i8*
	call void @free(i8* %2)
	ret void

}

define void @global.counter.retain_shared(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%2 = load i32, i32* %1
	%3 = add i32 %2, 1
	store i32 %3, i32* %1
	br label %exit


exit:
	ret void

}

define void @global.counter.release_shared(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%2 = load i32, i32* %1
	%3 = sub i32 %2, 1
	store i32 %3, i32* %1
	%4 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%5 = load i32, i32* %4
	%6 = icmp eq i32 %5, 0
	br i1 %6, label %17, label %7


exit:
	ret void


7:
	br label %exit


8:
	%9 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 4
	%10 = load void (i8*, i1)*, void (i8*, i1)** %9
	%11 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%12 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%13 = load i8*, i8** %12
	call void %10(i8* %13, i1 true)
	%14 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%15 = load i32, i32* %14
	%16 = icmp eq i32 %15, 0
	br i1 %16, label %18, label %17


17:
	br label %7


18:
	%19 = bitcast %global.counter* %0 to i8*
	%20 = bitcast %global.counter* %0 to i8*
	call void @free(i8* %20)
	br label %17

}

define void @global.counter.retain_weak(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%2 = load i32, i32* %1
	%3 = add i32 %2, 1
	store i32 %3, i32* %1
	br label %exit


exit:
	ret void

}

define void @global.counter.release_weak(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%2 = load i32, i32* %1
	%3 = sub i32 %2, 1
	store i32 %3, i32* %1
	%4 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%5 = load i32, i32* %4
	%6 = icmp eq i32 %5, 0
	%7 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%8 = load i32, i32* %7
	%9 = icmp eq i32 %8, 0
	%10 = and i1 %6, %9
	br i1 %10, label %12, label %11


exit:
	ret void


11:
	br label %exit


12:
	%13 = bitcast %global.counter* %0 to i8*
	%14 = bitcast %global.counter* %0 to i8*
	call void @free(i8* %14)
	br label %11

}

define i32 @main() {
entry:
	%0 = alloca i32
	br label %body


body:
	%1 = call i32 @puts(i8* bitcast ([12 x i8]* @string.cb091131e20d7842e7627e8736856b45 to i8*))
	store i32 0, i32* %0
	br label %exit


exit:
	%2 = load i32, i32* %0
	ret i32 %2

}

declare i32 @puts(i8* %text)

declare i32 @printf(i8* %format, ...)

declare i8* @malloc(i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)
