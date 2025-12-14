# json.Marshal
将传入的数据编码为json的数据格式（序列化）

## 1.反射转换any-->reflect.Value

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

## 2.拿到反射值（type,data）之后根据反射获得的type选择不同的编码器
在src中的这一句：typeEncoder(v.Type())
typeEncoder是根据type的不同返回不同的编码器，例如type=string，就会返回一个针对string的编码器。

## 3.编码器对反射值转换成json格式存入byte.Buffer
不同编码器针对不同的type，转换成[]byte之后存入最初初始化的EncodeState结构体，通过Bytes()复制
内存中保存的数据，使用append保存到新创建的[]byte

# json.Unmarshal
将传入的[]byte解析到指定格式的变量（指针类型，因为需要修改变量内容）内部（反序列化）

## 1.先对传入的指针进行反射拿到反射值（主要是获取变量类型）
使用reflect.Value（v）,根据data的类型选择不同的解码器。在src中首先根据传入的[]byte的类型（字面量
、数组、对象）三种类型选择不同的解码器。

## 2.在不同的解码器中对v进行处理


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
