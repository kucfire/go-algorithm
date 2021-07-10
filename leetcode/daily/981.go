/*
	leetcode tag:981 title:基于时间的键值存储
*/

package daily

import "sort"

type Team struct {
	value    string
	timetamp int
}

type TimeMap struct {
	Index map[string][]Team
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{map[string][]Team{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Index[key] = append(this.Index[key], Team{value: value, timetamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	team := this.Index[key]
	result := sort.Search(len(team), func(i int) bool { return team[i].timetamp > timestamp })
	if result > 0 {
		return team[result-1].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
