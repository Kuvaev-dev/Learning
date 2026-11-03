package namedreturn

// Only Short Functions
func Return(start string, end string) (endstart, startend string) {
	endstart = end + start
	startend = start + end
	return
}
