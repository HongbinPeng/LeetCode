# Go 排序 sort 常用语法（刷算法速查）

> Go 没有 Java 的 `Collections.sort`，内置 **`sort` 包**。刷题记住：**基本类型排序 + 切片自定义排序 + 降序** 就够。

---

## 一、基本类型排序

```go
import "sort"

// int 切片
nums := []int{3, 1, 2}
sort.Ints(nums)                                    // [1 2 3]  升序
sort.Sort(sort.Reverse(sort.IntSlice(nums)))       // [3 2 1]  降序

// float64
f := []float64{3.1, 1.2}
sort.Float64s(f)

// string
s := []string{"b", "a", "c"}
sort.Strings(s)
```

---

## 二、自定义排序（刷题重点）

### 1. 按结构体字段排序
```go
type Person struct {
    Name string
    Age  int
}

ps := []Person{{"a", 30}, {"b", 20}, {"c", 25}}
sort.Slice(ps, func(i, j int) bool {
    return ps[i].Age < ps[j].Age      // 升序；> 为降序
})
```

### 2. 多条件排序（先年龄，再姓名）
```go
sort.Slice(ps, func(i, j int) bool {
    if ps[i].Age != ps[j].Age {
        return ps[i].Age < ps[j].Age        // 主：年龄升序
    }
    return ps[i].Name < ps[j].Name          // 次：姓名升序
})
```

### 3. 自定义类型排序（实现 Less 接口）
```go
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

sort.Sort(ByAge(ps))
```

---

## 三、常用组合（刷题高频）

### 1. 找第 K 大 / 取前 K
```go
nums := []int{5, 2, 8, 1, 9}
sort.Sort(sort.Reverse(sort.IntSlice(nums)))   // 降序 [9 8 5 2 1]
top3 := nums[:3]                                // 前 3 大 [9 8 5]
```

### 2. 对 map 的 key 排序（有序输出）
```go
m := map[string]int{"a": 1, "c": 3, "b": 2}
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

### 3. 对 map 按 value 排序（频率榜）
```go
type kv struct {
    k string
    v int
}
slice := make([]kv, 0, len(m))
for k, v := range m {
    slice = append(slice, kv{k, v})
}
sort.Slice(slice, func(i, j int) bool {
    return slice[i].v > slice[j].v     // 按值降序
})
```

---

## 四、需要自己写排序的场景

**时间 O(N log N)** 的 `sort` 不是万能的。当题目要求 O(N log N) 以下、或不想破坏原数组时，可用**计数排序**等：

### 1. 计数排序（值域小，如字母频率）
```go
// 统计 26 个小写字母频率后再按序输出
count := [26]int{}
for _, c := range s {
    count[c-'a']++
}
```

### 2. 桶排序（值域有限）
```go
// 数字范围小（如成绩 0~100）时用桶
```

---

## 五、易错陷阱

| 场景 | 坑 |
|------|-----|
| `sort.Slice` 的 `less` 是不稳定排序 | 相等元素相对顺序**不保证**，需要稳定用 `sort.SliceStable` |
| 降序写法 | `>` 是降序，`<` 是升序 | 
| 修改原切片 | `sort` 会**原地修改**，想保留原数据先 `copy` |
| 对 `map` 排序 | map 本身无序，**必须先把 key/value 取到切片再排** |

---

## 六、刷题参考模板

```go
import "sort"

// 按频率从高到低输出元素（top K 常用）
type Pair struct {
    Val int
    Frq int
}

func sortByFreq(nums []int) []Pair {
    cnt := map[int]int{}
    for _, n := range nums {
        cnt[n]++
    }
    pairs := make([]Pair, 0, len(cnt))
    for v, f := range cnt {
        pairs = append(pairs, Pair{v, f})
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].Frq > pairs[j].Frq
    })
    return pairs
}
```

---

## 内置函数速查（`sort` 包）

| 函数 | 作用 |
|------|------|
| `sort.Ints(s)` / `Float64s` / `Strings` | 基本类型升序 |
| `sort.Reverse(x)` | 反向排序器，配合 `sort.Sort` |
| `sort.Slice(s, less)` | 切片自定义排序（最常用） |
| `sort.SliceStable(s, less)` | 稳定版本，保序 |
| `sort.Search(n, f)` | 二分查找 |

### 二分查找（排序数组常用）
```go
nums := []int{1, 3, 5, 7}
i := sort.Search(len(nums), func(i int) bool {
    return nums[i] >= 5     // 找第一个 >=5 的位置
})                          // i = 2
```