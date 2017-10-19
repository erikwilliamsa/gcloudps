package workers

//Consumes until toggled off
func Consume(toggle <-chan struct{}) int {
	i := 0
	for {
		select {
		case <-toggle:
			return i
		default:
			i++
		}

	}
}
