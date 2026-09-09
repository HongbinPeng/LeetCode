# Go 切片 slice 常用语法（刷算法速查）

> Go 里数组用的少，**切片 slice** 几乎是「数组+动态列表」的统一替代。刷题必须掌握。

---

## 一、常用核心操作

```go
// 创建
s := []int{}                 // 空切片（**可写**，和 map 一样）
s2 := make([]int, 3)         // 长度 3，元素零值 [0 0 0]
s3 := make([]int, 3, 5)      // 长度 3，容量 5
s4 := []int{1, 2, 3}         // 字面量
s5 := make([]int, 0, 10)     // 长度 0，容量 10（先预留，避免扩容，高效）

// 增（append）
s = append(s, 4)
s = append(s, 5, 6)          // 追加多个

// 取长度 / 容量
len(s)
cap(s)

// 取值 / 下标索引
v := s[0]

// 切片（取子段，闭开区间 [i:j]）
sub := s[1:3]                // 索引 1 到 2，不含 3
```
> **注意**：`s[:j]`、`s[i:]`、`s[i:j]` 都是**左闭右开**区间，`[:len(s)-1]` 是**去掉最后一个元素**（模拟出栈）。

---

## 二、刷题高频写法

### 1. 用切片模拟栈（后进先出 LIFO）
```go
// 入栈 push
stack = append(stack, x)

// 看栈顶 peek
top := stack[len(stack)-1]

// 出栈 pop
x := stack[len(stack)-1]
stack = stack[:len(stack)-1]
```

### 2. 用切片模拟队列（先进先出 FIFO）
```go
// 队头
front := queue[0]

// 出队
x := queue[0]
queue = queue[1:]

// 入队
queue = append(queue, x)
```
> 每次 `queue[1:]` 会不断移动底层数组起点，适合简单场景；大量出队时可改用循环数组或"移动指针"优化。

### 3. 遍历
```go
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])         // 用索引
}

for i, v := range s {         // i 索引，v 值
    fmt.Println(i, v)
}

for _, v := range s { ... }   // 只要值
for i := range s { ... }      // 只要索引
```

### 4. 删除元素
```go
// 删除下标 i 的元素（保留原顺序）
s = append(s[:i], s[i+1:]...)

// 删除最后一个元素（出栈）
s = s[:len(s)-1]
```

### 5. 拷贝
```go
s2 := make([]int, len(s))
copy(s2, s)                   // 深拷贝，改 s2 不影响 s

// 浅拷贝（共享底层数组，危险！）
s3 := s
s3[0] = 99                    // 会同时改掉 s[0]
```

### 6. 排序切片
```go
import "sort"
sort.Ints(s)                          // int 切片升序
sort.Sort(sort.Reverse(sort.IntSlice(s)))  // int 切片降序
sort.Strings(strSlice)

// 自定义排序（结构体等）
sort.Slice(s, func(i, j int) bool {
    return s[i].Val < s[j].Val
})
```

---

## 三、易错陷阱

| 场景 | 坑 |
|------|-----|
| 浅拷贝底层共享 | `s2 := s` 只是引用，改一个影响另一个；要拷贝用 `copy` |
| `append` 扩容 | 超过 cap 会**重新分配数组**，之前取得子切片/引用会指向旧数组 |
| 空切片 nil vs 空 | `var s []int`（nil）和 `s := []int{}`（空）都能 append 用 |
| 切片越界 | `s[0]` 空切片会 panic；先判 `len(s) > 0` |
| 切片区间左闭右开 | `s[1:3]` 是第 1、2 个，不含第 3 个 |

---

## 四、刷题参考模板

```go
// 读取 + 遍历 + 处理 综合
func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    for i, n := range nums {
        if j, ok := seen[target-n]; ok {
            return []int{j, i}
        }
        seen[n] = i
    }
    return []int{}
}
```

---

## 五、数组 vs 切片

| | 数组 `[3]int` | 切片 `[]int` |
|---|---|---|
| 长度 | 固定，是类型一部分 | 动态，可扩展 |
| 函数传参 | 值拷贝 | 引用（共享底层） |
| 刷题用 | 几乎不用 | **几乎都用** |

> 一句话：刷题当你想用"数组/列表"，**默认用切片 `[]T`**。
