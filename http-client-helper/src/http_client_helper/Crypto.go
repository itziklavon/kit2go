package http_client_helper

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"fmt"
)

var openSSLSaltHeader string = "Salted_"

type OpenSSLCreds struct {
	key []byte
	iv  []byte
}

var (
	ErrEmptyPassword  = fmt.Errorf("the password used was empty")
	ErrExceededLength = fmt.Errorf("derived key too long for md5")
)

const (
	AlgoPBEWithMD5AndDES = "PBEWithMD5AndDES"
)

const (
	MaxLenMD5 = 20
)

func PBKDF1MD5(pass, salt []byte, count, l int) ([]byte, error) {
	if l > MaxLenMD5 {
		return nil, ErrExceededLength
	}

	derived := make([]byte, len(pass)+len(salt))
	copy(derived, pass)
	copy(derived[len(pass):], salt)

	for i := 0; i < count; i++ {
		dr := md5.Sum(derived)
		derived = dr[:]
	}

	return derived[:l], nil
}

func DecryptPassword(encrypted []byte, password string) []byte {
	if len(encrypted) < des.BlockSize {
		return nil
	}

	salt := []byte{0xA9, 0x9B, 0xC8, 0x32,
		0x56, 0x35, 0xE3, 0x03}
	ct := encrypted[:]

	key, err := PBKDF1MD5([]byte(password), salt, 19, des.BlockSize*2)
	if err != nil {
		return nil
	}

	iv := key[des.BlockSize:]
	key = key[:des.BlockSize]

	b, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	dst := make([]byte, len(ct))
	bm := cipher.NewCBCDecrypter(b, iv)
	bm.CryptBlocks(dst, ct)

	return dst
}

type Decryptor struct {
	Password, Algorithm string
}

func (d Decryptor) Decrypt(bs []byte) ([]byte, error) {
	switch d.Algorithm {
	case AlgoPBEWithMD5AndDES:
		if d.Password == "" {
			return nil, ErrEmptyPassword
		}
	}
	return nil, fmt.Errorf("unknown jasypt algorithm")
}
