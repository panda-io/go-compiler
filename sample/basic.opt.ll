; ModuleID = './sample/basic.ll'
source_filename = "./sample/basic.ll"

@global.a = local_unnamed_addr global i32 0
@global.b = local_unnamed_addr global i32 1
@global.c = local_unnamed_addr global i32 2

; Function Attrs: norecurse nounwind readnone
define i32 @main() local_unnamed_addr #0 {
entry:
  ret i32 0
}

attributes #0 = { norecurse nounwind readnone }
