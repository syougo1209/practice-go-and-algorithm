package main

type Store struct {
	value     string
	timestamp int
}

type TimeMap struct {
	store map[string][]Store
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]Store)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Store{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.store[key]
	l := 0
	r := len(arr)
	ind := -1
	for l < r {
		mid := (l + r) / 2
		if arr[mid].timestamp == timestamp {
			return arr[mid].value
		} else if arr[mid].timestamp < timestamp {
			ind = mid
			l = mid + 1
		} else if arr[mid].timestamp > timestamp {
			r = mid
		}
	}
	if ind == -1 {
		return ""
	} else {
		return arr[ind].value
	}
}
