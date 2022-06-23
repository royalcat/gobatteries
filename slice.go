package gobatteries

func PaginateInt64[V any](x []V, skip int, size int) []V {
	if skip > len(x) {
		skip = len(x)
	}

	end := skip + size
	if end > len(x) {
		end = len(x)
	}

	return x[skip:end]
}

func PaginateInt64Invert[V any](x []V, skip int, size int) []V {
	if size < 0 || skip < 0 {
		return []V{}
	}

	if skip > len(x) {
		return []V{}
	}

	end := len(x) - skip - size
	if end < 0 {
		end = 0
	}

	out := make([]V, 0, size)
	for i := len(x) - 1 - skip; i >= end; i-- {
		out = append(out, x[i])
	}
	return out
}
