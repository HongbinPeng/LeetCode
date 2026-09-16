# Go 字符串 string 常用语法（刷算法速查）

> Go 的 `string` 是**只读字节序列**，底层是 UTF-8。刷题常跟 **ASCII**、**中文字符**、**子串**、**拼接**打交道。核心坑：**下标是字节，不是字符**。

---

## 一、常用核心操作

```go
s := "hello"
len(s)                  // 长度 = 字节数（不是字符数！）

// 拼接
s2 := s + " world"
s3 := fmt.Sprintf("%s-%d", s, 42)

// 下标取字节（返回 byte，不是字符）
b := s[0]               // byte = 'h'
s[0] = 'H'              // ❌ 编译错误！string 是只读的

// 取子串
sub := s[1:4]           // "ell"，左闭右开，按字节

// 判断相等
s == "hello"
strings.Contains(s, "ell")

// 大小写
strings.ToUpper(s)
strings.ToLower(s)

// 分割 / 拼接
parts := strings.Split("a,b,c", ",")   // []string{"a","b","c"}
joined := strings.Join(parts, "-")     // "a-b-c"

// 查找 / 替换
strings.Index(s, "ll")   // 第一个位置，找不到返回 -1
strings.Replace(s, "l", "L", -1)  // -1 表示全部
```

---

## 二、字符与字节（重点坑）

### 1. 遍历 ASCII 字符串（用下标）
```go
s := "abc"
for i := 0; i < len(s); i++ {
    b := s[i]            // byte，直接能用
}
```

### 2. 遍历可能含中文的字符串（用 for range）
```go
s := "go语言"
for i, c := range s {
    // c 是 rune（一个 Unicode 字符），i 是字节偏移
    fmt.Println(i, string(c))
}
```

### 3. 统计字符出现次数
> 因为 `for range` 按字符走，英文/中文都能数对。

```go
count := map[rune]int{}
for _, c := range s {
    count[c]++           // c 是 rune（int32）
}
```

### 4. 判断字母/数字
```go
import "unicode"
unicode.IsLetter('a')
unicode.IsDigit('5')
unicode.IsLetter(rune(b))   // 处理 byte 时转 rune
```

---

## 三、string ↔ byte / rune 转化

```go
// string → []byte（字节）
bs := []byte(s)
s = string(bs)

// string → []rune（字符，可安全处理中文）
rs := []rune(s)          // 中文字不再乱码/无法取整
s = string(rs)

// 单字符
c := 'a'                 // rune(int32 类型)
s = string(c)
```

> **刷题关键**：**遇到中文/多字节要按字符处理**时，用 `[]rune(s)` 转换为字符切片，`len(rs)` 才是字符个数。

---

## 四、高频写法

### 1. 字符串反转
```go
func reverse(s string) string {
    rs := []rune(s)
    for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
        rs[i], rs[j] = rs[j], rs[i]
    }
    return string(rs)
}
```

### 2. 判断回文
```go
func isPalindrome(s string) bool {
    rs := []rune(s)
    for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
        if rs[i] != rs[j] {
            return false
        }
    }
    return true
}
```

### 3. 字符频率（anagram 判断常用）
```go
func charCount(s string) map[rune]int {
    m := map[rune]int{}
    for _, c := range s {
        m[c]++
    }
    return m
}
```

---

## 五、易错陷阱

| 场景 | 坑 |
|------|-----|
| `len(s)` 是字节数 | 字符串含中文时，`len` ≠ 字符数 |
| `s[i]` 是 byte | 中文会被拆成 3 个 byte，下标取到乱码 |
| `string` 只读 | 不能 `s[i] = ...`，需转 `[]byte`/`[]rune` |
| `s[1:3]` 按字节切 | 中文会切破，可能乱码 |
| `+` 拼接效率 | 大循环拼接很慢，用 `strings.Builder` |

### 高效拼接（循环里）
```go
var b strings.Builder
for _, s := range strs {
    b.WriteString(s)
}
result := b.String()
```

---

## 六、刷题参考模板

```go
import "strings"

// 统计不区分大小写的字符频率
func solve(s string) map[byte]int {
    m := map[byte]int{}
    for i := 0; i < len(s); i++ {
        c := s[i]
        if 'A' <= c && c <= 'Z' {
            c += 32              // 转小写
        }
        m[c]++
    }
    return m
}
```

---

## 常见内置函数速查（`strings` 包）

| 函数 | 作用 |
|------|------|
| `len(s)` | 字节长度 |
| `strings.Contains(s, sub)` | 是否包含子串 |
| `strings.Split(s, sep)` | 按分隔符切割 |
| `strings.Join(slice, sep)` | 用分隔符拼接 |
| `strings.Index(s, sub)` | 子串位置，无返回 -1 |
| `strings.Replace(s, old, new, n)` | 替换，n=-1 全部 |
| `strings.HasPrefix/HasSuffix` | 前后缀 |
| `strings.ToUpper/ToLower` | 大小写 |
| `strings.TrimSpace(s)` | 去首尾空白 |
| `strings.Builder` | 高效拼接 |