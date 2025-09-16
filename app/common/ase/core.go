package ase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// AesEncryptGCM 使用 AES-256-GCM 模式加密数据
func AesEncryptGCM(key, plaintext []byte) ([]byte, error) {
	// 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// 创建 GCM 模式的 AEAD 接口
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// 生成一个随机的 Nonce (一次性使用的随机数)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// 加密数据
	// GCM.Seal() 会将 Nonce 和密文拼接在一起
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// AesDecryptGCM 使用 AES-256-GCM 模式解密数据
func AesDecryptGCM(key, ciphertext []byte) ([]byte, error) {
	// 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// 创建 GCM 模式的 AEAD 接口
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// 验证并解密
	// 因为加密时我们将 Nonce 和密文拼接了，所以这里需要先分离它们
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}

	return plaintext, nil
}
