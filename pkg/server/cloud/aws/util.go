package aws

func breakIntoChunks(vals []*string, chunkSize int) [][]*string {
	var chunked [][]*string
	for i := 0; i < len(vals); i += chunkSize {
		end := i + chunkSize
		if end > len(vals) {
			end = len(vals)
		}
		chunked = append(chunked, vals[i:end])
	}
	return chunked
}
