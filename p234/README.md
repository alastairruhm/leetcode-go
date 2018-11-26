# 234. Palindrome Linked List

Palindrome Linked List - LeetCode: https://leetcode.com/problems/palindrome-linked-list/description/

## description

Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true

Follow up:
**Could you do it in O(n) time and O(1) space?**

## Summary

方法一：遍历链表，取出其中的结点值拼成字符串，然后判断字符串是否是回文字符串

时间复杂度 O(n)
空间复杂度 O(n)

方法二：思想利用快慢指针法找到链表中点，然后一个将后半部分就地反转，分别再从头、中点遍历判断是否相等，该方法时间复杂度O(n)、空间复杂度O(1).

考虑偶数个节点的情况

```
slow
head   fast
1   ->  2  ->  2  ->  1  -> null

head   slow          fast
1   ->  2  ->  2  ->  1  -> null
```

考虑奇数个节点的情况

```
temp
fast
slow
head
1   ->  2  ->  1 -> null

temp
head   slow   fast
1   <-  2      1 -> null
```

```
slow
head   fast
1   ->  2  ->  3  ->  2  ->  1  ->  null 

head   slow          fast
1   ->  2  ->  3  ->  2  ->  1  ->  null 

head          slow                  fast
1   ->  2  ->  3  ->  2  ->  1  ->  null

head          slow         cur            fast
1   ->  2  ->  3  ->  null  2  ->  1  ->  null

head          slow temp      cur     q      fast
1   ->  2  ->  3   null  <-   2  ->  1  ->  null

head          slow           temp   cur(q)  fast
1   ->  2  ->  3   null  <-   2  ->  1  ->  null

head          slow                  temp    cur/q/fast
1   ->  2  ->  3   null  <-   2   <-  1  ->  null
```


```
temp
fast
slow
head
1   ->  2
```

```
temp
fast
slow
head
1   ->  null
```

## tips


