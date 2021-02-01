; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.counter.vtable.type = type { %global.counter* ()*, void (i8*, i1)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.counter = type { %global.counter.vtable.type*, i32, i32, i8*, void (i8*, i1)* }

@global.counter.vtable.data = global %global.counter.vtable.type { %global.counter* ()* @global.counter.create, void (i8*, i1)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.a = local_unnamed_addr global i32 0
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"
@string.3860584d72c3c1e2f6ebe3a6e165a60f = constant [8 x i8] c"default\00"
@string.43d2b345b462cd06e5f7841d678588a5 = constant [5 x i8] c"loop\00"
@string.e4774a52ffe13545ae9d5ce20dd839b9 = constant [9 x i8] c"a = %d \0A\00"

; Function Attrs: nofree nounwind
define i32 @main() local_unnamed_addr #0 {
entry:
  %0 = load i32, i32* @global.a, align 4
  %1 = add i32 %0, 1
  store i32 %1, i32* @global.a, align 4
  %2 = icmp eq i32 %1, 0
  br i1 %2, label %6, label %3

3:                                                ; preds = %6, %entry
  %4 = phi i32 [ %.pr, %6 ], [ %1, %entry ]
  %5 = icmp eq i32 %4, 1
  br i1 %5, label %12, label %8

6:                                                ; preds = %entry
  %7 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([7 x i8], [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268, i64 0, i64 0))
  %.pr = load i32, i32* @global.a, align 4
  br label %3

8:                                                ; preds = %12, %3
  %9 = phi i32 [ %.pre, %12 ], [ %4, %3 ]
  %switch.selectcmp = icmp eq i32 %9, 1
  %switch.select = select i1 %switch.selectcmp, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc, i64 0, i64 0), i8* getelementptr inbounds ([8 x i8], [8 x i8]* @string.3860584d72c3c1e2f6ebe3a6e165a60f, i64 0, i64 0)
  %switch.selectcmp5 = icmp eq i32 %9, 0
  %switch.select6 = select i1 %switch.selectcmp5, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268, i64 0, i64 0), i8* %switch.select
  %10 = tail call i32 @puts(i8* nonnull dereferenceable(1) %switch.select6)
  %.pr1 = load i32, i32* @global.a, align 4
  %11 = icmp slt i32 %.pr1, 5
  br i1 %11, label %.lr.ph, label %exit

12:                                               ; preds = %3
  %13 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([7 x i8], [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc, i64 0, i64 0))
  %.pre = load i32, i32* @global.a, align 4
  br label %8

exit:                                             ; preds = %.lr.ph, %8
  %.lcssa = phi i32 [ %.pr1, %8 ], [ %17, %.lr.ph ]
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.e4774a52ffe13545ae9d5ce20dd839b9, i64 0, i64 0), i32 %.lcssa)
  ret i32 0

.lr.ph:                                           ; preds = %8, %.lr.ph
  %15 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([5 x i8], [5 x i8]* @string.43d2b345b462cd06e5f7841d678588a5, i64 0, i64 0))
  %16 = load i32, i32* @global.a, align 4
  %17 = add i32 %16, 1
  store i32 %17, i32* @global.a, align 4
  %18 = icmp slt i32 %17, 5
  br i1 %18, label %.lr.ph, label %exit
}

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

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
  ret %global.counter* %1
}

; Function Attrs: nounwind
define void @global.counter.destroy(i8* nocapture %this, i1 %free) #1 {
entry:
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

attributes #0 = { nofree nounwind }
attributes #1 = { nounwind }
attributes #2 = { nofree norecurse nounwind }
