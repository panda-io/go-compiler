; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.counter.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.base_class.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)* }
%global.derive_class.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { i8* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.base_class.vtable.data = global %global.base_class.vtable.type { i8* ()* @global.base_class.create, void (i8*)* @global.base_class.destroy, void (i8*)* @global.base_class.echo }
@global.derive_class.vtable.data = global %global.derive_class.vtable.type { i8* ()* @global.derive_class.create, void (i8*)* @global.derive_class.destroy, void (i8*)* @global.derive_class.echo }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.c1432ab71496ebb8b3e30bbcf37605e7 = constant [20 x i8] c"retain shared: %d \0A\00"
@string.ce99a7174da84f1767b9d5235ea3c24f = constant [21 x i8] c"release shared: %d \0A\00"
@string.21c67ac9191c65481dbab306227b4840 = constant [17 x i8] c"free object %p \0A\00"
@string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 = constant [18 x i8] c"free counter %p \0A\00"
@string.8d9c52192bdfa908703a004b070ff63e = constant [18 x i8] c"retain weak: %d \0A\00"
@string.ad2ae91c8542c824c6842efed6523f49 = constant [19 x i8] c"release weak: %d \0A\00"
@string.726bd3560bd4c136648f7760895d8d62 = constant [18 x i8] c"base construction\00"
@string.362aeeddb3d01da539cb6755bde46953 = constant [17 x i8] c"base destruction\00"
@string.9bcbb503bda6c8ad83f772846b706f08 = constant [13 x i8] c"echo in base\00"
@string.33b7808bf372c3d58730520160cb2c15 = constant [20 x i8] c"derive construction\00"
@string.ef25b0542457581e67c27a0dddb7bda5 = constant [19 x i8] c"derive destruction\00"
@string.895758554639f423e017c6610cbf460b = constant [15 x i8] c"echo in derive\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare noalias i8* @malloc(i32) local_unnamed_addr #0

; Function Attrs: nounwind
declare void @free(i8* nocapture) local_unnamed_addr #1

declare void @memset(i8*, i32, i32) local_unnamed_addr

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
define void @global.counter.destroy(i8* nocapture %this) #1 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([16 x i8], [16 x i8]* @string.f8f86b3941cca26e8c147322b9a8309f, i64 0, i64 0))
  tail call void @free(i8* %this)
  ret void
}

; Function Attrs: nofree nounwind
define void @global.counter.retain_shared(i8* nocapture %this) #0 {
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
define void @global.counter.retain_weak(i8* nocapture %this) #0 {
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
define void @global.counter.release_weak(i8* %this) #1 {
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

define i32 @main() local_unnamed_addr {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.derive_class.vtable.type**
  store %global.derive_class.vtable.type* @global.derive_class.vtable.data, %global.derive_class.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.33b7808bf372c3d58730520160cb2c15, i64 0, i64 0))
  %3 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %3, i32 0, i32 32)
  %4 = bitcast i8* %3 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %4, align 8
  %5 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e, i64 0, i64 0))
  %6 = getelementptr i8, i8* %3, i64 8
  %7 = bitcast i8* %6 to i32*
  %8 = load i32, i32* %7, align 4
  %9 = add i32 %8, 1
  store i32 %9, i32* %7, align 4
  %10 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.c1432ab71496ebb8b3e30bbcf37605e7, i64 0, i64 0), i32 %9) #1
  %11 = getelementptr i8, i8* %3, i64 16
  %12 = bitcast i8* %11 to i8**
  store i8* %0, i8** %12, align 8
  %13 = getelementptr i8, i8* %3, i64 24
  %14 = bitcast i8* %13 to void (i8*)**
  store void (i8*)* @global.derive_class.destroy, void (i8*)** %14, align 8
  %15 = load %global.derive_class.vtable.type*, %global.derive_class.vtable.type** %1, align 8
  %16 = getelementptr %global.derive_class.vtable.type, %global.derive_class.vtable.type* %15, i64 0, i32 2
  %17 = load void (i8*)*, void (i8*)** %16, align 8
  tail call void %17(i8* %0)
  tail call void @global.counter.release_shared(i8* %3)
  ret i32 0
}

define void @global.echo(i8* nocapture readonly %dc) local_unnamed_addr {
entry:
  %0 = getelementptr i8, i8* %dc, i64 16
  %1 = bitcast i8* %0 to i8**
  %2 = load i8*, i8** %1, align 8
  %3 = bitcast i8* %2 to %global.derive_class.vtable.type**
  %4 = load %global.derive_class.vtable.type*, %global.derive_class.vtable.type** %3, align 8
  %5 = getelementptr %global.derive_class.vtable.type, %global.derive_class.vtable.type* %4, i64 0, i32 2
  %6 = load void (i8*)*, void (i8*)** %5, align 8
  tail call void %6(i8* %2)
  ret void
}

define i8* @global.base_class.create() {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.base_class.vtable.type**
  store %global.base_class.vtable.type* @global.base_class.vtable.data, %global.base_class.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.726bd3560bd4c136648f7760895d8d62, i64 0, i64 0))
  ret i8* %0
}

; Function Attrs: nofree nounwind
define void @global.base_class.destroy(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.362aeeddb3d01da539cb6755bde46953, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @global.base_class.echo(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([13 x i8], [13 x i8]* @string.9bcbb503bda6c8ad83f772846b706f08, i64 0, i64 0))
  ret void
}

define i8* @global.derive_class.create() {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.derive_class.vtable.type**
  store %global.derive_class.vtable.type* @global.derive_class.vtable.data, %global.derive_class.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.33b7808bf372c3d58730520160cb2c15, i64 0, i64 0))
  ret i8* %0
}

; Function Attrs: nofree nounwind
define void @global.derive_class.destroy(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.362aeeddb3d01da539cb6755bde46953, i64 0, i64 0)) #1
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.ef25b0542457581e67c27a0dddb7bda5, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @global.derive_class.echo(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([13 x i8], [13 x i8]* @string.9bcbb503bda6c8ad83f772846b706f08, i64 0, i64 0)) #1
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.895758554639f423e017c6610cbf460b, i64 0, i64 0))
  ret void
}

attributes #0 = { nofree nounwind }
attributes #1 = { nounwind }
