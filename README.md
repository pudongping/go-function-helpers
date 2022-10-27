# go-function-helpers

本项目收录一些高频、实用的 Go 语言助手函数，函数名称尽可能的和 PHP 中同功能函数一致，避免重复造轮子，提高开发效率。  
由于 Go 语言为强类型语言，因此本项目中的助手函数可能只会提供一种数据类型的作为参考，
要想兼容其他数据类型，请找到符合需求的助手函数之后，复制一份，然后自行调整为你所期望的数据类型。  
**因此，本项目更像是一个资料库，而不像是一个插件库。**

## 实用的工具推荐

- [JSON-to-Go](https://mholt.github.io/json-to-go/) —— 在线将 json 结构转换为 Go 类型定义
- [php2golang](https://www.php2golang.com/) —— PHP 函数到 Go 语言的转化

## 可用方法

### [数组 & 切片](#class-Arrayx)

a | b                                        | c
--- |------------------------------------------| --- 
[InArray](./arrayx#method-InArray) | [ArrayChunk](./arrayx#method-ArrayChunk) | [ArrayUnique](./arrayx#method-ArrayUnique)

### [字符串](#class-Strx)

a | b | c
--- | --- | --- 
[StrRandom](#method-StrRandom) | |

### [类型转换](#class-Convert)

a | b | c
--- | --- | --- 
[Struct2Map](#method-Struct2Map) | [Map2Struct](#method-Map2Struct) |


### [其他](#class-Other)
