package sha

func sha256Hash(messageString string) []uint32 {
	message := preprocessMessage(messageString)
	hashArray := preprocessHash()

	return hashArray
}
