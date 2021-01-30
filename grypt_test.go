package grypt

import "testing"

func TestEnd2EndCrypto(t *testing.T) {

	tests := []struct {
		name       string
		plaintext  string
		cipherbyte []byte
		encodedB64 string
	}{
		{name: "simpleText", plaintext: "hello", encodedB64: ""},
		{name: "emptyText", plaintext: "", encodedB64: ""},
	}

	for _, test := range tests {

		encryptedText, err := B64Encrypt([]byte(test.plaintext))
		if err != nil {
			t.Errorf("decrypting %s failed: %v", test.name, err)
		}

		decryptedText, err := B64Decrypt(encryptedText)
		if err != nil {
			t.Errorf("decrypting %s failed: %v", test.name, err)
		}

		if string(decryptedText) != test.plaintext {
			t.Errorf("encrypt/decrypt of %s failed: %v", test.name, err)
		}
	}
}

func testShortText(t *testing.T) {
	t.Error("Not Implemented!")
}

func testLongText(t *testing.T) {
	t.Error("Not Implemented!")
}

func testEmptyString(t *testing.T) {
	t.Error("Not Implemented!")
}

func testBinaryContent(t *testing.T) {
	t.Error("Not Implemented!")
}
