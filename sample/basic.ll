%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }
%global.counter.vtable.type = type { %global.counter* ()*, void (%global.counter*)*, void (%global.counter*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (%global.counter*)* @global.counter.destroy, void (%global.counter*)* @global.counter.retain_shared }

define %global.counter* @global.counter.create() {
entry:
	%0 = getelementptr %global.counter, %global.counter* null, i32 1
	%1 = ptrtoint %global.counter* %0 to i32
	%2 = call i8* @malloc(i32 %1)
	call void @memset(i8* %2, i32 0, i32 %1)
	%3 = bitcast i8* %2 to %global.counter*
	%4 = getelementptr %global.counter, %global.counter* %3, i32 0, i32 0
	store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %4
	ret %global.counter* %3
}

define void @global.counter.destroy(%global.counter* %this) {
entry:
	%0 = bitcast %global.counter* %this to i8*
	call void @free(i8* %0)
	ret void
}

define void @global.counter.retain_shared(%global.counter* %this) {
entry:
	ret void
}

declare i32 @puts(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

define i32 @main() {
entry:
	ret i32 0
}
