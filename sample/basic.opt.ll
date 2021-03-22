; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

@string.5bddf146c13b387514280200e83cf08b = constant [15 x i8] c"value #1: %d \0A\00"
@string.aee0f3d368512408f6bd5274bf51a219 = constant [15 x i8] c"value #2: %d \0A\00"

; Function Attrs: nofree nounwind
declare i32 @printf(i8* nocapture readonly, ...) local_unnamed_addr #0

; Function Attrs: nofree nounwind
define void @global.print_number(i32* nocapture readonly %value) local_unnamed_addr #0 {
entry:
  %0 = load i32, i32* %value, align 4
  %1 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.5bddf146c13b387514280200e83cf08b, i64 0, i64 0), i32 %0)
  %2 = load i32, i32* %value, align 4
  %3 = tail call i32 (i8*, ...) @printf(i8* nonnull dereferenceable(1) getelementptr inbounds ([15 x i8], [15 x i8]* @string.aee0f3d368512408f6bd5274bf51a219, i64 0, i64 0), i32 %2)
  ret void
}

attributes #0 = { nofree nounwind }
