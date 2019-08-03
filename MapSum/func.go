package mapSum

//https://leetcode-cn.com/problems/map-sum-pairs/submissions/

type Base struct {
	key  byte
	val  int
	next *MapSum
}

type MapSum struct {
	//XXX: only support string consisted of byte, maybe should support rune
	base [256]Base
}

func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	if key == "" {
		return
	}
	for _, k := range key[:len(key)-1] {
		if this.base[k].next == nil {
			this.base[k].next = &MapSum{}
		}
		this = this.base[k].next
	}
	this.base[key[len(key)-1]].val = val
}

func (this *MapSum) Sum(prefix string) int {
	for _, k := range prefix[:len(prefix)-1] {
		if this.base[k].next == nil {
			return 0
		}
		this = this.base[k].next
	}
	return this.base[prefix[len(prefix)-1]].val + this.base[prefix[len(prefix)-1]].next.nodeSum()
}

func (this *MapSum) nodeSum() int {
	sum := 0
	if this == nil {
		return sum
	}
	for _, b := range this.base {
		sum += b.val
		if b.next != nil {
			sum += b.next.nodeSum()
		}
	}
	return sum
}
