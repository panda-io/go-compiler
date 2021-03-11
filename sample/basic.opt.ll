; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

%global.counter.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)*, void (i8*)* }
%global.base.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)* }
%global.derive.vtable.type = type { i8* ()*, void (i8*)*, void (i8*)* }

@global.counter.vtable.data = global %global.counter.vtable.type { i8* ()* @global.counter.create, void (i8*)* @global.counter.destroy, void (i8*)* @global.counter.retain_shared, void (i8*)* @global.counter.release_shared, void (i8*)* @global.counter.retain_weak, void (i8*)* @global.counter.release_weak }
@global.base.vtable.data = global %global.base.vtable.type { i8* ()* @global.base.create, void (i8*)* @global.base.destroy, void (i8*)* @global.base.do }
@global.derive.vtable.data = global %global.derive.vtable.type { i8* ()* @global.derive.create, void (i8*)* @global.derive.destroy, void (i8*)* @global.derive.do }
@string.5bdaebb122965539cdd6ce77f212b65e = constant [15 x i8] c"create counter\00"
@string.f8f86b3941cca26e8c147322b9a8309f = constant [16 x i8] c"destroy counter\00"
@string.cf85dc053c0475520502efb2ba3c77a9 = constant [14 x i8] c"retain shared\00"
@string.3b2e33db0bf2dc0aef4015e91829dab5 = constant [15 x i8] c"release shared\00"
@string.c52993c907c8f30bf3854dc0e21a7eca = constant [19 x i8] c"shared count: %d \0A\00"
@string.21c67ac9191c65481dbab306227b4840 = constant [17 x i8] c"free object %p \0A\00"
@string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0 = constant [18 x i8] c"free counter %p \0A\00"
@string.5927c4441dce664e4b461e529f933750 = constant [12 x i8] c"retain weak\00"
@string.5662737e1a39fc068ead71add358dfd3 = constant [13 x i8] c"release weak\00"
@string.b6feae5df5d6172ffcb2a6bcd4d5c478 = constant [17 x i8] c"weak count: %d \0A\00"
@string.726bd3560bd4c136648f7760895d8d62 = constant [18 x i8] c"base construction\00"
@string.362aeeddb3d01da539cb6755bde46953 = constant [17 x i8] c"base destruction\00"
@string.b58017d6f5ff15cba2431d7ec3967243 = constant [21 x i8] c"do something in base\00"
@string.33b7808bf372c3d58730520160cb2c15 = constant [20 x i8] c"derive construction\00"
@string.ef25b0542457581e67c27a0dddb7bda5 = constant [19 x i8] c"derive destruction\00"
@string.7cf7144b5f7d7f8893615fa04d42f3f7 = constant [23 x i8] c"do something in derive\00"

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
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.cf85dc053c0475520502efb2ba3c77a9, i64 0, i64 0))
  %1 = getelementptr i8, i8* %this, i64 8
  %2 = bitcast i8* %1 to i32*
  %3 = load i32, i32* %2, align 4
  %4 = add i32 %3, 1
  store i32 %4, i32* %2, align 4
  ret void
}

define void @global.counter.release_shared(i8* %this) {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.3b2e33db0bf2dc0aef4015e91829dab5, i64 0, i64 0))
  %1 = icmp eq i8* %this, null
  br i1 %1, label %exit, label %2

exit:                                             ; preds = %10, %24, %2, %entry
  ret void

2:                                                ; preds = %entry
  %3 = getelementptr i8, i8* %this, i64 8
  %4 = bitcast i8* %3 to i32*
  %5 = load i32, i32* %4, align 4
  %6 = add i32 %5, -1
  store i32 %6, i32* %4, align 4
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.c52993c907c8f30bf3854dc0e21a7eca, i64 0, i64 0), i32 %6)
  %8 = load i32, i32* %4, align 4
  %9 = icmp eq i32 %8, 0
  br i1 %9, label %10, label %exit

10:                                               ; preds = %2
  %11 = getelementptr i8, i8* %this, i64 24
  %12 = bitcast i8* %11 to void (i8*)**
  %13 = load void (i8*)*, void (i8*)** %12, align 8
  %14 = getelementptr i8, i8* %this, i64 16
  %15 = bitcast i8* %14 to i8**
  %16 = load i8*, i8** %15, align 8
  tail call void %13(i8* %16)
  %17 = load i8*, i8** %15, align 8
  %18 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.21c67ac9191c65481dbab306227b4840, i64 0, i64 0), i8* %17)
  %19 = load i8*, i8** %15, align 8
  tail call void @free(i8* %19)
  store i8* null, i8** %15, align 8
  %20 = getelementptr i8, i8* %this, i64 12
  %21 = bitcast i8* %20 to i32*
  %22 = load i32, i32* %21, align 4
  %23 = icmp eq i32 %22, 0
  br i1 %23, label %24, label %exit

24:                                               ; preds = %10
  %25 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0, i64 0, i64 0), i8* nonnull %this)
  tail call void @free(i8* nonnull %this)
  br label %exit
}

; Function Attrs: nofree nounwind
define void @global.counter.retain_weak(i8* nocapture %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.5927c4441dce664e4b461e529f933750, i64 0, i64 0))
  %1 = getelementptr i8, i8* %this, i64 12
  %2 = bitcast i8* %1 to i32*
  %3 = load i32, i32* %2, align 4
  %4 = add i32 %3, 1
  store i32 %4, i32* %2, align 4
  ret void
}

; Function Attrs: nounwind
define void @global.counter.release_weak(i8* %this) #1 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([13 x i8], [13 x i8]* @string.5662737e1a39fc068ead71add358dfd3, i64 0, i64 0))
  %1 = icmp eq i8* %this, null
  br i1 %1, label %exit, label %2

exit:                                             ; preds = %2, %14, %entry
  ret void

2:                                                ; preds = %entry
  %3 = getelementptr i8, i8* %this, i64 12
  %4 = bitcast i8* %3 to i32*
  %5 = load i32, i32* %4, align 4
  %6 = add i32 %5, -1
  store i32 %6, i32* %4, align 4
  %7 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.b6feae5df5d6172ffcb2a6bcd4d5c478, i64 0, i64 0), i32 %6)
  %8 = getelementptr i8, i8* %this, i64 8
  %9 = bitcast i8* %8 to i32*
  %10 = load i32, i32* %9, align 4
  %11 = load i32, i32* %4, align 4
  %12 = or i32 %11, %10
  %13 = icmp eq i32 %12, 0
  br i1 %13, label %14, label %exit

14:                                               ; preds = %2
  %15 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.4fc1bf1a9ddd2be568f08ffc8ed6b9f0, i64 0, i64 0), i8* nonnull %this)
  tail call void @free(i8* nonnull %this)
  br label %exit
}

define i32 @main() local_unnamed_addr {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.derive.vtable.type**
  store %global.derive.vtable.type* @global.derive.vtable.data, %global.derive.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.33b7808bf372c3d58730520160cb2c15, i64 0, i64 0))
  %3 = tail call dereferenceable_or_null(32) i8* @malloc(i32 32)
  tail call void @memset(i8* %3, i32 0, i32 32)
  %4 = bitcast i8* %3 to %global.counter.vtable.type**
  store %global.counter.vtable.type* @global.counter.vtable.data, %global.counter.vtable.type** %4, align 8
  %5 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bdaebb122965539cdd6ce77f212b65e, i64 0, i64 0))
  %6 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([14 x i8], [14 x i8]* @string.cf85dc053c0475520502efb2ba3c77a9, i64 0, i64 0)) #1
  %7 = getelementptr i8, i8* %3, i64 8
  %8 = bitcast i8* %7 to i32*
  %9 = load i32, i32* %8, align 4
  %10 = add i32 %9, 1
  store i32 %10, i32* %8, align 4
  %11 = getelementptr i8, i8* %3, i64 16
  %12 = bitcast i8* %11 to i8**
  store i8* %0, i8** %12, align 8
  %13 = getelementptr i8, i8* %3, i64 24
  %14 = bitcast i8* %13 to void (i8*)**
  store void (i8*)* @global.derive.destroy, void (i8*)** %14, align 8
  tail call void @global.counter.release_shared(i8* %3)
  ret i32 0
}

define i8* @global.base.create() {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.base.vtable.type**
  store %global.base.vtable.type* @global.base.vtable.data, %global.base.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([18 x i8], [18 x i8]* @string.726bd3560bd4c136648f7760895d8d62, i64 0, i64 0))
  ret i8* %0
}

; Function Attrs: nofree nounwind
define void @global.base.destroy(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.362aeeddb3d01da539cb6755bde46953, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @global.base.do(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([21 x i8], [21 x i8]* @string.b58017d6f5ff15cba2431d7ec3967243, i64 0, i64 0))
  ret void
}

define i8* @global.derive.create() {
entry:
  %0 = tail call dereferenceable_or_null(8) i8* @malloc(i32 8)
  tail call void @memset(i8* %0, i32 0, i32 8)
  %1 = bitcast i8* %0 to %global.derive.vtable.type**
  store %global.derive.vtable.type* @global.derive.vtable.data, %global.derive.vtable.type** %1, align 8
  %2 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([20 x i8], [20 x i8]* @string.33b7808bf372c3d58730520160cb2c15, i64 0, i64 0))
  ret i8* %0
}

; Function Attrs: nofree nounwind
define void @global.derive.destroy(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([17 x i8], [17 x i8]* @string.362aeeddb3d01da539cb6755bde46953, i64 0, i64 0)) #1
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([19 x i8], [19 x i8]* @string.ef25b0542457581e67c27a0dddb7bda5, i64 0, i64 0))
  ret void
}

; Function Attrs: nofree nounwind
define void @global.derive.do(i8* nocapture readnone %this) #0 {
entry:
  %0 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([23 x i8], [23 x i8]* @string.7cf7144b5f7d7f8893615fa04d42f3f7, i64 0, i64 0))
  ret void
}

attributes #0 = { nofree nounwind }
attributes #1 = { nounwind }
