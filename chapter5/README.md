#### 函数 ####
特性：
（1）函数包含连续的执行语句, 可以在代码中通过调用函数来执行他们;
（2）函数将一个复杂的工作切分成更小的模块, 使得多人协作方式更容易;
（3）函数对它的使用者隐藏了实现细节。

# 函数的声明
  每个函数的声明都包含一个名字、一个形参列表、一个可选的返回列表以及函数体。
  (1) 函数的类型称为函数签名，当两个函数拥有相同的形参列表和返回列表时, 认为
     这两个函数的签名一致;
  (2) Go语言没有默认参数值的概念也不能指定实参名;
  (3) 实参是按值来传递的,函数接收到的是每个实参的副本;
  (4) 没有函数体表明是用其他语言实现。

# 函数变量
      函数是Go里面很重要的值: 函数变量有类型, 可以赋值给其他变量、传递或者
   从其函数返回。本身不可比较(nil除外)。
  (1) 函数变量使得函数将数据参数化, 也将函数的行为当做参数进行传递;


#### 注意 ####
1. 除非能够提供有效的错误消息或者能够很快地检测出错误, 否则在运行时检测断言
   条件就毫无意义;
2. 使用错误值来对预期错误和未知错误加以区分;
3. 当宕机发生时, 所有的延迟函数以倒序的方式执行, 从栈最上面的函数开始一直返
   回到main函数;
4. runtime包提供了转储栈的方法使得程序猿可以诊断错误;
      func main(){
         defer printStack()
         //...
      }
      func printStack() {
         var buf [4096]byte
         n := runtime.Stack(buf[:], false)
         os.Stdout.Write(buf[:n])
      }
   Go语言的宕机机制让延迟执行的函数在栈清理之前调用, 所以能转储栈信息。

5. 延迟执行的匿名函数能够改变外层函数返回给调用者的结果。
         func double(x int) int {
            return x + x
         }

         func triple(x int) (result int) {
            defer func() { result += x }()
            return double(x)
         }

6. 变长函数通常用于格式化字符串。
   ...int参数就像函数体内的slice, 但是变长函数的类型和一个带普通slice参数
   的函数的类型不同。

7. 当一个匿名函数需要递归时, 必须先声明一个变量然后将匿名函数赋给这个变量。