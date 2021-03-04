%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }
%global.counter.vtable.type = type { %global.counter* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.base = type { %global.base.vtable.type* }
%global.base.vtable.type = type { %global.base* ()*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.base.vtable.data = global %global.base.vtable.type { %global.base* ()* @global.base.create, void (i8*)* @global.base.destroy }
@string.319f93d592b03bf2f832e7b2592ace2e = constant [17 x i8] c"base contruction\00"
@string.362aeeddb3d01da539cb6755bde46953 = constant [17 x i8] c"base destruction\00"
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.cf85dc053c0475520502efb2ba3c77a9 = constant [14 x i8] c"retain shared\00"
@string.3b2e33db0bf2dc0aef4015e91829dab5 = constant [15 x i8] c"release shared\00"
@string.c52993c907c8f30bf3854dc0e21a7eca = constant [19 x i8] c"shared count: %d \0A\00"
@string.21c67ac9191c65481dbab306227b4840 = constant [17 x i8] c"free object %p \0A\00"
@string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 = constant [18 x i8] c"free counter %p \0A\00"
@string.5927c4441dce664e4b461e529f933750 = constant [12 x i8] c"retain weak\00"
@string.5662737e1a39fc068ead71add358dfd3 = constant [13 x i8] c"release weak\00"
@string.b6feae5df5d6172ffcb2a6bcd4d5c478 = constant [17 x i8] c"weak count: %d \0A\00"

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
	%1 = call i32 @puts(i8* bitcast ([14 x i8]* @string.cf85dc053c0475520502efb2ba3c77a9 to i8*))
	%2 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%3 = load i32, i32* %2
	%4 = add i32 %3, 1
	store i32 %4, i32* %2
	br label %exit


exit:
	ret void

}

define void @global.counter.release_shared(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = call i32 @puts(i8* bitcast ([15 x i8]* @string.3b2e33db0bf2dc0aef4015e91829dab5 to i8*))
	%2 = icmp eq %global.counter* %0, null
	br i1 %2, label %14, label %3


exit:
	ret void


3:
	%4 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%5 = load i32, i32* %4
	%6 = sub i32 %5, 1
	store i32 %6, i32* %4
	%7 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%8 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%9 = load i32, i32* %8
	%10 = call i32 (i8*, ...) @printf(i8* bitcast ([19 x i8]* @string.c52993c907c8f30bf3854dc0e21a7eca to i8*), i32 %9)
	%11 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%12 = load i32, i32* %11
	%13 = icmp eq i32 %12, 0
	br i1 %13, label %16, label %15


14:
	br label %exit


15:
	br label %exit


16:
	%17 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%18 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%19 = load i8*, i8** %18
	%20 = call i32 (i8*, ...) @printf(i8* bitcast ([17 x i8]* @string.21c67ac9191c65481dbab306227b4840 to i8*), i8* %19)
	%21 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 4
	%22 = load void (i8*)*, void (i8*)** %21
	%23 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%24 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%25 = load i8*, i8** %24
	call void %22(i8* %25)
	%26 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%27 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 3
	%28 = load i8*, i8** %27
	call void @free(i8* %28)
	%29 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%30 = load i32, i32* %29
	%31 = icmp eq i32 %30, 0
	br i1 %31, label %33, label %32


32:
	br label %15


33:
	%34 = bitcast %global.counter* %0 to i8*
	%35 = bitcast %global.counter* %0 to i8*
	%36 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 to i8*), i8* %35)
	%37 = bitcast %global.counter* %0 to i8*
	%38 = bitcast %global.counter* %0 to i8*
	call void @free(i8* %38)
	br label %32

}

define void @global.counter.retain_weak(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = call i32 @puts(i8* bitcast ([12 x i8]* @string.5927c4441dce664e4b461e529f933750 to i8*))
	%2 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%3 = load i32, i32* %2
	%4 = add i32 %3, 1
	store i32 %4, i32* %2
	br label %exit


exit:
	ret void

}

define void @global.counter.release_weak(i8* %this) {
entry:
	%0 = bitcast i8* %this to %global.counter*
	br label %body


body:
	%1 = call i32 @puts(i8* bitcast ([13 x i8]* @string.5662737e1a39fc068ead71add358dfd3 to i8*))
	%2 = icmp eq %global.counter* %0, null
	br i1 %2, label %18, label %3


exit:
	ret void


3:
	%4 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%5 = load i32, i32* %4
	%6 = sub i32 %5, 1
	store i32 %6, i32* %4
	%7 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%8 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%9 = load i32, i32* %8
	%10 = call i32 (i8*, ...) @printf(i8* bitcast ([17 x i8]* @string.b6feae5df5d6172ffcb2a6bcd4d5c478 to i8*), i32 %9)
	%11 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 1
	%12 = load i32, i32* %11
	%13 = icmp eq i32 %12, 0
	%14 = getelementptr %global.counter, %global.counter* %0, i32 0, i32 2
	%15 = load i32, i32* %14
	%16 = icmp eq i32 %15, 0
	%17 = and i1 %13, %16
	br i1 %17, label %20, label %19


18:
	br label %exit


19:
	br label %exit


20:
	%21 = bitcast %global.counter* %0 to i8*
	%22 = bitcast %global.counter* %0 to i8*
	%23 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 to i8*), i8* %22)
	%24 = bitcast %global.counter* %0 to i8*
	%25 = bitcast %global.counter* %0 to i8*
	call void @free(i8* %25)
	br label %19

}

define i32 @main() {
entry:
	%0 = alloca i32
	br label %body


body:
	%1 = call %global.base* @global.base.create()
	%2 = call %global.counter* @global.counter.create()
	%3 = bitcast %global.counter* %2 to i8*
	call void @global.counter.retain_shared(i8* %3)
	%4 = getelementptr %global.counter, %global.counter* %2, i32 0, i32 3
	%5 = bitcast %global.base* %1 to i8*
	store i8* %5, i8** %4
	%6 = getelementptr %global.counter, %global.counter* %2, i32 0, i32 4
	store void (i8*)* @global.base.destroy, void (i8*)** %6
	store i32 0, i32* %0
	br label %exit


exit:
	%7 = bitcast %global.counter* %2 to i8*
	call void @global.counter.release_shared(i8* %7)
	%8 = load i32, i32* %0
	ret i32 %8

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
