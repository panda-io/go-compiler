; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.counter.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { i8* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.c1432ab71496ebb8b3e30bbcf37605e7 = constant [20 x i8] c"retain shared: %d \0A\00"
@string.ce99a7174da84f1767b9d5235ea3c24f = constant [21 x i8] c"release shared: %d \0A\00"
@string.21c67ac9191c65481dbab306227b4840 = constant [17 x i8] c"free object %p \0A\00"
@string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 = constant [18 x i8] c"free counter %p \0A\00"
@string.8d9c52192bdfa908703a004b070ff63e = constant [18 x i8] c"retain weak: %d \0A\00"
@string.ad2ae91c8542c824c6842efed6523f49 = constant [19 x i8] c"release weak: %d \0A\00"
@string.e4b993b5f16d57ebba5166037b305638 = constant [15 x i8] c"value #3: %d \0A\00"
@string.5bddf146c13b387514280200e83cf08b = constant [15 x i8] c"value #1: %d \0A\00"
@string.aee0f3d368512408f6bd5274bf51a219 = constant [15 x i8] c"value #2: %d \0A\00"

define i8* @global.counter.create() {
entry:
  %0 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %0, i32 0, i32 32)
  %1 = bitcast i8* %0 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e, i64 0, i64 0))
  ret i8* %0
}

; Function Attrs: nounwind
define void @global.counter.destroy(i8* nocapture %this) #0 {
entry:
  tail call void @free(i8* %this)
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @global.counter.retain_shared(i8* nocapture %this) #1 {
entry:
  %0 = getelementptr i8, i8* %this, i64 8
  %1 = bitcast i8* %0 to i32*
  %2 = load i32, i32* %1, align 4
  %3 = add i32 %2, 1
  store i32 %3, i32* %1, align 4
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.c1432ab71496ebb8b3e30bbcf37605e7, i64 0, i64 0), i32 %3)
  ret void
}

define void @global.counter.release_shared(i8* %this) {
entry:
  %0 = icmp eq i8* %this, null
  br i1 %0, label %exit, label %1

exit:                                             ; preds = %9, %23, %1, %entry
  ret void

1:                                                ; preds = %entry
  %2 = getelementptr i8, i8* %this, i64 8
  %3 = bitcast i8* %2 to i32*
  %4 = load i32, i32* %3, align 4
  %5 = add i32 %4, -1
  store i32 %5, i32* %3, align 4
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([21 x i8], [21 x i8]* @string.ce99a7174da84f1767b9d5235ea3c24f, i64 0, i64 0), i32 %5)
  %7 = load i32, i32* %3, align 4
  %8 = icmp eq i32 %7, 0
  br i1 %8, label %9, label %exit

9:                                                ; preds = %1
  %10 = getelementptr i8, i8* %this, i64 24
  %11 = bitcast i8* %10 to void (i8*)**
  %12 = load void (i8*)*, void (i8*)** %11, align 8
  %13 = getelementptr i8, i8* %this, i64 16
  %14 = bitcast i8* %13 to i8**
  %15 = load i8*, i8** %14, align 8
  tail call void %12(i8* %15)
  %16 = load i8*, i8** %14, align 8
  %17 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.21c67ac9191c65481dbab306227b4840, i64 0, i64 0), i8* %16)
  %18 = load i8*, i8** %14, align 8
  tail call void @free(i8* %18)
  store i8* null, i8** %14, align 8
  %19 = getelementptr i8, i8* %this, i64 12
  %20 = bitcast i8* %19 to i32*
  %21 = load i32, i32* %20, align 4
  %22 = icmp eq i32 %21, 0
  br i1 %22, label %23, label %exit

23:                                               ; preds = %9
  %24 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0, i64 0, i64 0), i8* nonnull %this)
  tail call void @free(i8* nonnull %this)
  br label %exit
}

; Function Attrs: nofree nounwind
define void @global.counter.retain_weak(i8* nocapture %this) #1 {
entry:
  %0 = getelementptr i8, i8* %this, i64 12
  %1 = bitcast i8* %0 to i32*
  %2 = load i32, i32* %1, align 4
  %3 = add i32 %2, 1
  store i32 %3, i32* %1, align 4
  %4 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.8d9c52192bdfa908703a004b070ff63e, i64 0, i64 0), i32 %3)
  ret void
}

; Function Attrs: nounwind
define void @global.counter.release_weak(i8* %this) #0 {
entry:
  %0 = icmp eq i8* %this, null
  br i1 %0, label %exit, label %1

exit:                                             ; preds = %1, %13, %entry
  ret void

1:                                                ; preds = %entry
  %2 = getelementptr i8, i8* %this, i64 12
  %3 = bitcast i8* %2 to i32*
  %4 = load i32, i32* %3, align 4
  %5 = add i32 %4, -1
  store i32 %5, i32* %3, align 4
  %6 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.ad2ae91c8542c824c6842efed6523f49, i64 0, i64 0), i32 %5)
  %7 = getelementptr i8, i8* %this, i64 8
  %8 = bitcast i8* %7 to i32*
  %9 = load i32, i32* %8, align 4
  %10 = load i32, i32* %3, align 4
  %11 = or i32 %10, %9
  %12 = icmp eq i32 %11, 0
  br i1 %12, label %13, label %exit

13:                                               ; preds = %1
  %14 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0, i64 0, i64 0), i8* nonnull %this)
  tail call void @free(i8* nonnull %this)
  br label %exit
}

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #1

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #1

; Function Attrs: nofree nounwind
declare noalias i8* @malloc(i32) local_unnamed_addr #1

; Function Attrs: nounwind
declare void @free(i8* nocapture) local_unnamed_addr #0

declare void @memset(i8*, i32, i32) local_unnamed_addr

; Function Attrs: nofree nounwind
define i32 @main() local_unnamed_addr #1 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.e4b993b5f16d57ebba5166037b305638, i64 0, i64 0), i32 1)
  ret i32 0
}

; Function Attrs: nofree nounwind
define void @global.print_number(i32* nocapture readonly %value) local_unnamed_addr #1 {
entry:
  %0 = load i32, i32* %value, align 4
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bddf146c13b387514280200e83cf08b, i64 0, i64 0), i32 %0)
  %2 = load i32, i32* %value, align 4
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.aee0f3d368512408f6bd5274bf51a219, i64 0, i64 0), i32 %2)
  ret void
}

attributes #0 = { nounwind }
attributes #1 = { nofree nounwind }
