; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.counter.vtable.type = type { %global.counter* ()*, void (%global.counter*, i1)*, void (%global.counter*)*, void (%global.counter*)*, void (%global.counter*)*, i32 (%global.counter*)*, i32 (%global.counter*)*, i8* (%global.counter*)* }
%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (%global.counter*, i1)* @global.counter.destroy, void (%global.counter*)* @global.counter.retain_shared, void (%global.counter*)* @global.counter.retain_weak, void (%global.counter*)* @global.counter.release_weak, i32 (%global.counter*)* @global.counter.shared_count, i32 (%global.counter*)* @global.counter.weak_count, i8* (%global.counter*)* @global.counter.get_object }

; Function Attrs: norecurse nounwind readnone
define i32 @main() local_unnamed_addr #0 {
entry:
  ret i32 0
}

; Function Attrs: nofree nounwind
declare noalias i8* @malloc(i32) local_unnamed_addr #1

; Function Attrs: nounwind
declare void @free(i8* nocapture) local_unnamed_addr #2

declare void @memset(i8*, i32, i32) local_unnamed_addr

define %global.counter* @global.counter.create() {
entry:
  %0 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %0, i32 0, i32 32)
  %1 = bitcast i8* %0 to %global.counter*
  %2 = bitcast i8* %0 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %2, align 8
  ret %global.counter* %1
}

; Function Attrs: nounwind
define void @global.counter.destroy(%global.counter* nocapture %this, i1 %free) #2 {
entry:
  %0 = bitcast %global.counter* %this to i8*
  tail call void @free(i8* %0)
  ret void
}

; Function Attrs: nofree norecurse nounwind
define void @global.counter.retain_shared(%global.counter* nocapture %this) #3 {
entry:
  %0 = getelementptr %global.counter, %global.counter* %this, i64 0, i32 1
  %1 = load i32, i32* %0, align 4
  %2 = add i32 %1, 1
  store i32 %2, i32* %0, align 4
  ret void
}

; Function Attrs: nofree norecurse nounwind
define void @global.counter.retain_weak(%global.counter* nocapture %this) #3 {
entry:
  %0 = getelementptr %global.counter, %global.counter* %this, i64 0, i32 2
  %1 = load i32, i32* %0, align 4
  %2 = add i32 %1, 1
  store i32 %2, i32* %0, align 4
  ret void
}

; Function Attrs: nofree norecurse nounwind
define void @global.counter.release_weak(%global.counter* nocapture %this) #3 {
entry:
  %0 = getelementptr %global.counter, %global.counter* %this, i64 0, i32 2
  %1 = load i32, i32* %0, align 4
  %2 = add i32 %1, -1
  store i32 %2, i32* %0, align 4
  ret void
}

; Function Attrs: norecurse nounwind readonly
define i32 @global.counter.shared_count(%global.counter* nocapture readonly %this) #4 {
entry:
  %0 = getelementptr %global.counter, %global.counter* %this, i64 0, i32 1
  %1 = load i32, i32* %0, align 4
  ret i32 %1
}

; Function Attrs: norecurse nounwind readonly
define i32 @global.counter.weak_count(%global.counter* nocapture readonly %this) #4 {
entry:
  %0 = getelementptr %global.counter, %global.counter* %this, i64 0, i32 2
  %1 = load i32, i32* %0, align 4
  ret i32 %1
}

; Function Attrs: norecurse nounwind readonly
define i8* @global.counter.get_object(%global.counter* nocapture readonly %this) #4 {
entry:
  %0 = getelementptr %global.counter, %global.counter* %this, i64 0, i32 3
  %1 = load i8*, i8** %0, align 8
  ret i8* %1
}

attributes #0 = { norecurse nounwind readnone }
attributes #1 = { nofree nounwind }
attributes #2 = { nounwind }
attributes #3 = { nofree norecurse nounwind }
attributes #4 = { norecurse nounwind readonly }
