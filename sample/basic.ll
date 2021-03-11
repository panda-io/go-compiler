%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }
%global.counter.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.base_class = type { %global.base_class.vtable.type* }
%global.base_class.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)* }
%global.derive = type { %global.derive.vtable.type* }
%global.derive.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { i8* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.base_class.vtable.data = global %global.base_class.vtable.type { i8* ()* @global.base_class.create, void (i8*)* @global.base_class.destroy, void (i8*)* @global.base_class.do }
@global.derive.vtable.data = global %global.derive.vtable.type { i8* ()* @global.derive.create, void (i8*)* @global.derive.destroy, void (i8*)* @global.derive.do }
@string.726bd3560bd4c136648f7760895d8d62 = constant [18 x i8] c"base construction\00"
@string.362aeeddb3d01da539cb6755bde46953 = constant [17 x i8] c"base destruction\00"
@string.b58017d6f5ff15cba2431d7ec3967243 = constant [21 x i8] c"do something in base\00"
@string.33b7808bf372c3d58730520160cb2c15 = constant [20 x i8] c"derive construction\00"
@string.ef25b0542457581e67c27a0dddb7bda5 = constant [19 x i8] c"derive destruction\00"
@string.7cf7144b5f7d7f8893615fa04d42f3f7 = constant [23 x i8] c"do something in derive\00"
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
	%0 = call i32 @puts(i8* bitcast ([16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f to i8*))
	call void @free(i8* %this)
	br label %exit


exit:
	ret void

}

define void @global.counter.retain_shared(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([14 x i8]* @string.cf85dc053c0475520502efb2ba3c77a9 to i8*))
	%1 = bitcast i8* %this to %global.counter*
	%2 = getelementptr %global.counter, %global.counter* %1, i32 0, i32 1
	%3 = load i32, i32* %2
	%4 = add i32 %3, 1
	store i32 %4, i32* %2
	br label %exit


exit:
	ret void

}

define void @global.counter.release_shared(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([15 x i8]* @string.3b2e33db0bf2dc0aef4015e91829dab5 to i8*))
	%1 = icmp eq i8* %this, null
	br i1 %1, label %15, label %2


exit:
	ret void


2:
	%3 = bitcast i8* %this to %global.counter*
	%4 = getelementptr %global.counter, %global.counter* %3, i32 0, i32 1
	%5 = load i32, i32* %4
	%6 = sub i32 %5, 1
	store i32 %6, i32* %4
	%7 = bitcast i8* %this to %global.counter*
	%8 = getelementptr %global.counter, %global.counter* %7, i32 0, i32 1
	%9 = load i32, i32* %8
	%10 = call i32 (i8*, ...) @printf(i8* bitcast ([19 x i8]* @string.c52993c907c8f30bf3854dc0e21a7eca to i8*), i32 %9)
	%11 = bitcast i8* %this to %global.counter*
	%12 = getelementptr %global.counter, %global.counter* %11, i32 0, i32 1
	%13 = load i32, i32* %12
	%14 = icmp eq i32 %13, 0
	br i1 %14, label %17, label %16


15:
	br label %exit


16:
	br label %exit


17:
	%18 = bitcast i8* %this to %global.counter*
	%19 = getelementptr %global.counter, %global.counter* %18, i32 0, i32 4
	%20 = load void (i8*)*, void (i8*)** %19
	%21 = bitcast i8* %this to %global.counter*
	%22 = getelementptr %global.counter, %global.counter* %21, i32 0, i32 3
	%23 = load i8*, i8** %22
	call void %20(i8* %23)
	%24 = bitcast i8* %this to %global.counter*
	%25 = getelementptr %global.counter, %global.counter* %24, i32 0, i32 3
	%26 = load i8*, i8** %25
	%27 = call i32 (i8*, ...) @printf(i8* bitcast ([17 x i8]* @string.21c67ac9191c65481dbab306227b4840 to i8*), i8* %26)
	%28 = bitcast i8* %this to %global.counter*
	%29 = getelementptr %global.counter, %global.counter* %28, i32 0, i32 3
	%30 = load i8*, i8** %29
	call void @free(i8* %30)
	%31 = bitcast i8* %this to %global.counter*
	%32 = getelementptr %global.counter, %global.counter* %31, i32 0, i32 3
	%33 = load i8*, i8** %32
	%34 = bitcast i8* %this to %global.counter*
	%35 = getelementptr %global.counter, %global.counter* %34, i32 0, i32 3
	store i8* null, i8** %35
	%36 = bitcast i8* %this to %global.counter*
	%37 = getelementptr %global.counter, %global.counter* %36, i32 0, i32 2
	%38 = load i32, i32* %37
	%39 = icmp eq i32 %38, 0
	br i1 %39, label %41, label %40


40:
	br label %16


41:
	%42 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 to i8*), i8* %this)
	call void @free(i8* %this)
	br label %40

}

define void @global.counter.retain_weak(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([12 x i8]* @string.5927c4441dce664e4b461e529f933750 to i8*))
	%1 = bitcast i8* %this to %global.counter*
	%2 = getelementptr %global.counter, %global.counter* %1, i32 0, i32 2
	%3 = load i32, i32* %2
	%4 = add i32 %3, 1
	store i32 %4, i32* %2
	br label %exit


exit:
	ret void

}

define void @global.counter.release_weak(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([13 x i8]* @string.5662737e1a39fc068ead71add358dfd3 to i8*))
	%1 = icmp eq i8* %this, null
	br i1 %1, label %20, label %2


exit:
	ret void


2:
	%3 = bitcast i8* %this to %global.counter*
	%4 = getelementptr %global.counter, %global.counter* %3, i32 0, i32 2
	%5 = load i32, i32* %4
	%6 = sub i32 %5, 1
	store i32 %6, i32* %4
	%7 = bitcast i8* %this to %global.counter*
	%8 = getelementptr %global.counter, %global.counter* %7, i32 0, i32 2
	%9 = load i32, i32* %8
	%10 = call i32 (i8*, ...) @printf(i8* bitcast ([17 x i8]* @string.b6feae5df5d6172ffcb2a6bcd4d5c478 to i8*), i32 %9)
	%11 = bitcast i8* %this to %global.counter*
	%12 = getelementptr %global.counter, %global.counter* %11, i32 0, i32 1
	%13 = load i32, i32* %12
	%14 = icmp eq i32 %13, 0
	%15 = bitcast i8* %this to %global.counter*
	%16 = getelementptr %global.counter, %global.counter* %15, i32 0, i32 2
	%17 = load i32, i32* %16
	%18 = icmp eq i32 %17, 0
	%19 = and i1 %14, %18
	br i1 %19, label %22, label %21


20:
	br label %exit


21:
	br label %exit


22:
	%23 = call i32 (i8*, ...) @printf(i8* bitcast ([18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 to i8*), i8* %this)
	call void @free(i8* %this)
	br label %21

}

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = alloca i8*
	br label %body


body:
	%2 = call i8* @global.derive.create()
	%3 = call i8* @global.counter.create()
	call void @global.counter.retain_shared(i8* %3)
	%4 = bitcast i8* %3 to %global.counter*
	%5 = getelementptr %global.counter, %global.counter* %4, i32 0, i32 3
	store i8* %2, i8** %5
	%6 = bitcast i8* %3 to %global.counter*
	%7 = getelementptr %global.counter, %global.counter* %6, i32 0, i32 4
	store void (i8*)* @global.derive.destroy, void (i8*)** %7
	store i8* %3, i8** %1
	%8 = load i8*, i8** %1
	%9 = bitcast i8* %8 to %global.counter*
	%10 = getelementptr %global.counter, %global.counter* %9, i32 0, i32 3
	%11 = load i8*, i8** %10
	%12 = bitcast i8* %11 to %global.derive*
	%13 = getelementptr %global.derive, %global.derive* %12, i32 0, i32 0
	%14 = load %global.derive.vtable.type*, %global.derive.vtable.type** %13
	%15 = getelementptr %global.derive.vtable.type, %global.derive.vtable.type* %14, i32 0, i32 2
	%16 = load void (i8*)*, void (i8*)** %15
	call void %16(i8* %11)
	store i32 0, i32* %0
	br label %exit


exit:
	%17 = load i8*, i8** %1
	call void @global.counter.release_shared(i8* %17)
	%18 = load i32, i32* %0
	ret i32 %18

}

define i8* @global.base_class.create() {
entry:
	%0 = alloca i8*
	%1 = getelementptr %global.base_class, %global.base_class* null, i32 1
	%2 = ptrtoint %global.base_class* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	call void @memset(i8* %3, i32 0, i32 %2)
	%4 = bitcast i8* %3 to %global.base_class*
	%5 = getelementptr %global.base_class, %global.base_class* %4, i32 0, i32 0
	store %global.base_class.vtable.type* @global.base_class.vtable.data, %global.base_class.vtable.type** %5
	store i8* %3, i8** %0
	br label %body


body:
	%6 = call i32 @puts(i8* bitcast ([18 x i8]* @string.726bd3560bd4c136648f7760895d8d62 to i8*))
	br label %exit


exit:
	%7 = load i8*, i8** %0
	ret i8* %7

}

define void @global.base_class.destroy(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([17 x i8]* @string.362aeeddb3d01da539cb6755bde46953 to i8*))
	br label %exit


exit:
	ret void

}

define void @global.base_class.do(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([21 x i8]* @string.b58017d6f5ff15cba2431d7ec3967243 to i8*))
	br label %exit


exit:
	ret void

}

define i8* @global.derive.create() {
entry:
	%0 = alloca i8*
	%1 = getelementptr %global.derive, %global.derive* null, i32 1
	%2 = ptrtoint %global.derive* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	call void @memset(i8* %3, i32 0, i32 %2)
	%4 = bitcast i8* %3 to %global.derive*
	%5 = getelementptr %global.derive, %global.derive* %4, i32 0, i32 0
	store %global.derive.vtable.type* @global.derive.vtable.data, %global.derive.vtable.type** %5
	store i8* %3, i8** %0
	br label %body


body:
	%6 = call i32 @puts(i8* bitcast ([20 x i8]* @string.33b7808bf372c3d58730520160cb2c15 to i8*))
	br label %exit


exit:
	%7 = load i8*, i8** %0
	ret i8* %7

}

define void @global.derive.destroy(i8* %this) {
entry:
	br label %body


body:
	call void @global.base_class.destroy(i8* %this)
	%0 = call i32 @puts(i8* bitcast ([19 x i8]* @string.ef25b0542457581e67c27a0dddb7bda5 to i8*))
	br label %exit


exit:
	ret void

}

define void @global.derive.do(i8* %this) {
entry:
	br label %body


body:
	%0 = call i32 @puts(i8* bitcast ([23 x i8]* @string.7cf7144b5f7d7f8893615fa04d42f3f7 to i8*))
	br label %exit


exit:
	ret void

}
