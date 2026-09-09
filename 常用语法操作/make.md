# Go make 用法（刷算法速查）

> `make` 只用于构造 **map、slice、channel** 这三种**引用类型**（内置的、需要底层数据结构初始化的类型）。它是 Go 特有的初始化函数。

---

## 一、make 是什么

`make` 用于**创建并初始化**三种内置引用类型：

```go
make(map[K]V)      // 创建 map
make([]T, len, cap)  // 创建 slice
make(chan T, size)   // 创建 channel
```

> **关键区别**：`make` 返回的是**已初始化、可直接用**的值。
> 而 `var m map[string]int` 只是声明了个 nil map，**不能直接写入**。

---

## 二、三类的 make 用法

### 1. map
```go
m := make(map[string]int)            // 空的，可写
m["a"] = 1                           // ✅ 没问题

// 预先分配容量（可选，减少扩容，更快）
m2 := make(map[string]int, 100)      // 预分配 100 个槽位
```

### 2. slice（最常用）
```go
// 长度 0，容量 0
s := make([]int, 0)

// 长度 3，容量 3，元素是零值 [0 0 0]
s2 := make([]int, 3)

// 长度 3，容量 5（预留空间，避免多次扩容）
s3 := make([]int, 3, 5)
```

### 3. channel（并发用，刷题基本不碰）
```go
ch := make(chan int)          // 无缓冲
ch2 := make(chan int, 10)     // 缓冲 10
```

---

## 三、slice 的 len 和 cap（重点）

`make([]T, len, cap)` 的**两个参数**最容易混：

| 参数 | 含义 | 结果 |
|------|------|------|
| `len` | 当前长度 | 你能 `s[i]` 访问的范围 |
| `cap` | 容量 | 底层数组大小，决定何时扩容 |

```go
s := make([]int, 3, 5)
len(s)   // 3
cap(s)   // 5

s[0] = 1    // ✅ 索引 0~2 可用
s[3] = 1    // ❌ panic! 超过 len（虽然 cap 够，但 length 不够）

s = append(s, 9)      // 追加第 4 个，len 变 4
s = append(s, 9)      // 第 5 个，len 变 5（还没扩容）
s = append(s, 9)      // 超过 cap，会重新分配更大数组（深层拷贝）
```

> **记忆**：`len` 是你现在能用的；`cap` 是底层预留的总空间。刷题时常用 `make([]T, 0, n)` 预留够容量，避免 append 反复扩容。

---

## 四、标准刷题姿势

### 1. 预分配结果切片（高频）
```go
// 已知结果大小，先分配满，用下标写
res := make([]int, n)
for i := range res {
    res[i] = ...      // 直接用下标赋值
}
```

### 2. 动态追加，但预留容量
```go
// 用 make 预留容量，再 append，避免多次扩容
res := make([]int, 0, n)
for i := 0; i < n; i++ {
    res = append(res, ...)
}
```

### 3. map 预分配大小
```go
// 提前知道元素个数时，map 也预留容量
m := make(map[int]int, len(s))
for _, v := range s {
    m[v] = m[v] + 1
}
```

---

## 五、make vs new vs 字面量（易混）

| 写法 | 返回 | 适用 | 能否直接用 |
|------|------|------|-----------|
| `make([]int, 3)` | 已初始化的 slice | slice/map/channel | ✅ |
| `new([]int)` | **指向 nil 切片 的指针** | 任意类型 | ❌ 得再赋值 |
| `[]int{1,2,3}` | 字面量 slice | 已知内容 | ✅ |
| `map[string]int{}` | 字面量 map | 已知键值 | ✅ |

```go
// new 的坑：拿到的 slice 还是 nil
s := new([]int)
*s = append(*s, 1)      // 要先解引用再赋值

// 直接写会 panic 的例子
var m map[string]int    // nil map
m["a"] = 1              // ❌ panic: assignment to entry in nil map
```

> **结论**：刷题时 map 和 slice 统一用 **`make` 或用字面量**，**别用 `var m map[..]` 去直接写**，也别用 `new`。

---

## 六、make 的其它注意事项

### 1. 不用 make 的写法
```go
// slice 空值直接可用，不用 make
var s []int
s = append(s, 1)      // ✅ 空 nil 切片 append 也行

// map 必须初始化才能写
// var m map[string]int  ← 这样写，m["a"]=1 会 panic
```

### 2. make 只能用于这三种类型
```go
make(int)        // ❌ 编译错误，int 不是引用类型
make(MyStruct)   // ❌ 编译错误
make([]int, 3)   // ✅
make(map[int]int) // ✅
make(chan int)   // ✅
```

### 3. 预分配容量的作用
- 大量 append 前用 `make([]T, 0, n)`，能避免底层数组反复扩容拷贝，**提速明显**。
- 大量写入前用 `make(map[K]V, n)`，减少哈希表扩容。
- 刷题时如果数据量能估算，预分配是好习惯。

---

## 七、快速判断用哪个

| 想干嘛 | 用哪个 |
|--------|--------|
| 建一个空的、能直接写 key 的 map | `make(map[K]V)` 或 `map[K]V{}` |
| 建一个 index 可直接用的数组 | `make([]T, n)` |
| 从空慢慢 append 的列表 | `make([]T, 0, n)` 或 `[]T{}` |
| 已知内容的列表 | `[]T{a, b, c}` 字面量 |
| 建一个 channel | `make(chan T)` |

> **一句话**：`make` = 给「引用类型」分配并初始化底层结构，让 map/slice/channel 变得**直接可用**。