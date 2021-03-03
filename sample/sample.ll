; ModuleID = 'sample.c'
source_filename = "sample.c"
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

; Function Attrs: noinline nounwind optnone uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %2 = alloca float, align 4
  %3 = alloca float, align 4
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  %6 = alloca i32, align 4
  store i32 0, i32* %1, align 4
  store float 0x3FB99999A0000000, float* %2, align 4
  %7 = load float, float* %2, align 4
  %8 = fsub float -0.000000e+00, %7
  store float %8, float* %3, align 4
  store i32 1, i32* %4, align 4
  %9 = load i32, i32* %4, align 4
  %10 = sub nsw i32 0, %9
  store i32 %10, i32* %5, align 4
  %11 = load i32, i32* %4, align 4
  store i32 %11, i32* %6, align 4
  ret i32 0
}

; Function Attrs: noinline nounwind optnone uwtable
define zeroext i1 @test(i32*) #0 {
  %2 = alloca i1, align 1
  %3 = alloca i32*, align 8
  store i32* %0, i32** %3, align 8
  %4 = load i32*, i32** %3, align 8
  %5 = icmp eq i32* %4, null
  br i1 %5, label %6, label %7

; <label>:6:                                      ; preds = %1
  store i1 true, i1* %2, align 1
  br label %8

; <label>:7:                                      ; preds = %1
  store i1 false, i1* %2, align 1
  br label %8

; <label>:8:                                      ; preds = %7, %6
  %9 = load i1, i1* %2, align 1
  ret i1 %9
}

attributes #0 = { noinline nounwind optnone uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.module.flags = !{!0}
!llvm.ident = !{!1}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{!"clang version 6.0.0-1ubuntu2 (tags/RELEASE_600/final)"}
