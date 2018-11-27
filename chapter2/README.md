程序结构 -- 总结
==============

遵循普遍的构建规则:
    - 变量存储值
    - 简单表达式通过加和减等操作合并成大的
    - 基本类型通过数组和结构体进行聚合
    - 表达式通过if、for等控制语句来决定执行顺序
    - 语句被组织成函数用于隔离和复用
    - 函数被组织成源文件和包

命名:
    - 规则
        开头是字母或_, 后面可以跟任意数量的字符、数字和下划线, 并区分大小写
    - 风格
        "驼峰式", 像ASCII和HTML需要使用相同的大小写, 如: HTMLEscape/htmlEscape
    - 大小写
        包名总是小写, 能导出包开头大写
    - 名称长度
        尽可能短, 作用域越长就使用越长且更有意义的名称
    
Go有25个关键字, 只能出现在语法允许的地方, 不能作为名称:
    break    default     func   interface  select 
    case     defer       go     map        struct
    chan     else        goto   package    switch
    const    fallthrough if     range      type
    continue for         import return     var

#!+
变量的生命周期:
    在程序执行的过程中, 变量存在的时间段。通过是否可达来确定, 所以局部变量可在包
    含它的循坏的一次迭代之外继续存在。即包含它的循环已经返回, 它的存在可能还延续

    涉及到逃逸的概念:
        - 每一次变量逃逸都需要一次额外的内存分配过程。
#!-
