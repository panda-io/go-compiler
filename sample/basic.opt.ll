; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

@global.a = local_unnamed_addr global i32 0
@string.7ac4a7d2c8bf695dc8a116f6b70a8e1a = constant [9 x i8] c"b = %d \0A\00"
@string.cb091131e20d7842e7627e8736856b45 = constant [12 x i8] c"hello world\00"

; Function Attrs: nofree nounwind
declare i32 @puts(i8* nocapture readonly) local_unnamed_addr #0

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

; Function Attrs: nofree nounwind
define i32 @main() local_unnamed_addr #0 {
entry:
  %0 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([9 x i8], [9 x i8]* @string.7ac4a7d2c8bf695dc8a116f6b70a8e1a, i64 0, i64 0), i32 9)
  %1 = tail call i32 @puts(i8* nonnull dereferenceable(1) getelementptr inbounds ([12 x i8], [12 x i8]* @string.cb091131e20d7842e7627e8736856b45, i64 0, i64 0))
  ret i32 0
}

attributes #0 = { nofree nounwind }
