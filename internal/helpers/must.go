package helpers

//Must panic on err != nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
