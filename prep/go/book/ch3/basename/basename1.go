package main

// basename removes directory components and a .suffix
func basename(s string) string {
	//discard last '/' and everythign before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			//create a substring that is one character past the last '/'
			s = s[i+1:]
			//if we start at the _end_ we can dip out once we hit the first one
			break
		}
	}

	//preserve everything before the last .
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
      //grab everything before the .
      //weve already gotten rid of the directory stuff
			s = s[:i]
			break
		}
	}
  return s
}
