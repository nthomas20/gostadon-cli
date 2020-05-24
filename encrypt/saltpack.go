package encrypt

import (
	"bytes"

	"github.com/keybase/saltpack"
)

// First run informed by: https://godoc.org/github.com/keybase/saltpack
/**
 * Key Management: Saltpack makes no attempt to manage keys. We assume the wrapping application has a story for key management.
 * Encoding: Saltpack has two encoding modes: binary and armored. In armored mode, saltpack outputs in Base62-encoding, suitable for publication into any manner of Web settings without fear of markup-caused mangling.
 */

func Encrypt() {
	saltpack.NewEncryptStream()
}

func exampleEncode(plaintext []byte) []byte {
	buf := bytes.NewBuffer(nil)
	encoder := newEncoder(buf)
	var headerBytes []byte
	// A real implementation would encode the header into
	// headerBytes.
	err := encoder.Encode(headerBytes)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(plaintext); i++ {
		block := exampleBlock{
			PayloadCiphertext: []byte{^plaintext[i]},
			Seqno:             packetSeqno(i + 1),
			IsFinal:           i == len(plaintext)-1,
		}
		err := encoder.Encode(block)
		if err != nil {
			panic(err)
		}
	}
	return buf.Bytes()
}
