package sanmodel

const (
	// 服务器返回客户端
	Nil    uint16 = iota
	Syntax        // syntax拼写错误
	Err
	Suc

	//客户端给服务器
	Database
	Str
	Set

	// str
	Str_Put
	Str_Del
	Str_Get
	Str_Clean
	Str_Merge

	//set
	Set_Add
	Set_Card
	Set_Member
	Set_Pop
	Set_DelByKey
	Set_Merge
	Set_Clean
	Set_IsMember

	// set_status
	Set_Del
)

//msgtype
