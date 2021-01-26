; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

@global.a = local_unnamed_addr global i32 0
@global.b = local_unnamed_addr global i32 1
@global.c = local_unnamed_addr global i32 2
@string.459521b87e7c4e2aa0de9b45c0a81268 = constant [7 x i8] c"a is 0\00"
@string.318ab47b7b6a7bd68c90f6696d16b2fc = constant [7 x i8] c"a is 1\00"

; Function Attrs: nofree nounwind
define i32 @main() local_unnamed_addr #0 {
entry:
  %0 = load i32, i32* @global.a, align 4
  %1 = add i32 %0, 1
  store i32 %1, i32* @global.a, align 4
  %2 = icmp eq i32 %1, 0
  br i1 %2, label %3, label %5

3:                                                ; preds = %entry
  %4 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([7 x i8], [7 x i8]* @string.459521b87e7c4e2aa0de9b45c0a81268, i64 0, i64 0))
  %.pr = load i32, i32* @global.a, align 4
  br label %5

5:                                                ; preds = %3, %entry
  %6 = phi i32 [ %.pr, %3 ], [ %1, %entry ]
  %7 = icmp eq i32 %6, 1
  br i1 %7, label %8, label %exit

8:                                                ; preds = %5
  %9 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([7 x i8], [7 x i8]* @string.318ab47b7b6a7bd68c90f6696d16b2fc, i64 0, i64 0))
  br label %exit

exit:                                             ; preds = %8, %5
  ret i32 0
}

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

attributes #0 = { nofree nounwind }
