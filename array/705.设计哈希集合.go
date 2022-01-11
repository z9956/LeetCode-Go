package array

type MyHashSet struct {
	m map[int]int
}

func Constructor() MyHashSet {
	return MyHashSet{m: make(map[int]int)}
}

func (this *MyHashSet) Add(key int) {
	this.m[key] = key
}

func (this *MyHashSet) Remove(key int) {
	delete(this.m, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.m[key]
	return ok
}
