package comma

func Comma(s string) string {
	n := len(s)

  //if length less than three return - BASE CASE
	if n <= 3 {
		return s
	}

  //If larger than three, call self recursively with last three digits removed
  // Add "," to result and attach the 
	return Comma(s[:n-3]) + "," + s[n-3:]
}
