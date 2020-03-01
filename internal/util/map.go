package util

func CombineMapString(base, added map[string]string) map[string]string {
	for key, val := range added {
		base[key] = val
	}
	return base
}
