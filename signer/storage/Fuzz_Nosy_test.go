package storage

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_AESEncryptedStorage_Del__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filename string, k2 []byte, k3 string) {
		s := NewAESEncryptedStorage(filename, k2)
		s.Del(k3)
	})
}

func Fuzz_Nosy_AESEncryptedStorage_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filename string, k2 []byte, k3 string) {
		s := NewAESEncryptedStorage(filename, k2)
		s.Get(k3)
	})
}

func Fuzz_Nosy_AESEncryptedStorage_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filename string, k2 []byte, k3 string, value string) {
		s := NewAESEncryptedStorage(filename, k2)
		s.Put(k3, value)
	})
}

func Fuzz_Nosy_AESEncryptedStorage_readEncryptedStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filename string, key []byte) {
		s := NewAESEncryptedStorage(filename, key)
		s.readEncryptedStorage()
	})
}

func Fuzz_Nosy_AESEncryptedStorage_writeEncryptedStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var creds map[string]storedCredential
		fill_err = tp.Fill(&creds)
		if fill_err != nil {
			return
		}

		s := NewAESEncryptedStorage(filename, key)
		s.writeEncryptedStorage(creds)
	})
}

func Fuzz_Nosy_EphemeralStorage_Del__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EphemeralStorage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Del(key)
	})
}

func Fuzz_Nosy_EphemeralStorage_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EphemeralStorage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Get(key)
	})
}

func Fuzz_Nosy_EphemeralStorage_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EphemeralStorage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Put(key, value)
	})
}

func Fuzz_Nosy_NoStorage_Del__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *NoStorage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Del(key)
	})
}

func Fuzz_Nosy_NoStorage_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *NoStorage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Get(key)
	})
}

func Fuzz_Nosy_NoStorage_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *NoStorage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Put(key, value)
	})
}

func Fuzz_Nosy_Storage_Del__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key string) {
		_x1 := NewEphemeralStorage()
		_x1.Del(key)
	})
}

func Fuzz_Nosy_Storage_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key string) {
		_x1 := NewEphemeralStorage()
		_x1.Get(key)
	})
}

func Fuzz_Nosy_Storage_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key string, value string) {
		_x1 := NewEphemeralStorage()
		_x1.Put(key, value)
	})
}

func Fuzz_Nosy_decrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, nonce []byte, ciphertext []byte, additionalData []byte) {
		decrypt(key, nonce, ciphertext, additionalData)
	})
}

func Fuzz_Nosy_encrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, plaintext []byte, additionalData []byte) {
		encrypt(key, plaintext, additionalData)
	})
}
