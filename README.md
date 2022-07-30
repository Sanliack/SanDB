# SanDBV1.0

基于TCP实现的一个Key-value数据库（数据可以持久化，保存到本地，需要数据时从文件中读取出指定数据）。提供有Put，Get，Del,Clean,Merge方法。

使用方法： 
server端：
  server := sanmodel.NewServerModel("SanDB Server V1.0", "127.0.0.1:3665")
	server.Server()
  
  
client端：  
  client := sanmodel.NewClientModel()
	control, err := client.Connect("127.0.0.1:3665")
  err := control.Put([]byte("key1"),[]byte("val1"))
  
 
 
 
 
 

