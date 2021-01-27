; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

@global.a = local_unnamed_addr global i32 0
@global.b = local_unnamed_addr global i32 1
@global.c = local_unnamed_addr global i32 2
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

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
  br i1 %5, label %10, label %exit

6:                                                ; preds = %entry
  %7 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([7 x i8], [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268, i64 0, i64 0))
  %.pr = load i32, i32* @global.a, align 4
  br label %3

exit:                                             ; preds = %10, %3
  %8 = phi i32 [ %.pre, %10 ], [ %4, %3 ]
  %9 = icmp eq i32 %8, 1
  %.0 = zext i1 %9 to i32
  ret i32 %.0

10:                                               ; preds = %3
  %11 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([7 x i8], [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc, i64 0, i64 0))
  %.pre = load i32, i32* @global.a, align 4
  br label %exit
}

attributes #0 = { nofree nounwind }
