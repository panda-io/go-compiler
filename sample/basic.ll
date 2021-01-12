@global.allocator = global { i8*, i32 }

declare i32 @printf(i8* %text)

declare i8* @malloc(i32 %size)

declare i8* @realloc(i8* %address, i32 %size)

declare void @free(i8* %address)

declare i32 @memcmp(i8* %dest, i8* %source, i32 %size)

declare void @memcpy(i8* %dest, i8* %source, i32 %size)

declare void @memset(i8* %source, i32 %value, i32 %size)

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
