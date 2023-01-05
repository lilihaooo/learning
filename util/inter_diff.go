package util

// 交集
func Intersect(slice1 []int64, slice2 []int64) []int64 { // 取两个切片的交集
	m := make(map[int64]int64)
	n := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

// 差集
func Difference(a []int64, b []int64) []int64 {
	var inter []int64
	mp := make(map[int64]bool)
	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			delete(mp, s)
		}
	}
	for key := range mp {
		inter = append(inter, key)
	}
	return inter
}
