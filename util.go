package main

func stringSliceContains(s []string, c ...string) bool {
	for _, str := range s {
		for _, chk := range c {
			if chk == str {
				return true
			}
		}
	}

	return false
}