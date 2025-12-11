# json


## 1

//反射转换any-->reflect.Value
func unpackEface(i any) Value {
e := (*abi.EmptyInterface)(unsafe.Pointer(&i))
// NOTE: don't read e.word until we know whether it is really a pointer or not.
t := e.Type
if t == nil {
return Value{}
}
f := flag(t.Kind())
if t.IfaceIndir() {
f |= flagIndir
}
return Value{t, e.Data, f}
}

e := (*abi.EmptyInterface)(unsafe.Pointer(&i))
unsafe.Pointer能把任意指针变成可转换为其他指针类型的指针类型（通用指针，不携带类型信息）

由于语法规定，不能直接访问any（空接口）的data、type（为什么这么设计？）。通过取地址转换成通用
指针再使用类型转换，最终访问到any的具体内存，拿到any的data和type



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

## ai解释（后续验证）

### 不能直接访问any的data和type的原因
这样会类型不安全，比如：
i := any(42)
i.Data = float
这样直接访问i，把i当作int处理，实际上是float，会产生垃圾甚至崩溃。
