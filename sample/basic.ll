%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }
%global.counter.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { i8* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.c1432ab71496ebb8b3e30bbcf37605e7 = constant [20 x i8] c"retain shared: %d \0A\00"
@string.ce99a7174da84f1767b9d5235ea3c24f = constant [21 x i8] c"release shared: %d \0A\00"
@string.21c67ac9191c65481dbab306227b4840 = constant [17 x i8] c"free object %p \0A\00"
@string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 = constant [18 x i8] c"free counter %p \0A\00"
@string.8d9c52192bdfa908703a004b070ff63e = constant [18 x i8] c"retain weak: %d \0A\00"
@string.ad2ae91c8542c824c6842efed6523f49 = constant [19 x i8] c"release weak: %d \0A\00"
@string.e4b993b5f16d57ebba5166037b305638 = constant [15 x i8] c"value #3: %d \0A\00"
@string.5bddf146c13b387514280200e83cf08b = constant [15 x i8] c"value #1: %d \0A\00"
@string.aee0f3d368512408f6bd5274bf51a219 = constant [15 x i8] c"value #2: %d \0A\00"

define i8* @global.counter.create() {
entry:
	%0 = alloca i8*
	%1 = getelementptr %global.counter, %global.counter* null, i32 1
	%2 = ptrtoint %global.counter* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	call void @memset(i8* %3, i32 0, i32 %2)
	%4 = bitcast i8* %3 to %global.counter*
	%5 = getelementptr %global.counter, %global.counter* %4, i32 0, i32 0
	store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %5
	store i8* %3, i8** %0
	br label %body


body:
	%6 = call i32 @puts(i8* bitcast ([15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e to i8*))
	br label %exit


exit:
	%7 = load i8*, i8** %0
	ret i8* %7

}

define void @global.counter.destroy(i8* %this) {
entry:
	br label %body


body:
	call void @free(i8* %this)
	%0 = call i32 @puts(i8* bitcast ([16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f to i8*))
	br label %exit


exit:
	ret void

}

define void @global.counter.retain_shared(i8* %this) {
entry:
	br label %body


body:
	%0 = bitcast i8* %this to %global.counter*
	%1 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%2 = load i32, i32* %1
	%3 = add i32 %2, 1
	store i32 %3, i32* %1
	%4 = bitcast i8* %this to %global.counter*
	%5 = getelementptr %global.counter, %global.counter* %4, i32 0, i32 1
	%6 = load i32, i32* %5
	%7 = call i32 (i8*, ...) @printf(i8* bitcast ([20 x i8]* @string.c1432ab71496ebb8b3e30bbcf37605e7 to i8*), i32 %6)
	br label %exit


exit:
	ret void

}

define void @global.counter.release_shared(i8* %this) {
entry:
	br label %body


body:
	%0 = icmp eq i8* %this, null
	br i1 %0, label %14, label %1


exit:
	ret void


1:
	%2 = bitcast i8* %this to %global.counter*
	%3 = getelementptr %global.counter, %global.counter* %2, i32 0, i32 1
	%4 = load i32, i32* %3
	%5 = sub i32 %4, 1
	store i32 %5, i32* %3
	%6 = bitcast i8* %this to %global.counter*
	%7 = getelementptr %global.counter, %global.counter* %6, i32 0, i32 1
	%8 = load i32, i32* %7
	%9 = call i32 (i8*, ...) @printf(i8* bitcast ([21 x i8]* @string.ce99a7174da84f1767b9d5235ea3c24f to i8*), i32 %8)
	%10 = bitcast i8* %this to %global.counter*
	%11 = getelementptr %global.counter, %global.counter* %10, i32 0, i32 1
	%12 = load i32, i32* %11
	%13 = icmp eq i32 %12, 0
	br i1 %13, label %16, label %15


14:
	br label %exit


15:
	br label %exit


16:
	%17 = bitcast i8* %this to %global.counter*
	%18 = getelementptr %global.counter, %global.counter* %17, i32 0, i32 4
	%19 = load void (i8*)*, void (i8*)** %18
	%20 = bitcast i8* %this to %global.counter*
	%21 = getelementptr %global.counter, %global.counter* %20, i32 0, i32 3
	%22 = load i8*, i8** %21
	call void %19(i8* %22)
	%23 = bitcast i8* %this to %global.counter*
	%24 = getelementptr %global.counter, %global.counter* %23, i32 0, i32 3
	%25 = load i8*, i8** %24
	%26 = call i32 (i8*, ...) @printf(i8* bitcast ([17 x i8]* @string.21c67ac9191c65481dbab306227b4840 to i8*), i8* %25)
	%27 = bitcast i8* %this to %global.counter*
	%28 = getelementptr %global.counter, %global.counter* %27, i32 0, i32 3
	%29 = load i8*, i8** %28
	call void @free(i8* %29)
	%30 = bitcast i8* %this to %global.counter*
	%31 = getelementptr %global.counter, %global.counter* %30, i32 0, i32 3
	%32 = load i8*, i8** %31
	%33 = bitcast i8* %this to %global.counter*
	%34 = getelementptr %global.counter, %global.counter* %33, i32 0, i32 3
	store i8* null, i8** %34
	%35 = bitcast i8* %this to %global.counter*
	%36 = getelementptr %global.counter, %global.counter* %35, i32 0, i32 2
	%37 = load i32, i32* %36
	%38 = icmp eq i32 %37, 0
	br i1 %38, label %40, label %39


39:
	br label %15


40:
	%41 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 to i8*), i8* %this)
	call void @free(i8* %this)
	br label %39

}

define void @global.counter.retain_weak(i8* %this) {
entry:
	br label %body


body:
	%0 = bitcast i8* %this to %global.counter*
	%1 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%2 = load i32, i32* %1
	%3 = add i32 %2, 1
	store i32 %3, i32* %1
	%4 = bitcast i8* %this to %global.counter*
	%5 = getelementptr %global.counter, %global.counter* %4, i32 0, i32 2
	%6 = load i32, i32* %5
	%7 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.8d9c52192bdfa908703a004b070ff63e to i8*), i32 %6)
	br label %exit


exit:
	ret void

}

define void @global.counter.release_weak(i8* %this) {
entry:
	br label %body


body:
	%0 = icmp eq i8* %this, null
	br i1 %0, label %19, label %1


exit:
	ret void


1:
	%2 = bitcast i8* %this to %global.counter*
	%3 = getelementptr %global.counter, %global.counter* %2, i32 0, i32 2
	%4 = load i32, i32* %3
	%5 = sub i32 %4, 1
	store i32 %5, i32* %3
	%6 = bitcast i8* %this to %global.counter*
	%7 = getelementptr %global.counter, %global.counter* %6, i32 0, i32 2
	%8 = load i32, i32* %7
	%9 = call i32 (i8*, ...) @printf(i8* bitcast ([19 x i8]* @string.ad2ae91c8542c824c6842efed6523f49 to i8*), i32 %8)
	%10 = bitcast i8* %this to %global.counter*
	%11 = getelementptr %global.counter, %global.counter* %10, i32 0, i32 1
	%12 = load i32, i32* %11
	%13 = icmp eq i32 %12, 0
	%14 = bitcast i8* %this to %global.counter*
	%15 = getelementptr %global.counter, %global.counter* %14, i32 0, i32 2
	%16 = load i32, i32* %15
	%17 = icmp eq i32 %16, 0
	%18 = and i1 %13, %17
	br i1 %18, label %21, label %20


19:
	br label %exit


20:
	br label %exit


21:
	%22 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 to i8*), i8* %this)
	call void @free(i8* %this)
	br label %20

}

declare i32 @puts(i8* %text)

declare i32 @printf(i8* %format, ...)

declare i8* @malloc(i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = alloca i32
	br label %body


body:
	store i32 1, i32* %1
	%2 = load i32, i32* %1
	%3 = call i32 (i8*, ...) @printf(i8* bitcast ([15 x i8]* @string.e4b993b5f16d57ebba5166037b305638 to i8*), i32 %2)
	store i32 0, i32* %0
	br label %exit


exit:
	%4 = load i32, i32* %0
	ret i32 %4

}

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
