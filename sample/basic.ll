%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*, i1)* }
%global.counter.vtable.type = type { %global.counter* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.base = type { %global.base.vtable.type* }
%global.base.vtable.type = type { %global.base* ()*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.base.vtable.data = global %global.base.vtable.type { %global.base* ()* @global.base.create, void (i8*)* @global.base.destroy }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.319f93d592b03bf2f832e7b2592ace2e = constant [17 x i8] c"base contruction\00"
@string.362aeeddb3d01da539cb6755bde46953 = constant [17 x i8] c"base destruction\00"

declare i32 @puts(i8* %text)

declare i32 @printf(i8* %format, ...)

declare i8* @malloc(i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

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
	%6 = call i32 @puts(i8* bitcast ([15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e to i8*))
	br label %exit


exit:
	%7 = load %global.counter*, %global.counter** %0
	ret %global.counter* %7

}

define void @global.counter.destroy(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = call i32 @puts(i8* bitcast ([16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f to i8*))
	br label %exit


exit:
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
	%1 = icmp eq %global.counter* %0, null
	br i1 %1, label %9, label %2


exit:
	ret void


2:
	%3 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%4 = load i32, i32* %3
	%5 = sub i32 %4, 1
	store i32 %5, i32* %3
	%6 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%7 = load i32, i32* %6
	%8 = icmp eq i32 %7, 0
	br i1 %8, label %20, label %10


9:
	br label %exit


10:
	br label %exit


11:
	%12 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 4
	%13 = load void (i8*, i1)*, void (i8*, i1)** %12
	%14 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%15 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%16 = load i8*, i8** %15
	call void %13(i8* %16, i1 true)
	%17 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%18 = load i32, i32* %17
	%19 = icmp eq i32 %18, 0
	br i1 %19, label %21, label %20


20:
	br label %10


21:
	%22 = bitcast %global.counter* %0 to i8*
	%23 = bitcast %global.counter* %0 to i8*
	call void @free(i8* %23)
	br label %20

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
	%1 = icmp eq %global.counter* %0, null
	br i1 %1, label %13, label %2


exit:
	ret void


2:
	%3 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%4 = load i32, i32* %3
	%5 = sub i32 %4, 1
	store i32 %5, i32* %3
	%6 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%7 = load i32, i32* %6
	%8 = icmp eq i32 %7, 0
	%9 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%10 = load i32, i32* %9
	%11 = icmp eq i32 %10, 0
	%12 = and i1 %8, %11
	br i1 %12, label %15, label %14


13:
	br label %exit


14:
	br label %exit


15:
	%16 = bitcast %global.counter* %0 to i8*
	%17 = bitcast %global.counter* %0 to i8*
	call void @free(i8* %17)
	br label %14

}

define i32 @main() {
entry:
	%0 = alloca i32
	br label %body


body:
	%1 = call %global.counter* @global.counter.create()
	store i32 0, i32* %0
	br label %exit


exit:
	%2 = bitcast %global.counter* %1 to i8*
	call void @global.counter.destroy(i8* %2)
	%3 = load i32, i32* %0
	ret i32 %3

}

define %global.base* @global.base.create() {
entry:
	%0 = alloca %global.base*
	%1 = getelementptr %global.base, %global.base* null, i32 1
	%2 = ptrtoint %global.base* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	call void @memset(i8* %3, i32 0, i32 %2)
	%4 = bitcast i8* %3 to %global.base*
	%5 = getelementptr %global.base, %global.base* %4, i32 0, i32 0
	store %global.base.vtable.type* @global.base.vtable.data, %global.base.vtable.type** %5
	store %global.base* %4, %global.base** %0
	br label %body


body:
	%6 = call i32 @puts(i8* bitcast ([17 x i8]* @string.319f93d592b03bf2f832e7b2592ace2e to i8*))
	br label %exit


exit:
	%7 = load %global.base*, %global.base** %0
	ret %global.base* %7

}

define void @global.base.destroy(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.base*
	br label %body


body:
	%1 = call i32 @puts(i8* bitcast ([17 x i8]* @string.362aeeddb3d01da539cb6755bde46953 to i8*))
	br label %exit


exit:
	ret void

}
