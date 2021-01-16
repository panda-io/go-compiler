@global.allocator = global { %global.allocator.vtable.type*, i8*, i32 }
@global.allocator.vtable.type = global { void (%global.allocator*, i32), void (%global.allocator*) }
@global.allocator.vtable.data = global %global.allocator.vtable.type { void (%global.allocator*, i32)* @global.allocator.allocate, void (%global.allocator*)* @global.allocator.destroy }
@global.string = global { %global.string.vtable.type*, i8*, i32, i32 }
@global.string.vtable.type = global { void (%global.string*, i32), void (%global.allocator*) }
@global.string.vtable.data = global %global.string.vtable.type { void (%global.string*, i32)* @global.string.allocate, void (%global.allocator*)* @global.allocator.destroy }

define void @global.allocator.allocate(%global.allocator* %this, i32 %size) {
entry:
	%0 = load %global.allocator*, %global.allocator* %this
	%1 = load i32, i32 %size
	ret <nil> null
}

define void @global.allocator.destroy(%global.allocator* %this) {
entry:
	%0 = load %global.allocator*, %global.allocator* %this
	ret <nil> null
}

define void @global.string.allocate(%global.string* %this, i32 %size) {
entry:
	%0 = load %global.string*, %global.string* %this
	%1 = load i32, i32 %size
	ret <nil> null
}

declare i32 @printf(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)
