## 基础语法
#### 1. 定义变量
使用 var 定义变量, 或者使用 := 定义变量, 但是这种写法只能用在函数体内
#### 2. 定义常量
使用 const 来定义常量, 使用 const 还可以定义枚举, 注意 iota 的用法
#### 3. 运算符
goLang 的运算符和其他语言类似, 只有 & 变量存储地址 * 指针变量. 目前尚未接触过
#### 4. 条件语句
if else switch 和其他语言类似, 新增的 select 判断, 类似于 switch case . 只是 case 中是 io 操作
#### 5. 循环判断语句
基本和其他语言一样 for 循环
#### 6. 函数
语法: 
```$xslt
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
go 除了普通的函数外还有匿名函数并具有闭包功能, 可以类似其他语言的的元祖返回多个值
#### 7.方法
语法: 
```$xslt
func (variable_name variable_data_type) function_name() [return_type]{
   /* 函数体*/
}
```
一个方法就是一个包含了接受者的函数

#### 8.变量作用域
1. 函数内定义的变量称为局部变量
2. 函数外定义的变量称为全局变量
3. 函数定义中的变量称为形式参数
类似于 js 中，局部变量和全部变量可以名称一致，但是会以局部变量为主

#### 9. 数组
数组的形式基本和其他语言差不多
#### 10.指针
指针在以前接触过的语言中没有接触过, 使用 & 来获取指针地址, 使用 * 来获取指针的值, 使用 nil 来辨识空指针
#### 11.结构体
语法: 
```$xslt
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}
```
类似于 java 的 javabean
#### 12.切片 slice
语法: 
```$xslt
var identifier []type

var slice1 []type = make([]type, len)
make([]T, length, capacity)
也可以简写为

slice1 := make([]type, len)
```
go 数组的长度是不可改变的, 在一些特定的场景下，我们需要更灵活的方法, 使用切片可以满足这些需求
切片不需要说明长度

#### 13.range
range 关键字用语 for 循环中迭代数组(array), 切片(slice)，链表(channel)或集合(map) 的元素
#### 14.map
语法: 
```$xslt
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable = make(map[key_data_type]value_data_type)
```
map 集合类似于java中的map 和 python 中的字典, 使用 range 可以遍历. 是无需的键值对的集合
#### 15.递归函数
go 的递归和其他语言类似, 都是在方法或者函数本身调用本身
#### 16.类型转换
语法: 
```$xslt
type_name(expression)
```
用语将一种数据类型的变量转换为另外一种类型的变量,和 java 类似可以在变量前加括号强制转换
#### 17.接口
golang 的接口和 java 的接口 和 scala 的特质有些不一样,其他任何类型只有实现了这些方法就是实现了这个接口
空接口或者最小接口, 不包含任何方法, 它对实现不做任何要求
语法: 
```cgo
/*赋任何类型的值*/
type Any interface{} 
```
#### 18.错误处理
定义: 
```$xslt
type error interface {
    Error() string
}
```
错误处理
#### 19. 反射 reflect
reflect 包实现了运行时反射, 允许程序操作任意类型的对象
语法: 
```cgo
/*TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。*/
func TypeOf(i interface{}) Type
/*ValueOf返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值。*/
func ValueOf(i interface{}) Value
/*获取变量的类别, 返回一个常量*/
reflect.Value.Kind
/*转换成 interface{} 类型*/
reflect.Value.Interface()

```
#### 20. 文件读取 os.File

参看文档: `https://www.cnblogs.com/suoning/p/7225096.html`

#### 21.Goroutine
语法: 
```cgo
go 
```
Go 语言的主要的功能在于令人简易使用的并行设计, 这个方法叫做 Goroutine, 通过 
#### channel
channel, 管道, 队列, 先进先出, 用来异步传递数据. channel 加上 goroutine , 就形成了一种既简单又强大的请求处理模型, 使高并发和线程同步质检代码的编写变得异常简单
线程安全, 多个 goroutine 同时访问, 不需要加锁
channel 是有类型的, 一个整数的 channel 只能存放整数