- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)

1. group.Go 内で error を返すとerrgroup がそれを検知
2. 関連するコンテキストがキャンセルされる
    * errgroup.WithContext によって生成されたコンテキストは最初のエラー発生時点でキャンセルされる。
    * ctx.Done() チャネルがすべての関連ゴルーチンに通知される。
3. ゴルーチンが ctx.Done() を確認することでキャンセルを検知

```mermaid

graph TD
    T1_STEP1["Task 1"]
    T2_STEP1["Task 2"]
    T3_STEP1["Task 3"]
    T1_STEP2["Task 1 (Finished)"]
    T2_STEP2["Task 2 (Error)"]
    T3_STEP2["Task 3 (Cancelled)"]
    E_STEP1["ErrGroup"]
    E_STEP2["Group Wait"]
    E_STEP3["Error Propagation"]

    M["Main Context"]
    M --> E_STEP1
    E_STEP1 --> T1_STEP1 --> E_STEP2 --> T1_STEP2
    E_STEP1 --> T2_STEP1 --> E_STEP2 --> T2_STEP2 --> E_STEP3 --> T3_STEP2
    E_STEP1 --> T3_STEP1 --> E_STEP2
```
