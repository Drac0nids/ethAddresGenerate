# ethAddresGenerate
 以0开头的以太坊地址生成器
 
编译

```
sudo go build -o ethAddresGenerate main.go
```

使用
```
Usage of ./ethAddresGenerate:                                                                                                                                       ─╯
  -num int
        0的位数 (default 4)
  -thread int
        线程数 (default 10)
```

示例：
<img width="950" alt="image" src="https://user-images.githubusercontent.com/31087166/188784381-c961326b-e139-4c5b-9df5-be1152f02a0d.png">

输出格式为：
```
私钥 地址
```
