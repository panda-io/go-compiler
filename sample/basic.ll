@global.color.red = global i32 0
@global.color.green = global i32 1
@global.color.blue = global i32 2
@global.counter = global { %global.counter.vtable.type*, i32, i32, i8*, void (i8*) }
@global.counter.vtable.type = global { void (), void (%global.counter*), void (%global.counter*), void (%global.counter*), void (%global.counter*), void (%global.counter*), i32 (%global.counter*), i32 (%global.counter*), i8* (%global.counter*) }
@global.counter.vtable.data = global %global.counter.vtable.type { void ()* @global.counter.create, void (%global.counter*)* @global.counter.destroy, void (%global.counter*)* @global.counter.retain_shared, void (%global.counter*)* @global.counter.release_shared, void (%global.counter*)* @global.counter.retain_weak, void (%global.counter*)* @global.counter.release_weak, i32 (%global.counter*)* @global.counter.shared_count, i32 (%global.counter*)* @global.counter.weak_count, i8* (%global.counter*)* @global.counter.get_object }
@global.nothing = global { %global.nothing.vtable.type* }
@global.nothing.vtable.type = global { void (), void (%global.nothing*) }
@global.nothing.vtable.data = global %global.nothing.vtable.type { void ()* @global.nothing.create, void (%global.nothing*)* @global.nothing.destroy }

declare void @global.counter.create()

declare void @global.counter.destroy(%global.counter* %this)

define void @global.counter.retain_shared(%global.counter* %this) {
entry:
	ret void
}

define void @global.counter.release_shared(%global.counter* %this) {
entry:
	ret void
}

define void @global.counter.retain_weak(%global.counter* %this) {
entry:
	ret void
}

define void @global.counter.release_weak(%global.counter* %this) {
entry:
	ret void
}

define i32 @global.counter.shared_count(%global.counter* %this) {
entry:
	ret void
}

define i32 @global.counter.weak_count(%global.counter* %this) {
entry:
	ret void
}

define i8* @global.counter.get_object(%global.counter* %this) {
entry:
	ret void
}

declare void @global.nothing.create()

declare void @global.nothing.destroy(%global.nothing* %this)

define i32 @global.add(i32 %a, i32 %b) {
entry:
	%0 = alloca i32
	store i32 %a, i32* %0
	%1 = alloca i32
	store i32 %b, i32* %1
	%2 = load i32, i32* %0
	%3 = load i32, i32* %1
	%4 = add i32 %2, %3
	ret i32 %4
}

define i32 @main() {
entry:
	ret i32 0
}
