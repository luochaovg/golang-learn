## Golang中GC回收机制三色标记与混合写屏障
- https://github.com/luochaovg/golang/blob/main/5%E3%80%81Golang%E4%B8%89%E8%89%B2%E6%A0%87%E8%AE%B0%2B%E6%B7%B7%E5%90%88%E5%86%99%E5%B1%8F%E9%9A%9CGC%E6%A8%A1%E5%BC%8F%E5%85%A8%E5%88%86%E6%9E%90.md
- https://www.bilibili.com/video/BV1wz4y1y7Kd?p=7&spm_id_from=pageDriver

#### 1,GC
```text
自动释放
垃圾回收
三色标记法
内存管理
STW

三色标记法
强弱三色不变式
屏障机制


七、总结
​ 以上便是Golang的GC全部的标记-清除逻辑及场景演示全过程。
GoV1.3- 普通标记清除法，整体过程需要启动STW，效率极低。
GoV1.5- 三色标记法， 堆空间启动写屏障，栈空间不启动，全部扫描之后，需要重新扫描一次栈(需要STW)，效率普通
GoV1.8-三色标记法，混合写屏障机制， 栈空间不启动，堆空间启动。整个过程几乎不需要STW，效率较高。
```

