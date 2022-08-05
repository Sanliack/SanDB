# SanDBV1.3

基于TCP实现的一个Key-value数据库（数据可以持久化，保存到本地，需要数据时从文件中读取出指定数据）。

使用方法： 
server端：

```go
server := sanmodel.NewServerModel("ServerName", "127.0.0.1:3665")
server.Server()
```

client端：  

```go
client := sanmodel.NewClientModel()
control, err := client.Connect("127.0.0.1:3665","database_name")
if err != nil{
	return
}

err = control.Set().Sadd([]byte("c1"), []byte("valc1"))
...
```

## 可存储类型

### 1.Set

格式：一个key多个value。例子：key1 -- value1--value2--value3

 方法：

```
1.Sadd(key []byte, val []byte) error
// 为key添加一个val

2.Scard(key []byte) (int, error)
// 查询一个key的val个数

3.Smember(key []byte) ([][]byte, error)
// 查询一个key的所有val

4.Spop(key []byte, val []byte) error
// 删除一个key下指定val的值

5.SIsMember(key []byte, val []byte) (bool, error)
// 判断一个key-val是否存在

6.DelByKey(key []byte) error
// 删除一个key 以及key下所有val

7.MergeFile() error
// 整理文件（用于删除持久化文件中存在而内存中以及不在的数据）

8.Clean() error
//删除所有数据
```

### 2.str

格式：key-val（一个key对应一个val）

方法

```
1.Put(key []byte, val []byte) error
// 写入一个key-val对

2.Get(key []byte) ([]byte, error)
// 通过key获取val

3.Merge() error
// 整理文件（用于删除持久化文件中存在而内存中以及不在的数据）

4.Del(key []byte) error
// 删除指定key-val对

5.Clean() error
// 删除所有数据
```

## SanDB-Server架构设计

![](.\SanDB_Pic.png)

## SanDB存储单元设计

单个存储单元：

![](.\entry.png)

存储文件：

![](.\File.png)