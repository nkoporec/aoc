package helpers


// Remove empty lines from the data
func RemoveEmptyItemsInSlice(data []string) []string {
	var result []string
	for _, line := range data {
		if line != "" {
			result = append(result, line)
		}
	}	
	return result
}
