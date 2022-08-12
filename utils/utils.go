package utils

func GoSafe(f func()) {
	go func() {
		defer func() {
			recover()
		}()
		f()
	}()
}
