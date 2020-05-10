func main() {
	defer func() {
		if err := recover(); err != nil {
			if myErr, ok := err.(*MyErr); ok {
				fmt.Println(myErr)
			}
			panic(err) // not allowed
		}
	}()
	panic("") // allowed
}