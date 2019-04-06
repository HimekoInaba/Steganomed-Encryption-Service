package types

type Encryption struct {
	Message string `json:"message"`
	File    []byte `json:"file"`
}

type Decryption struct {
	File []byte `json:"file"`
}
