package util

func createSet(es ...string) map[string]bool {
	m := map[string]bool{}
	for _, v := range es {
		m[v] = true
	}
	return m
}

func setDifference(first map[string]bool, second map[string]bool) map[string]bool {
	res := map[string]bool{}
	for key := range first {
		if _, ok := second[key]; !ok {
			res[key] = true
		}
	}
	return res
}

func setUnion(first map[string]bool, second map[string]bool) map[string]bool {
	res := map[string]bool{}
	for key := range first {
		res[key] = true
	}
	for key := range second {
		res[key] = true
	}
	return res
}
