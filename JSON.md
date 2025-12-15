# json.Marshal
将传入的数据编码为json的数据格式（序列化）

序列化的大概步骤：
对v进行反射，获取储存在内存的类型和data-->v.type获取底层类型，使用对应类型的编码器进行加工-->
针对不同类型（处理基础类型string\int\bool）进行递归处理为基础类型，基础类型可以直接序列化为
json格式

## 1.反射转换any-->reflect.Value
e := (*abi.EmptyInterface)(unsafe.Pointer(&i))
unsafe.Pointer能把任意指针变成可转换为其他指针类型的指针类型（通用指针，不携带类型信息）

由于语法规定，不能直接访问any（空接口）的data、type（为什么这么设计？）。通过取地址转换成通用
指针再使用类型转换，最终访问到any的具体内存，拿到any的data和type

## 2.拿到反射值（type,data）之后根据反射获得的type选择不同的编码器
在src中的这一句：typeEncoder(v.Type())
typeEncoder是根据type的不同返回不同的编码器，例如type=string，就会返回一个针对string的编码器。

## 3.编码器对反射值转换成json格式存入byte.Buffer
不同编码器针对不同的type，转换成[]byte之后存入最初初始化的EncodeState结构体，通过Bytes()复制
内存中保存的数据，使用append保存到新创建的[]byte

# json.Unmarshal
将传入的[]byte解析到指定格式的变量（指针类型，因为需要修改变量内容）内部（反序列化）

反序列化到结构体的大概步骤：
首先获取v的反射值-->对传入的[]byte进行scan，获取第一节内容（根据每一节内容的类型，也就是
输入json的格式，数组、 对象、 字符串，根据格式的不同选择不同的解码器）-->拿到v的内存地址，
方便后面修改内容-->外层for循环：遍历结构体字段，内层for循环根据结构体字段的位置索引进行初始
化，拿到字段地址方便后续修改内容-->将json序列化到对应的字段，使用反射修改内存值


## 背景知识
any = interface{} = EmptyInterface

type EmptyInterface struct {
Type *Type
Data unsafe.Pointer
}
其中保存了动态类型信息（int,*int,结构体等）和数据信息(data是指向实际值的指针)，比如：
var a = 32
var i = any(&a),type = *int，data = 32

非空接口 = NonEmptyInterface

type NonEmptyInterface struct {
ITab *ITab
Data unsafe.Pointer
}
其中保存了方法集合、类型信息等和指向实际值的指针

## ai解释

### 不能直接访问any的data和type的原因
这样会类型不安全，比如：
i := any(42)
i.Data = float
这样直接访问i，把i当作int处理，实际上是float，会产生垃圾甚至崩溃。
