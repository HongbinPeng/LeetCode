# Go map 常用语法（刷算法速查）

> 刷题只需要掌握这一份就够。核心操作 + 高频套路 + 与其它语言的差异。

---

## 一、最常用核心操作

```go
// 创建
m := map[string]int{}          // 空 map，可写
m2 := make(map[string]int)     // 同上，常用
m3 := map[string]int{"a": 1}   // 带初始值

// 增 / 改
m["a"] = 1
m["b"] = 2

// 查（取不到时返回零值）
v := m["c"]                   // v = 0，键不存在时是 int 零值

// 判断键是否存在（重点！多返回值）
v, ok := m["c"]
if ok {                       // ok 为 true 表示存在
    fmt.Println("存在", v)
} else {
    fmt.Println("不存在")      // v 是零值
}

// 删
delete(m, "a")

// 长度
len(m)
```

> **`v, ok := m[k]` 这个「多返回值」是 Go 刷题里最常用的** —— 判断键在不在，比如统计频率、记集合。

---

## 二、刷题高频写法

### 1. 统计频率（counter / 词频）
```go
count := map[string]int{}
for _, s := range strs {
    count[s]++            // 不存在时自动从 0 +1，超好用
}
```

### 2. 集合 Set（Go 没有内置 Set，用 map 模拟）
```go
// 用 struct{} 占位，更省内存
seen := map[int]struct{}{}
seen[1] = struct{}{}

if _, ok := seen[1]; ok {
    fmt.Println("在集合里")
}
```

### 3. 记录最早出现位置 / 下标
```go
pos := map[string]int{}
for i, s := range nums {
    if j, ok := pos[s]; ok {
        // 说明之前出现过
    }
    pos[s] = i        // 记录最新/最早下标
}
```

### 4. 父节点 / 指针映射（刷树题很常用）
```go
// 如：二叉树指向父节点
parent := map[*common.TreeNode]*common.TreeNode{}
parent[root] = nil
parent[node.Left] = node
```

---

## 三、遍历（重要：无序！）

```go
for k, v := range m {
    fmt.Println(k, v)
}

// 只要键 / 只要值
for k := range m { ... }
for _, v := range m { ... }
```

⚠️ **关键坑**：Go 的 map **遍历顺序是随机的、不保证**。如果需要有序输出，得先把 key 取出来排序。

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

---

## 四、Go 没有的（踩坑提醒）

| 你想要的功能 | Go 情况 |
|------------|---------|
| 像 Python 的 `dict.get(k, default)` | 没有，用 `if v,ok := m[k]; ok` 替代 |
| 像 Java 的 `map.putIfAbsent` | 没有，自己判断 |
| 像 Java 的 `map.containsKey` | 用 `_, ok := m[k]` |
| 有序 map（TreeMap） | 没有，用 `sort` + map |
| `map` 传给函数是引用还是拷贝 | 传的是**引用**（浅拷贝），函数内改 map 会影响外面的 |

---

## 五、刷题参考模板

```go
// 频率统计 + 存在判断 综合
func solution(nums []int) int {
    m := map[int]int{}
    for _, n := range nums {
        m[n]++
    }
    // 判断 / 返回值
    if v, ok := m[0]; ok {
        return v
    }
    return -1
}
```

---

## 与其它语言对照

| 需求 | Python | Java | Go |
|------|--------|------|-----|
| 哈希表 | `dict` | `HashMap` | 内置 `map` |
| 有序 map | `dict`（有顺序） | `TreeMap` | 无内置，用 `sort` + map |
| 集合 | `set` | `HashSet` | `map[type]struct{}` |
| 判断存在 | `k in d` | `containsKey` | `_, ok := m[k]` |
| 取值带默认 | `d.get(k, d)` | `getOrDefault` | `if v, ok := m[k]; ok` |