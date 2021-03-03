; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.base.vtable.type = type { %global.base* ()*, void (i8*)* }
%global.base = type { %global.base.vtable.type* }
%global.counter.vtable.type = type { %global.counter* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*)* }

@global.base.vtable.data = global %global.base.vtable.type { %global.base* ()* @global.base.create, void (i8*)* @global.base.destroy }
@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.319f93d592b03bf2f832e7b2592ace2e = constant [17 x i8] c"base contruction\00"
@string.362aeeddb3d01da539cb6755bde46953 = constant [17 x i8] c"base destruction\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare noalias i8* @malloc(i32) local_unnamed_addr #0

; Function Attrs: nounwind
declare void @free(i8* nocapture) local_unnamed_addr #1

declare void @memset(i8*, i32, i32) local_unnamed_addr

define %global.counter* @global.counter.create() {
entry:
  %0 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %0, i32 0, i32 32)
  %1 = bitcast i8* %0 to %global.counter*
  %2 = bitcast i8* %0 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %2, align 8
  %3 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e, i64 0, i64 0))
  ret %global.counter* %1
}

; Function Attrs: nofree nounwind
define void @global.counter.destroy(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree norecurse nounwind
define void @global.counter.retain_shared(i8* nocapture %this) #2 {
entry:
  %0 = getelementptr i8, i8* %this, i64 8
  %1 = bitcast i8* %0 to i32*
  %2 = load i32, i32* %1, align 4
  %3 = add i32 %2, 1
  store i32 %3, i32* %1, align 4
  ret void
}

; Function Attrs: nofree norecurse nounwind
define void @global.counter.release_shared(i8* %this) #2 {
entry:
  %0 = icmp eq i8* %this, null
  br i1 %0, label %exit, label %1

exit:                                             ; preds = %entry
  ret void

1:                                                ; preds = %entry
  %2 = getelementptr i8, i8* %this, i64 8
  %3 = bitcast i8* %2 to i32*
  %4 = load i32, i32* %3, align 4
  %5 = add i32 %4, -1
  store i32 %5, i32* %3, align 4
  ret void
}

; Function Attrs: nofree norecurse nounwind
define void @global.counter.retain_weak(i8* nocapture %this) #2 {
entry:
  %0 = getelementptr i8, i8* %this, i64 12
  %1 = bitcast i8* %0 to i32*
  %2 = load i32, i32* %1, align 4
  %3 = add i32 %2, 1
  store i32 %3, i32* %1, align 4
  ret void
}

; Function Attrs: nounwind
define void @global.counter.release_weak(i8* %this) #1 {
entry:
  %0 = icmp eq i8* %this, null
  br i1 %0, label %exit, label %1

exit:                                             ; preds = %1, %11, %entry
  ret void

1:                                                ; preds = %entry
  %2 = getelementptr i8, i8* %this, i64 12
  %3 = bitcast i8* %2 to i32*
  %4 = load i32, i32* %3, align 4
  %5 = add i32 %4, -1
  store i32 %5, i32* %3, align 4
  %6 = getelementptr i8, i8* %this, i64 8
  %7 = bitcast i8* %6 to i32*
  %8 = load i32, i32* %7, align 4
  %9 = or i32 %8, %5
  %10 = icmp eq i32 %9, 0
  br i1 %10, label %11, label %exit

11:                                               ; preds = %1
  tail call void @free(i8* nonnull %this)
  br label %exit
}

define i32 @main() local_unnamed_addr {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.base.vtable.type**
  store %global.base.vtable.type* @global.base.vtable.data, %global.base.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.319f93d592b03bf2f832e7b2592ace2e, i64 0, i64 0))
  %3 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %3, i32 0, i32 32)
  %4 = bitcast i8* %3 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %4, align 8
  %5 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e, i64 0, i64 0))
  %6 = getelementptr i8, i8* %3, i64 16
  %7 = bitcast i8* %6 to i8**
  store i8* %0, i8** %7, align 8
  %8 = getelementptr i8, i8* %3, i64 24
  %9 = bitcast i8* %8 to void (i8*)**
  store void (i8*)* @global.base.destroy, void (i8*)** %9, align 8
  %10 = icmp eq i8* %3, null
  br i1 %10, label %global.counter.release_shared.exit, label %11

11:                                               ; preds = %entry
  %12 = getelementptr i8, i8* %3, i64 8
  %13 = bitcast i8* %12 to i32*
  %14 = load i32, i32* %13, align 4
  %15 = add i32 %14, -1
  store i32 %15, i32* %13, align 4
  br label %global.counter.release_shared.exit

global.counter.release_shared.exit:               ; preds = %entry, %11
  ret i32 0
}

define %global.base* @global.base.create() {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.base*
  %2 = bitcast i8* %0 to %global.base.vtable.type**
  store %global.base.vtable.type* @global.base.vtable.data, %global.base.vtable.type** %2, align 8
  %3 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.319f93d592b03bf2f832e7b2592ace2e, i64 0, i64 0))
  ret %global.base* %1
}

; Function Attrs: nofree nounwind
define void @global.base.destroy(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.362aeeddb3d01da539cb6755bde46953, i64 0, i64 0))
  ret void
}

attributes #0 = { nofree nounwind }
attributes #1 = { nounwind }
attributes #2 = { nofree norecurse nounwind }
