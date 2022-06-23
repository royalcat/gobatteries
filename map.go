package gobatteries

func KeysOfMap[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
