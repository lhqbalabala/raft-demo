package pkg

import (
	"fmt"

	"math/rand"
	"sync"
	"time"
)

type Value struct {
	ReSetTime   time.Time
	TimeoutTime time.Duration
	Data        string
}
type Node struct {
	Key   string
	Value *Value
	Hash  int
	Next  *Node
}

func (this *Node) GetValue() *Value {
	if this == nil {
		return nil
	}
	return this.Value
}
func (this *Node) GetKey() string {
	return this.Key
}
func (this *KvStorage) Hash(key string) int {
	hash := 0
	val := []byte(key)
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(val[i])
	}
	tmp := hash
	return hash ^ tmp>>16
}

type KvStorage struct {
	Table      []*Node
	Threshold  int
	Size       int
	LoadFactor float32
	Lock       sync.Mutex
}

func (this *KvStorage) PutKey(key string, value string, duration time.Duration) {

	this.PutVal(key, &Value{Data: value, ReSetTime: time.Now(), TimeoutTime: duration})
}
func (this *KvStorage) DeleteKey(key string) {

	this.Remove(key)
}
func (this *KvStorage) GetKey(key string) string {

	return this.GetNode(key).Value.Data
}
func (this *KvStorage) runercleanTimer() {
	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		<-ticker.C
		n := len(this.Table)
		if this.Table != nil {
			for {
				this.CheckAndRemoveTimeoutNode(this.Table[rand.Intn(n)])
				n >>= 1
				if n == 0 {
					break
				}
			}
		}
	}
}
func (this *KvStorage) PutVal(key string, value *Value) {
	hash := this.Hash(key)
	var table []*Node
	var p *Node
	table = this.Table
	n := 0
	if table == nil || len(table) == 0 {
		table = this.resize()
	}
	n = len(table)
	fmt.Println(n, this.Size, this.Threshold)
	i := (n - 1) & hash
	p = table[i]
	if p == nil {
		table[i] = &Node{Key: key, Value: value, Hash: hash, Next: nil}
	} else {
		var e *Node
		var k string
		k = p.Key
		//key完全相同，直接替换
		if p.Hash == hash && k == key {
			e = p
		} else {
			for {
				e = p.Next
				if e == nil {
					p.Next = &Node{Key: key, Value: value, Hash: hash, Next: nil}
					break
				}
				k = e.Key
				if e.Hash == hash && k == key {
					break
				}
				p = e
			}
		}
		if e != nil {
			e.Value = value
		}
	}
	//检查，移除过期的数据
	//顺带清除边上的一些数据,清理log(n)次
	index := (n - 1) & hash
	cleanlen := n
	cleanindex := index
	for {
		//fmt.Println("清除了一次,当前clean为", cleanlen, cleanindex)
		if cleanindex+1 >= n {
			cleanindex = 0
		} else {
			cleanindex++
		}
		table[cleanindex] = this.CheckAndRemoveTimeoutNode(table[cleanindex])
		cleanlen >>= 1
		if cleanlen == 0 {
			break
		}
	}
	this.Size++
	if (this.Size) > this.Threshold {
		this.resize()
	}

}
func (this *KvStorage) GetNode(key string) *Node {
	hash := this.Hash(key)
	var tab []*Node
	var first, e *Node
	var n int
	var k string

	/*这里有3个判断，前2个判断是确定数组不为空，如果为空直接返回null；
	 *第3个判断，通过key的hash值找到该key值在数组中的下标；
	 *如果该数组下标对应的位置为空，则说明该key没有对应的value；
	 */
	tab = this.Table
	n = len(tab)

	//检查，移除过期的数据
	//顺带清除边上的一些数据,清理log(n)次
	index := (n - 1) & hash
	cleanlen := n
	cleanindex := index
	for {
		//fmt.Println("清除了一次,当前clean为", cleanlen, cleanindex)
		if cleanindex >= n {
			cleanindex = 0
		} else {
			cleanindex++
		}
		tab[cleanindex] = this.CheckAndRemoveTimeoutNode(tab[cleanindex])
		cleanlen >>= 1
		if cleanlen == 0 {
			break
		}
	}

	first = tab[index]

	if tab != nil && n > 0 && first != nil {
		//比较hash值，hash值相同之后才比较key是否相同，如果key相同，则返回该对象；
		k = first.Key
		if first.Hash == hash && k == key {
			return first
		}
		e = first.Next
		if e != nil {
			for {
				k = e.Key
				if e.Hash == hash && k == key {
					//更新访问时间
					e.Value.ReSetTime = time.Now()
					return e
				}
				e = e.Next
				if e == nil {
					break
				}
			}
		}
	}
	return nil
}
func (this *KvStorage) CheckAndRemoveTimeoutNode(node *Node) *Node {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	if node == nil {
		return nil
	} else if node.Next == nil {
		if (time.Since(node.Value.ReSetTime) > node.Value.TimeoutTime) && node.Value.TimeoutTime != 0 {
			this.Size--
			node = nil
		}
		return node
	} else {
		e := node
		var newnode *Node
		var newnodetail *Node
		for {
			//如果未超时或者永远不超时
			if time.Since(e.Value.ReSetTime) <= e.Value.TimeoutTime || (e.Value.TimeoutTime == 0) {
				if newnode == nil {
					newnode = e
				} else {
					newnodetail.Next = e
				}
				newnodetail = e
			} else {
				this.Size--
			}
			e = e.Next
			if e == nil {
				break
			}
		}
		if newnodetail != nil {
			newnodetail.Next = nil
		}
		node = newnode
	}
	return node
}
func (this *KvStorage) resize() []*Node {
	oldTab := this.Table
	if oldTab == nil || len(oldTab) == 0 {
		newtable := make([]*Node, 16)
		this.Table = newtable
		this.Size = 0
		this.Threshold = 12
		this.LoadFactor = 0.75
		return newtable
	}
	var oldCap int
	oldCap = len(oldTab)
	oldThr := this.Threshold
	newCap, newThr := 0, 0
	if oldCap > 0 {
		newCap = oldCap << 1
		if oldCap >= 1073741824 {
			this.Threshold = 2147483647
			return oldTab
		} else if newCap < 1073741824 && oldCap >= 16 {
			newThr = oldThr << 1
		}
	} else if oldThr > 0 {
		newCap = oldThr
	} else {
		newCap = 16
		newThr = 12
	}
	if newThr == 0 {
		ft := float32(newCap) * this.LoadFactor
		if newCap < 1073741824 && ft < float32(1073741824) {
			newThr = int(ft)
		} else {
			newThr = 2147483647
		}
	}
	this.Threshold = newThr
	newTab := make([]*Node, newCap)
	this.Table = newTab
	if oldTab != nil {
		for j := 0; j < oldCap; j++ {
			e := oldTab[j]
			if e != nil {
				oldTab[j] = nil
				if e.Next == nil {
					newTab[e.Hash&(newCap-1)] = e
				} else { // preserve order
					var loHead, loTail, hiHead, hiTail, next *Node

					for {
						next = e.Next
						if (e.Hash & oldCap) == 0 {
							if time.Since(e.Value.ReSetTime) <= e.Value.TimeoutTime || (e.Value.TimeoutTime == 0) {
								if loTail == nil {
									loHead = e
								} else {
									loTail.Next = e
								}
								loTail = e
							} else {
								this.Size--
							}
						} else {
							if time.Since(e.Value.ReSetTime) <= e.Value.TimeoutTime || (e.Value.TimeoutTime == 0) {
								if hiTail == nil {
									hiHead = e
								} else {
									hiTail.Next = e
								}
								hiTail = e
							} else {
								this.Size--
							}
						}
						e = next
						if e == nil {
							break
						}
					}

					if loTail != nil {
						loTail.Next = nil
						newTab[j] = loHead
					}
					if hiTail != nil {
						hiTail.Next = nil
						newTab[j+oldCap] = hiHead
					}
				}
			}
		}
	}
	return newTab
}
func (this *KvStorage) Remove(key string) {
	hash := this.Hash(key)
	var tab []*Node
	var p *Node
	var n, index int

	tab = this.Table
	if tab == nil {
		return
	}
	n = len(tab)
	//fmt.Println((n - 1) & hash)
	p = tab[(n-1)&hash]
	index = (n - 1) & hash
	//fmt.Println(p, 111111)
	if tab != nil && n > 0 && p != nil {
		var node, e *Node
		k := p.Key
		e = p.Next
		if p.Hash == hash && k == key {
			node = p
		} else if e != nil {
			for {
				k = e.Key
				if e.Hash == hash && (key == k) {
					node = e
					break
				}
				p = e
				e = e.Next
				if e == nil {
					break
				}
			}
		}
		if node != nil {
			if node == p {
				tab[index] = node.Next
			} else {
				p.Next = node.Next
			}
			this.Size--
			return
		}
	}
	return
}
