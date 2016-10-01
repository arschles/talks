i := 0
for {
	select {
	case <-timer.C:
		return
	case ticker.C:
		fn(i)
		i++
	}
}
