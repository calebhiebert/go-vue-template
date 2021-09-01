package graph

func sanitizeLimitOffset(lim *int, off *int) (limit int, offset int) {
	if lim == nil {
		limit = 10
	} else if *lim <= 0 {
		limit = 10
	} else if *lim > 100 {
		limit = 100
	} else {
		limit = *lim
	}

	if off == nil {
		offset = 0
	} else if *off <= 0 {
		offset = 0
	} else {
		offset = *off
	}

	return
}

