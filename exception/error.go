package exception

func PanicLogging(err any) {
	if err != nil {
		panic(err)
	}
}