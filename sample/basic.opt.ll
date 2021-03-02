; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.counter.vtable.type = type { %global.counter* ()*, void (i8*, i1)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*, i1)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (i8*, i1)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@global.a = local_unnamed_addr global i32 0
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

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

; Function Attrs: nounwind
define void @global.counter.destroy(i8* nocapture %this, i1 %free) #1 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f, i64 0, i64 0))
  tail call void @free(i8* %this)
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
define void @global.counter.release_shared(i8* nocapture %this) #2 {
entry:
  %0 = getelementptr i8, i8* %this, i64 8
  %1 = bitcast i8* %0 to i32*
  %2 = load i32, i32* %1, align 4
  %3 = add i32 %2, -1
  store i32 %3, i32* %1, align 4
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
define void @global.counter.release_weak(i8* nocapture %this) #1 {
entry:
  %0 = getelementptr i8, i8* %this, i64 12
  %1 = bitcast i8* %0 to i32*
  %2 = load i32, i32* %1, align 4
  %3 = add i32 %2, -1
  store i32 %3, i32* %1, align 4
  %4 = getelementptr i8, i8* %this, i64 8
  %5 = bitcast i8* %4 to i32*
  %6 = load i32, i32* %5, align 4
  %7 = or i32 %6, %3
  %8 = icmp eq i32 %7, 0
  br i1 %8, label %9, label %exit

exit:                                             ; preds = %9, %entry
  ret void

9:                                                ; preds = %entry
  tail call void @free(i8* nonnull %this)
  br label %exit
}

define i32 @main() local_unnamed_addr {
entry:
  %0 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %0, i32 0, i32 32)
  %1 = bitcast i8* %0 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e, i64 0, i64 0))
  %3 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i64 0, i64 0))
  %4 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f, i64 0, i64 0)) #1
  tail call void @free(i8* %0) #1
  ret i32 0
}

attributes #0 = { nofree nounwind }
attributes #1 = { nounwind }
attributes #2 = { nofree norecurse nounwind }
