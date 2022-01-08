package main

func sliceContains[T comparable](slice []T, search T) bool {
	for _, elem := range slice {
		if elem == search {
			return true
		}
	}
	return false
}

func sliceToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, elem := range slice {
		set[elem] = true
	}
	return set
}

func setToSlice[T comparable](set map[T]bool) []T {
	var slice []T
	for elem, in := range set {
		if in {
			slice = append(slice, elem)
		}
	}
	return slice
}

func uniquifySlice[T comparable](slice []T) []T {
	set := sliceToSet(slice)
	return setToSlice(set)
}
