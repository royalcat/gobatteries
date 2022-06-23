package gobatteries

// operations on constant size slices

func InSlice(arr []int64, val int64) bool {

	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func SliceDiff[V comparable](a, b []V) []V {
	m := make(map[V]bool)
	diff := make([]V, 0, len(a))

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

func CompareByteArrays(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
