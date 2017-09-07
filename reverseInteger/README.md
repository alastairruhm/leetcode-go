# 7. Two Sum

[Reverse Integer - LeetCode](https://leetcode.com/problems/reverse-integer/description/)

## description

Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

## Summary


## tips

### golang 基础类型

bool         1      true/false,默认false, 不能把非0值当做true(不用数字代表true/false)
byte         1      uint8 别名
rune         4      int32别名。 代表一个unicode code point
int/unit            依赖运行的平台，32bit/64bit
int8/uint8   1     -128 ~ 127; 0 ~ 255
int16/uint16 2     -32768 ~ 32767; 0 ~ 65535
int32/uint32 4     -21亿 ~ 21亿， 0 ~ 42亿
int64/uint64 8

float32      4     精确到7位小数,相当于c的float
float64      8     精确到15位小数,相当于c的double

complex64    8
complex128   16

uintptr            足够保存指针的32位、64位整数,指针(可以存指针的整数型)
array              值类型,数组
struct             值类型,结构体
string             值类型,字符串类型，常用
slice              引用类型,切片
map                引用类型,字典
channel            引用类型,通道
interface          接口类型,接口
function           函数类型,函数