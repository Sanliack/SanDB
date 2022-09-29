package main

type LRUCache struct {
	amap map[int]*ListNode
	Cap  int
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Pre  *ListNode
	Next *ListNode
	Val  int
}

func Constructor(capacity int) LRUCache {
	h := &ListNode{
		nil, nil, 0,
	}
	p := &ListNode{
		h, nil, 0,
	}
	h.Next = p
	return LRUCache{
		make(map[int]*ListNode, capacity),
		capacity,
		h,
		p,
	}
}

func (this *LRUCache) Get(key int) int {
	if data, ok := this.amap[key]; ok {
		this.InsertHead(data)
		return data.Val
	}
	return -1

}

func (this *LRUCache) InsertHead(data *ListNode) {
	data.Pre.Next = data.Next
	data.Next.Pre = data.Pre

	data.Next = this.Head.Next
	this.Head.Next.Pre = data

	this.Head.Next = data
	data.Pre = this.Head
}

func (this *LRUCache) Put(key int, value int) {
	if data, ok := this.amap[key]; ok {
		data.Val = value
		this.InsertHead(data)
		return
	}
	if len(this.amap)+1 >= this.Cap {
		this.Tail.Pre.Pre.Next = this.Tail
		this.Tail.Pre = this.Tail.Pre.Pre
	}
	aa := &ListNode{this.Head, this.Head.Next, value}
	this.Head.Next.Pre = aa
	this.Head.Next = aa
}
