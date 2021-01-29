package main

type Slice struct {
	array [5]int
	len   int
}

func Append(sl *Slice, elem int) *Slice {
	if len == len(sl.array) {
		lenInitial := len(sl.array)

		arrayNew = make([10]int)
		for idx, elem := range sl.array {
			arrayNew[idx] = elem
		}
		sl.array = arrayNew
	}
	sl.array[sl.len] = elem
	sl.len++
	return sl
}
