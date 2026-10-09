# LeetCode 刷题记录（Go）

用 Go 语言刷 LeetCode 的个人题解仓库。

## 项目结构

```
.
├── common/                  # 共享数据结构定义
│   ├── listnode.go          # ListNode（链表节点）
│   ├── treenode.go          # TreeNode（二叉树节点）
│   └── node.go              # Node（带 Random 指针的链表节点）
├── 1_两数之和/              # 每题一个目录：题号_题目名
│   └── main.go              # 题解代码
├── 2_两数相加/
│   └── main.go
└── ...
```

## 运行方式

每个题目目录下是一个独立的 `main.go`，可以直接运行：

```bash
go run ./2_两数相加/main.go
```

或者进入对应目录：

```bash
cd 2_两数相加
go run .
```

## 公共结构体

`common` 包定义了刷题常用的数据结构，题解中通过 `leetcode/common` 引用：

```go
import "leetcode/common"

// 链表节点
type ListNode struct {
    Val  int
    Next *ListNode
}

// 二叉树节点
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 带随机指针的链表节点
type Node struct {
    Val    int
    Next   *Node
    Random *Node
}
```

## 环境

- Go 1.25+
- 模块名：`leetcode`

## 进度

目前收录约 40 道题，主要覆盖：

- **链表**：反转链表、环形链表、排序链表、合并 K 个升序链表、K 个一组翻转链表等
- **二叉树**：中序遍历、验证 BST、前序+中序构造、最近公共祖先、翻转二叉树等
- **二分/搜索**：搜索旋转排序数组、搜索二维矩阵、找左右边界等
- **动态规划 / 贪心**：括号匹配、有效括号字符串、子数组等
- **周赛**：`00_周赛` 目录
