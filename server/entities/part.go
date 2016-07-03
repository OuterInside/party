package entities

// Part structure
type Part struct {
	ID    int   `json:"id"`
	Count int64 `json:"count"`
}

// Parts is part slice alias
type Parts []Part

func (parts Parts) Len() int {
	return len(parts)
}

func (parts Parts) Less(i, j int) bool {
	return parts[i].Count < parts[j].Count
}

func (parts Parts) Swap(i, j int) {
	parts[i], parts[j] = parts[j], parts[i]
}
