@global.allocator = global { i8*, i32 }

declare i32 @printf(i8* %0)

declare i8* @malloc(i32 %0)

declare i8* @realloc(i8* %0, i32 %1)

declare void @free(i8* %0)

declare i32 @memcmp(i8* %0, i8* %1, i32 %2)

declare void @memcpy(i8* %0, i8* %1, i32 %2)

declare void @memset(i8* %0, i32 %1, i32 %2)

define void @global.allocator.allocate(%global.allocator* %0, i32 %1) {
2:
	ret <nil> null
}

define void @global.allocator.destroy(%global.allocator* %0) {
1:
	ret <nil> null
}
