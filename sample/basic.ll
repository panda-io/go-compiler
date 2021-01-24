%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }
%global.counter.vtable.type = type { %global.counter* ()*, void (%global.counter*)*, void (%global.counter*)*, void (%global.counter*)*, void (%global.counter*)*, i32 (%global.counter*)*, i32 (%global.counter*)*, i8* (%global.counter*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (%global.counter*)* @global.counter.destroy, void (%global.counter*)* @global.counter.retain_shared, void (%global.counter*)* @global.counter.retain_weak, void (%global.counter*)* @global.counter.release_weak, i32 (%global.counter*)* @global.counter.shared_count, i32 (%global.counter*)* @global.counter.weak_count, i8* (%global.counter*)* @global.counter.get_object }
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

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
	br label %exit

exit:
	%6 = load %global.counter*, %global.counter** %0
	ret %global.counter* %6
}

define void @global.counter.destroy(%global.counter* %this) {
entry:
	br label %exit

exit:
	%0 = bitcast %global.counter* %this to i8*
	call void @free(i8* %0)
	ret void
}

define void @global.counter.retain_shared(%global.counter* %this) {
entry:
	%0 = getelementptr %global.counter, %global.counter* %this, i32 0, i32 1
	%1 = load i32, i32* %0
	%2 = add i32 %1, 1
	store i32 %2, i32* %0
	br label %exit

exit:
	ret void
}

define void @global.counter.retain_weak(%global.counter* %this) {
entry:
	%0 = getelementptr %global.counter, %global.counter* %this, i32 0, i32 2
	%1 = load i32, i32* %0
	%2 = add i32 %1, 1
	store i32 %2, i32* %0
	br label %exit

exit:
	ret void
}

define void @global.counter.release_weak(%global.counter* %this) {
entry:
	%0 = getelementptr %global.counter, %global.counter* %this, i32 0, i32 2
	%1 = load i32, i32* %0
	%2 = sub i32 %1, 1
	store i32 %2, i32* %0
	br label %exit

exit:
	ret void
}

define i32 @global.counter.shared_count(%global.counter* %this) {
entry:
	%0 = alloca i32
	%1 = getelementptr %global.counter, %global.counter* %this, i32 0, i32 1
	%2 = load i32, i32* %1
	store i32 %2, i32* %0
	br label %exit

exit:
	%3 = load i32, i32* %0
	ret i32 %3
}

define i32 @global.counter.weak_count(%global.counter* %this) {
entry:
	%0 = alloca i32
	%1 = getelementptr %global.counter, %global.counter* %this, i32 0, i32 2
	%2 = load i32, i32* %1
	store i32 %2, i32* %0
	br label %exit

exit:
	%3 = load i32, i32* %0
	ret i32 %3
}

define i8* @global.counter.get_object(%global.counter* %this) {
entry:
	%0 = alloca i8*
	%1 = getelementptr %global.counter, %global.counter* %this, i32 0, i32 3
	%2 = load i8*, i8** %1
	store i8* %2, i8** %0
	br label %exit

exit:
	%3 = load i8*, i8** %0
	ret i8* %3
}

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = getelementptr [12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i32 0, i32 0
	%2 = call i32 @puts(i8* %1)
	store i32 0, i32* %0
	br label %exit

exit:
	%3 = load i32, i32* %0
	ret i32 %3
}
