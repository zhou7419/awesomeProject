package auth

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
)

func Verify(str string, key string) (string, error) {
	//str := "k1zzDeOeCPDlf9ZftpR3gwdJ2w0+MYokcTFSvxec6O+SUuCDKSIk6ePukUuXKbcx" // {'type':'note','time':1689220531}
	//key := []byte("87ca2f3b550d6b51")

	dst, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return string(dst), err
	}
	dst, _ = openssl.AesECBDecrypt(dst, []byte(key), openssl.PKCS7_PADDING)
	if err != nil {
		return string(dst), err
	}

	return string(dst), err
}

//func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
//	cipher, _ := aes.NewCipher(generateKey(key))
//	length := (len(origData) + aes.BlockSize) / aes.BlockSize
//	plain := make([]byte, length*aes.BlockSize)
//	copy(plain, origData)
//	pad := byte(len(plain) - len(origData))
//	for i := len(origData); i < len(plain); i++ {
//		plain[i] = pad
//	}
//	encrypted = make([]byte, len(plain))
//	// 分组分块加密
//	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
//	}
//
//	return encrypted
//}
//func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
//	cipher, _ := aes.NewCipher(generateKey(key))
//	decrypted = make([]byte, len(encrypted))
//	//
//	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
//	}
//
//	trim := 0
//	if len(decrypted) > 0 {
//		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
//	}
//
//	return decrypted[:trim]
//}
//func generateKey(key []byte) (genKey []byte) {
//	genKey = make([]byte, 16)
//	copy(genKey, key)
//	for i := 16; i < len(key); {
//		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
//			genKey[j] ^= key[i]
//		}
//	}
//	return genKey
//}
