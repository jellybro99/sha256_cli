package main

func sha256Hash(message string) []uint32 {
	paddedMessageBytes := preprocessMessage(message)
	hashArray := preprocessHash()

	return hashArray
}

func preprocessMessage(message string) []byte {
	messageBytes := []byte(message)
	// append 1
	// append k 0s
	// append 64 bit representation of l

	paddedMessage := messageBytes

	return paddedMessage
}

func preprocessHash() []uint32 {
	return []uint32{
		0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19,
	}
}
