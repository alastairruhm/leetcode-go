# 232. Implement Queue using Stacks

https://leetcode.com/problems/implement-queue-using-stacks/

## description

Given a singly linked list, determine if it is a palindrome.Implement the following oper                                                                ations of a queue using stacks.

* push(x) -- Push element x to the back of queue.
* pop() -- Removes the element from in front of queue.
* peek() -- Get the front element.
* empty() -- Return whether the queue is empty.

Example:

```
MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
```

Notes:

* You must use only standard operations of a stack -`   `   ``          `   `   ``  `   - which means only push to top, peek/pop from top, size, and is empty operations are valid.
* Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
* You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

## tips

一个栈接受 push 操作的数据，当需要 Peek 和 Pop 时，将数据倒腾到另外一个栈，形成逆序

```
-> push 1

s1: 1
s2: 

-> push 2

s1: 2 | 1
s2: 

-> Peek()

Peek 把 s1 数据倒腾给 s2
s1: 
s2: 1 | 2  


-> Pop()

s2 栈的顶部是 1

s1: 
s2: 2  s2.Pop()
```