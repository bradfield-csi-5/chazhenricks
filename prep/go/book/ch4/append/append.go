package append

func appendInt(x []int, y int) []int {
	var z []int //create an empty slice

	zlen := len(x) + 1 //get a potential length for z that is 1 + slice length passed in
	if zlen <= cap(x) {
		// if our potential length is less (or equal to) the cap of the current slice, then we can add to it
		//setting z equal to a slice of x uses the same underlying array
		z = x[:zlen] //make a new slice from existing array
	} else {
		zcap := zlen
		//if our new cap isnt at least double the size of the current length, then double it
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		//make a _brand new_ slice/underlying array
		z = make([]int, zlen, zcap)
		//copy contents from x to z
		copy(z, x)

	}
	z[len(x)] = y
	return z
}
