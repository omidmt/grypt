package grypt

// Grypt is public interface defines methos can be used by clients and implemented internally.
type Grypt interface {
	Encrypt(text string, key ...string) (cipher, k []byte, err error)
	Decrypt(cipher, key []byte) (text string, err error)

	// EncryptBinary(data []byte, key ...[]byte) (cipher, k []byte, err error)
	// DecryptBinary(cipher, key []byte) (data []byte, err error)

	// Not Finalized
	// EncryptStream(data []byte, key ...[]byte) (cipher, k []byte, err error)
	// DecryptStream(cipher, key []byte) (data []byte, err error)
}
