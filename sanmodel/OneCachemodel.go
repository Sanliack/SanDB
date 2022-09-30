package sanmodel

import (
	"SanDB/sanface"
	"SanDB/tools"
	"fmt"
	"sync"
)

type OneCache struct {
	DBname    string
	CacheMap  map[string]*tools.ListNode
	Head      *tools.ListNode
	Tail      *tools.ListNode
	CMLock    sync.Mutex
	MaxLength int
	Length    int
}

func (o *OneCache) Put(k string, v []byte) {
	o.CMLock.Lock()
	defer o.CMLock.Unlock()
	lnode, ok := o.CacheMap[k]
	if ok {
		lnode.Val = v
		o.MoveToHead(lnode)
		return
	}
	Newnode := &tools.ListNode{
		Val: v,
		Key: k,
	}
	if o.Length >= o.MaxLength {
		tailkey := o.RemoveTail()
		o.Length--
		delete(o.CacheMap, tailkey)
	}
	o.InsertHead(Newnode)
	o.Length++
	o.CacheMap[k] = Newnode
}

func (o *OneCache) Get(key string) ([]byte, bool) {
	o.CMLock.Lock()
	defer o.CMLock.Unlock()
	data, ok := o.CacheMap[key]
	if !ok {
		return nil, false
	}
	o.MoveToHead(data)
	return data.Val, true
}

func (o *OneCache) Clean() {
	o.CMLock.Lock()
	defer o.CMLock.Unlock()
	o.CacheMap = make(map[string]*tools.ListNode)
	o.Length = 0
	o.Head.Next = o.Tail
	o.Tail.Pre = o.Head
}

func (o *OneCache) Del(key string) {
	o.CMLock.Lock()
	defer o.CMLock.Unlock()
	node, ok := o.CacheMap[key]
	if !ok {
		return
	}
	o.RemoveNode(node)
	o.Length--
}

func (o *OneCache) TestDebug() {
	ahead := o.Head
	count := 0
	fmt.Printf("dbname:%s 当前cache长度为%d 最大长度为%d\n", o.DBname, o.Length, o.MaxLength)
	for ahead != nil {
		fmt.Printf("count:%d , key %s Val: %s\n", count, ahead.Key, string(ahead.Val))
		ahead = ahead.Next
		count++
	}
}

func (o *OneCache) RemoveNode(node *tools.ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
}

func (o *OneCache) RemoveTail() string {
	key := o.Tail.Pre.Key
	tmp := o.Tail.Pre.Pre
	o.Tail.Pre.Pre = nil
	o.Tail.Pre.Next = nil
	o.Tail.Pre = tmp
	tmp.Next = o.Tail
	return key
}

func (o *OneCache) MoveToHead(node *tools.ListNode) {
	o.RemoveNode(node)
	o.InsertHead(node)
}

func (o *OneCache) InsertHead(node *tools.ListNode) {
	node.Next = o.Head.Next
	node.Pre = o.Head
	o.Head.Next.Pre = node
	o.Head.Next = node
}

func NewOneCache(dbname string, maxlen int) sanface.OneCacheFace {
	newcache := &OneCache{
		DBname:    dbname,
		CacheMap:  make(map[string]*tools.ListNode),
		MaxLength: maxlen,
		Length:    0,
		Head:      &tools.ListNode{},
		Tail:      &tools.ListNode{},
	}
	newcache.Head.Next = newcache.Tail
	newcache.Tail.Pre = newcache.Head
	return newcache
}
