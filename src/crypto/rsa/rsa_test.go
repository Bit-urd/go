// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa_test

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/internal/boring"
	"crypto/rand"
	. "crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"internal/testenv"
	"math/big"
	"strings"
	"testing"
)

func TestKeyGeneration(t *testing.T) {
	for _, size := range []int{128, 1024, 2048, 3072} {
		priv, err := GenerateKey(rand.Reader, size)
		if err != nil {
			t.Errorf("GenerateKey(%d): %v", size, err)
		}
		if bits := priv.N.BitLen(); bits != size {
			t.Errorf("key too short (%d vs %d)", bits, size)
		}
		testKeyBasics(t, priv)
		if testing.Short() {
			break
		}
	}
}

func Test3PrimeKeyGeneration(t *testing.T) {
	size := 768
	if testing.Short() {
		size = 256
	}

	priv, err := GenerateMultiPrimeKey(rand.Reader, 3, size)
	if err != nil {
		t.Errorf("failed to generate key")
	}
	testKeyBasics(t, priv)
}

func Test4PrimeKeyGeneration(t *testing.T) {
	size := 768
	if testing.Short() {
		size = 256
	}

	priv, err := GenerateMultiPrimeKey(rand.Reader, 4, size)
	if err != nil {
		t.Errorf("failed to generate key")
	}
	testKeyBasics(t, priv)
}

func TestNPrimeKeyGeneration(t *testing.T) {
	primeSize := 64
	maxN := 24
	if testing.Short() {
		primeSize = 16
		maxN = 16
	}
	// Test that generation of N-prime keys works for N > 4.
	for n := 5; n < maxN; n++ {
		priv, err := GenerateMultiPrimeKey(rand.Reader, n, 64+n*primeSize)
		if err == nil {
			testKeyBasics(t, priv)
		} else {
			t.Errorf("failed to generate %d-prime key", n)
		}
	}
}

func TestImpossibleKeyGeneration(t *testing.T) {
	// This test ensures that trying to generate toy RSA keys doesn't enter
	// an infinite loop.
	for i := 0; i < 32; i++ {
		GenerateKey(rand.Reader, i)
		GenerateMultiPrimeKey(rand.Reader, 3, i)
		GenerateMultiPrimeKey(rand.Reader, 4, i)
		GenerateMultiPrimeKey(rand.Reader, 5, i)
	}
}

func TestGnuTLSKey(t *testing.T) {
	// This is a key generated by `certtool --generate-privkey --bits 128`.
	// It's such that de ≢ 1 mod φ(n), but is congruent mod the order of
	// the group.
	priv := parseKey(testingKey(`-----BEGIN RSA TESTING KEY-----
MGECAQACEQDar8EuoZuSosYtE9SeXSyPAgMBAAECEBf7XDET8e6jjTcfO7y/sykC
CQDozXjCjkBzLQIJAPB6MqNbZaQrAghbZTdQoko5LQIIUp9ZiKDdYjMCCCCpqzmX
d8Y7
-----END RSA TESTING KEY-----`))
	testKeyBasics(t, priv)
}

func testKeyBasics(t *testing.T, priv *PrivateKey) {
	if err := priv.Validate(); err != nil {
		t.Errorf("Validate() failed: %s", err)
	}
	if priv.D.Cmp(priv.N) > 0 {
		t.Errorf("private exponent too large")
	}

	msg := []byte("hi!")
	enc, err := EncryptPKCS1v15(rand.Reader, &priv.PublicKey, msg)
	if err != nil {
		t.Errorf("EncryptPKCS1v15: %v", err)
		return
	}

	dec, err := DecryptPKCS1v15(nil, priv, enc)
	if err != nil {
		t.Errorf("DecryptPKCS1v15: %v", err)
		return
	}
	if !bytes.Equal(dec, msg) {
		t.Errorf("got:%x want:%x (%+v)", dec, msg, priv)
	}
}

func TestAllocations(t *testing.T) {
	if boring.Enabled {
		t.Skip("skipping allocations test with BoringCrypto")
	}
	testenv.SkipIfOptimizationOff(t)

	m := []byte("Hello Gophers")
	c, err := EncryptPKCS1v15(rand.Reader, &test2048Key.PublicKey, m)
	if err != nil {
		t.Fatal(err)
	}

	if allocs := testing.AllocsPerRun(100, func() {
		p, err := DecryptPKCS1v15(nil, test2048Key, c)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(p, m) {
			t.Fatalf("unexpected output: %q", p)
		}
	}); allocs > 10 {
		t.Errorf("expected less than 10 allocations, got %0.1f", allocs)
	}
}

var allFlag = flag.Bool("all", false, "test all key sizes up to 2048")

func TestEverything(t *testing.T) {
	min := 32
	max := 560 // any smaller than this and not all tests will run
	if testing.Short() {
		min = max
	}
	if *allFlag {
		max = 2048
	}
	for size := min; size <= max; size++ {
		size := size
		t.Run(fmt.Sprintf("%d", size), func(t *testing.T) {
			t.Parallel()
			priv, err := GenerateKey(rand.Reader, size)
			if err != nil {
				t.Errorf("GenerateKey(%d): %v", size, err)
			}
			if bits := priv.N.BitLen(); bits != size {
				t.Errorf("key too short (%d vs %d)", bits, size)
			}
			testEverything(t, priv)
		})
	}
}

func testEverything(t *testing.T, priv *PrivateKey) {
	if err := priv.Validate(); err != nil {
		t.Errorf("Validate() failed: %s", err)
	}

	msg := []byte("test")
	enc, err := EncryptPKCS1v15(rand.Reader, &priv.PublicKey, msg)
	if err == ErrMessageTooLong {
		t.Log("key too small for EncryptPKCS1v15")
	} else if err != nil {
		t.Errorf("EncryptPKCS1v15: %v", err)
	}
	if err == nil {
		dec, err := DecryptPKCS1v15(nil, priv, enc)
		if err != nil {
			t.Errorf("DecryptPKCS1v15: %v", err)
		}
		err = DecryptPKCS1v15SessionKey(nil, priv, enc, make([]byte, 4))
		if err != nil {
			t.Errorf("DecryptPKCS1v15SessionKey: %v", err)
		}
		if !bytes.Equal(dec, msg) {
			t.Errorf("got:%x want:%x (%+v)", dec, msg, priv)
		}
	}

	label := []byte("label")
	enc, err = EncryptOAEP(sha256.New(), rand.Reader, &priv.PublicKey, msg, label)
	if err == ErrMessageTooLong {
		t.Log("key too small for EncryptOAEP")
	} else if err != nil {
		t.Errorf("EncryptOAEP: %v", err)
	}
	if err == nil {
		dec, err := DecryptOAEP(sha256.New(), nil, priv, enc, label)
		if err != nil {
			t.Errorf("DecryptOAEP: %v", err)
		}
		if !bytes.Equal(dec, msg) {
			t.Errorf("got:%x want:%x (%+v)", dec, msg, priv)
		}
	}

	hash := sha256.Sum256(msg)
	sig, err := SignPKCS1v15(nil, priv, crypto.SHA256, hash[:])
	if err == ErrMessageTooLong {
		t.Log("key too small for SignPKCS1v15")
	} else if err != nil {
		t.Errorf("SignPKCS1v15: %v", err)
	}
	if err == nil {
		err = VerifyPKCS1v15(&priv.PublicKey, crypto.SHA256, hash[:], sig)
		if err != nil {
			t.Errorf("VerifyPKCS1v15: %v", err)
		}
		sig[1] ^= 0x80
		err = VerifyPKCS1v15(&priv.PublicKey, crypto.SHA256, hash[:], sig)
		if err == nil {
			t.Errorf("VerifyPKCS1v15 success for tampered signature")
		}
		sig[1] ^= 0x80
		hash[1] ^= 0x80
		err = VerifyPKCS1v15(&priv.PublicKey, crypto.SHA256, hash[:], sig)
		if err == nil {
			t.Errorf("VerifyPKCS1v15 success for tampered message")
		}
		hash[1] ^= 0x80
	}

	opts := &PSSOptions{SaltLength: PSSSaltLengthAuto}
	sig, err = SignPSS(rand.Reader, priv, crypto.SHA256, hash[:], opts)
	if err == ErrMessageTooLong {
		t.Log("key too small for SignPSS with PSSSaltLengthAuto")
	} else if err != nil {
		t.Errorf("SignPSS: %v", err)
	}
	if err == nil {
		err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], sig, opts)
		if err != nil {
			t.Errorf("VerifyPSS: %v", err)
		}
		sig[1] ^= 0x80
		err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], sig, opts)
		if err == nil {
			t.Errorf("VerifyPSS success for tampered signature")
		}
		sig[1] ^= 0x80
		hash[1] ^= 0x80
		err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], sig, opts)
		if err == nil {
			t.Errorf("VerifyPSS success for tampered message")
		}
		hash[1] ^= 0x80
	}

	opts.SaltLength = PSSSaltLengthEqualsHash
	sig, err = SignPSS(rand.Reader, priv, crypto.SHA256, hash[:], opts)
	if err == ErrMessageTooLong {
		t.Log("key too small for SignPSS with PSSSaltLengthEqualsHash")
	} else if err != nil {
		t.Errorf("SignPSS: %v", err)
	}
	if err == nil {
		err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], sig, opts)
		if err != nil {
			t.Errorf("VerifyPSS: %v", err)
		}
		sig[1] ^= 0x80
		err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], sig, opts)
		if err == nil {
			t.Errorf("VerifyPSS success for tampered signature")
		}
		sig[1] ^= 0x80
		hash[1] ^= 0x80
		err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], sig, opts)
		if err == nil {
			t.Errorf("VerifyPSS success for tampered message")
		}
		hash[1] ^= 0x80
	}

	// Check that an input bigger than the modulus is handled correctly,
	// whether it is longer than the byte size of the modulus or not.
	c := bytes.Repeat([]byte{0xff}, priv.Size())
	err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], c, opts)
	if err == nil {
		t.Errorf("VerifyPSS accepted a large signature")
	}
	_, err = DecryptPKCS1v15(nil, priv, c)
	if err == nil {
		t.Errorf("DecryptPKCS1v15 accepted a large ciphertext")
	}
	c = append(c, 0xff)
	err = VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], c, opts)
	if err == nil {
		t.Errorf("VerifyPSS accepted a long signature")
	}
	_, err = DecryptPKCS1v15(nil, priv, c)
	if err == nil {
		t.Errorf("DecryptPKCS1v15 accepted a long ciphertext")
	}
}

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

func parseKey(s string) *PrivateKey {
	p, _ := pem.Decode([]byte(s))
	k, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		panic(err)
	}
	return k
}

var test2048Key = parseKey(testingKey(`-----BEGIN RSA TESTING KEY-----
MIIEnwIBAAKCAQBxY8hCshkKiXCUKydkrtQtQSRke28w4JotocDiVqou4k55DEDJ
akvWbXXDcakV4HA8R2tOGgbxvTjFo8EK470w9O9ipapPUSrRRaBsSOlkaaIs6OYh
4FLwZpqMNBVVEtguVUR/C34Y2pS9kRrHs6q+cGhDZolkWT7nGy5eSEvPDHg0EBq1
1hu6HmPmI3r0BInONqJg2rcK3U++wk1lnbD3ysCZsKOqRUms3n/IWKeTqXXmz2XK
J2t0NSXwiDmA9q0Gm+w0bXh3lzhtUP4MlzS+lnx9hK5bjzSbCUB5RXwMDG/uNMQq
C4MmA4BPceSfMyAIFjdRLGy/K7gbb2viOYRtAgEDAoIBAEuX2tchZgcGSw1yGkMf
OB4rbZhSSiCVvB5r1ew5xsnsNFCy1ducMo7zo9ehG2Pq9X2E8jQRWfZ+JdkX1gdC
fiCjSkHDxt+LceDZFZ2F8O2bwXNF7sFAN0rvEbLNY44MkB7jgv9c/rs8YykLZy/N
HH71mteZsO2Q1JoSHumFh99cwWHFhLxYh64qFeeH6Gqx6AM2YVBWHgs7OuKOvc8y
zUbf8xftPht1kMwwDR1XySiEYtBtn74JflK3DcT8oxOuCZBuX6sMJHKbVP41zDj+
FJZBmpAvNfCEYJUr1Hg+DpMLqLUg+D6v5vpliburbk9LxcKFZyyZ9QVe7GoqMLBu
eGsCgYEAummUj4MMKWJC2mv5rj/dt2pj2/B2HtP2RLypai4et1/Ru9nNk8cjMLzC
qXz6/RLuJ7/eD7asFS3y7EqxKxEmW0G8tTHjnzR/3wnpVipuWnwCDGU032HJVd13
LMe51GH97qLzuDZjMCz+VlbCNdSslMgWWK0XmRnN7Yqxvh6ao2kCgYEAm7fTRBhF
JtKcaJ7d8BQb9l8BNHfjayYOMq5CxoCyxa2pGBv/Mrnxv73Twp9Z/MP0ue5M5nZt
GMovpP5cGdJLQ2w5p4H3opcuWeYW9Yyru2EyCEAI/hD/Td3QVP0ukc19BDuPl5Wg
eIFs218uiVOU4pw3w+Et5B1PZ/F+ZLr5LGUCgYB8RmMKV11w7CyRnVEe1T56Ru09
Svlp4qQt0xucHr8k6ovSkTO32hd10yxw/fyot0lv1T61JHK4yUydhyDHYMQ81n3O
IUJqIv/qBpuOxvQ8UqwIQ3iU69uOk6TIhSaNlqlJwffQJEIgHf7kOdbOjchjMA7l
yLpmETPzscvUFGcXmwKBgGfP4i1lg283EvBp6Uq4EqQ/ViL6l5zECXce1y8Ady5z
xhASqiHRS9UpN9cU5qiCoyae3e75nhCGym3+6BE23Nede8UBT8G6HuaZZKOzHSeW
IVrVW1QLVN6T4DioybaI/gLSX7pjwFBWSJI/dFuNDexoJS1AyUK+NO/2VEMnUMhD
AoGAOsdn3Prnh/mjC95vraHCLap0bRBSexMdx77ImHgtFUUcSaT8DJHs+NZw1RdM
SZA0J+zVQ8q7B11jIgz5hMz+chedwoRjTL7a8VRTKHFmmBH0zlEuV7L79w6HkRCQ
VRg10GUN6heGLv0aOHbPdobcuVDH4sgOqpT1QnOuce34sQs=
-----END RSA TESTING KEY-----`))

var test3072Key = parseKey(testingKey(`-----BEGIN RSA TESTING KEY-----
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
-----END RSA TESTING KEY-----`))

var test4096Key = parseKey(testingKey(`-----BEGIN RSA TESTING KEY-----
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
-----END RSA TESTING KEY-----`))

func BenchmarkDecryptPKCS1v15(b *testing.B) {
	b.Run("2048", func(b *testing.B) { benchmarkDecryptPKCS1v15(b, test2048Key) })
	b.Run("3072", func(b *testing.B) { benchmarkDecryptPKCS1v15(b, test3072Key) })
	b.Run("4096", func(b *testing.B) { benchmarkDecryptPKCS1v15(b, test4096Key) })
}

func benchmarkDecryptPKCS1v15(b *testing.B, k *PrivateKey) {
	r := bufio.NewReaderSize(rand.Reader, 1<<15)

	m := []byte("Hello Gophers")
	c, err := EncryptPKCS1v15(r, &k.PublicKey, m)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	var sink byte
	for i := 0; i < b.N; i++ {
		p, err := DecryptPKCS1v15(r, k, c)
		if err != nil {
			b.Fatal(err)
		}
		if !bytes.Equal(p, m) {
			b.Fatalf("unexpected output: %q", p)
		}
		sink ^= p[0]
	}
}

func BenchmarkEncryptPKCS1v15(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		r := bufio.NewReaderSize(rand.Reader, 1<<15)
		m := []byte("Hello Gophers")

		var sink byte
		for i := 0; i < b.N; i++ {
			c, err := EncryptPKCS1v15(r, &test2048Key.PublicKey, m)
			if err != nil {
				b.Fatal(err)
			}
			sink ^= c[0]
		}
	})
}

func BenchmarkDecryptOAEP(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		r := bufio.NewReaderSize(rand.Reader, 1<<15)

		m := []byte("Hello Gophers")
		c, err := EncryptOAEP(sha256.New(), r, &test2048Key.PublicKey, m, nil)
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		var sink byte
		for i := 0; i < b.N; i++ {
			p, err := DecryptOAEP(sha256.New(), r, test2048Key, c, nil)
			if err != nil {
				b.Fatal(err)
			}
			if !bytes.Equal(p, m) {
				b.Fatalf("unexpected output: %q", p)
			}
			sink ^= p[0]
		}
	})
}

func BenchmarkEncryptOAEP(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		r := bufio.NewReaderSize(rand.Reader, 1<<15)
		m := []byte("Hello Gophers")

		var sink byte
		for i := 0; i < b.N; i++ {
			c, err := EncryptOAEP(sha256.New(), r, &test2048Key.PublicKey, m, nil)
			if err != nil {
				b.Fatal(err)
			}
			sink ^= c[0]
		}
	})
}

func BenchmarkSignPKCS1v15(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		hashed := sha256.Sum256([]byte("testing"))

		var sink byte
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s, err := SignPKCS1v15(rand.Reader, test2048Key, crypto.SHA256, hashed[:])
			if err != nil {
				b.Fatal(err)
			}
			sink ^= s[0]
		}
	})
}

func BenchmarkVerifyPKCS1v15(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		hashed := sha256.Sum256([]byte("testing"))
		s, err := SignPKCS1v15(rand.Reader, test2048Key, crypto.SHA256, hashed[:])
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			err := VerifyPKCS1v15(&test2048Key.PublicKey, crypto.SHA256, hashed[:], s)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSignPSS(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		hashed := sha256.Sum256([]byte("testing"))

		var sink byte
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s, err := SignPSS(rand.Reader, test2048Key, crypto.SHA256, hashed[:], nil)
			if err != nil {
				b.Fatal(err)
			}
			sink ^= s[0]
		}
	})
}

func BenchmarkVerifyPSS(b *testing.B) {
	b.Run("2048", func(b *testing.B) {
		hashed := sha256.Sum256([]byte("testing"))
		s, err := SignPSS(rand.Reader, test2048Key, crypto.SHA256, hashed[:], nil)
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			err := VerifyPSS(&test2048Key.PublicKey, crypto.SHA256, hashed[:], s, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

type testEncryptOAEPMessage struct {
	in   []byte
	seed []byte
	out  []byte
}

type testEncryptOAEPStruct struct {
	modulus string
	e       int
	d       string
	msgs    []testEncryptOAEPMessage
}

func TestEncryptOAEP(t *testing.T) {
	sha1 := sha1.New()
	n := new(big.Int)
	for i, test := range testEncryptOAEPData {
		n.SetString(test.modulus, 16)
		public := PublicKey{N: n, E: test.e}

		for j, message := range test.msgs {
			randomSource := bytes.NewReader(message.seed)
			out, err := EncryptOAEP(sha1, randomSource, &public, message.in, nil)
			if err != nil {
				t.Errorf("#%d,%d error: %s", i, j, err)
			}
			if !bytes.Equal(out, message.out) {
				t.Errorf("#%d,%d bad result: %x (want %x)", i, j, out, message.out)
			}
		}
	}
}

func TestDecryptOAEP(t *testing.T) {
	random := rand.Reader

	sha1 := sha1.New()
	n := new(big.Int)
	d := new(big.Int)
	for i, test := range testEncryptOAEPData {
		n.SetString(test.modulus, 16)
		d.SetString(test.d, 16)
		private := new(PrivateKey)
		private.PublicKey = PublicKey{N: n, E: test.e}
		private.D = d

		for j, message := range test.msgs {
			out, err := DecryptOAEP(sha1, nil, private, message.out, nil)
			if err != nil {
				t.Errorf("#%d,%d error: %s", i, j, err)
			} else if !bytes.Equal(out, message.in) {
				t.Errorf("#%d,%d bad result: %#v (want %#v)", i, j, out, message.in)
			}

			// Decrypt with blinding.
			out, err = DecryptOAEP(sha1, random, private, message.out, nil)
			if err != nil {
				t.Errorf("#%d,%d (blind) error: %s", i, j, err)
			} else if !bytes.Equal(out, message.in) {
				t.Errorf("#%d,%d (blind) bad result: %#v (want %#v)", i, j, out, message.in)
			}
		}
		if testing.Short() {
			break
		}
	}
}

func Test2DecryptOAEP(t *testing.T) {
	random := rand.Reader

	msg := []byte{0xed, 0x36, 0x90, 0x8d, 0xbe, 0xfc, 0x35, 0x40, 0x70, 0x4f, 0xf5, 0x9d, 0x6e, 0xc2, 0xeb, 0xf5, 0x27, 0xae, 0x65, 0xb0, 0x59, 0x29, 0x45, 0x25, 0x8c, 0xc1, 0x91, 0x22}
	in := []byte{0x72, 0x26, 0x84, 0xc9, 0xcf, 0xd6, 0xa8, 0x96, 0x04, 0x3e, 0x34, 0x07, 0x2c, 0x4f, 0xe6, 0x52, 0xbe, 0x46, 0x3c, 0xcf, 0x79, 0x21, 0x09, 0x64, 0xe7, 0x33, 0x66, 0x9b, 0xf8, 0x14, 0x22, 0x43, 0xfe, 0x8e, 0x52, 0x8b, 0xe0, 0x5f, 0x98, 0xef, 0x54, 0xac, 0x6b, 0xc6, 0x26, 0xac, 0x5b, 0x1b, 0x4b, 0x7d, 0x2e, 0xd7, 0x69, 0x28, 0x5a, 0x2f, 0x4a, 0x95, 0x89, 0x6c, 0xc7, 0x53, 0x95, 0xc7, 0xd2, 0x89, 0x04, 0x6f, 0x94, 0x74, 0x9b, 0x09, 0x0d, 0xf4, 0x61, 0x2e, 0xab, 0x48, 0x57, 0x4a, 0xbf, 0x95, 0xcb, 0xff, 0x15, 0xe2, 0xa0, 0x66, 0x58, 0xf7, 0x46, 0xf8, 0xc7, 0x0b, 0xb5, 0x1e, 0xa7, 0xba, 0x36, 0xce, 0xdd, 0x36, 0x41, 0x98, 0x6e, 0x10, 0xf9, 0x3b, 0x70, 0xbb, 0xa1, 0xda, 0x00, 0x40, 0xd5, 0xa5, 0x3f, 0x87, 0x64, 0x32, 0x7c, 0xbc, 0x50, 0x52, 0x0e, 0x4f, 0x21, 0xbd}

	n := new(big.Int)
	d := new(big.Int)
	n.SetString(testEncryptOAEPData[0].modulus, 16)
	d.SetString(testEncryptOAEPData[0].d, 16)
	priv := new(PrivateKey)
	priv.PublicKey = PublicKey{N: n, E: testEncryptOAEPData[0].e}
	priv.D = d
	sha1 := crypto.SHA1
	sha256 := crypto.SHA256

	out, err := priv.Decrypt(random, in, &OAEPOptions{MGFHash: sha1, Hash: sha256})

	if err != nil {
		t.Errorf("error: %s", err)
	} else if !bytes.Equal(out, msg) {
		t.Errorf("bad result %#v (want %#v)", out, msg)
	}
}

func TestEncryptDecryptOAEP(t *testing.T) {
	sha256 := sha256.New()
	n := new(big.Int)
	d := new(big.Int)
	for i, test := range testEncryptOAEPData {
		n.SetString(test.modulus, 16)
		d.SetString(test.d, 16)
		priv := new(PrivateKey)
		priv.PublicKey = PublicKey{N: n, E: test.e}
		priv.D = d

		for j, message := range test.msgs {
			label := []byte(fmt.Sprintf("hi#%d", j))
			enc, err := EncryptOAEP(sha256, rand.Reader, &priv.PublicKey, message.in, label)
			if err != nil {
				t.Errorf("#%d,%d: EncryptOAEP: %v", i, j, err)
				continue
			}
			dec, err := DecryptOAEP(sha256, rand.Reader, priv, enc, label)
			if err != nil {
				t.Errorf("#%d,%d: DecryptOAEP: %v", i, j, err)
				continue
			}
			if !bytes.Equal(dec, message.in) {
				t.Errorf("#%d,%d: round trip %q -> %q", i, j, message.in, dec)
			}
		}
	}
}

// testEncryptOAEPData contains a subset of the vectors from RSA's "Test vectors for RSA-OAEP".
var testEncryptOAEPData = []testEncryptOAEPStruct{
	// Key 1
	{"a8b3b284af8eb50b387034a860f146c4919f318763cd6c5598c8ae4811a1e0abc4c7e0b082d693a5e7fced675cf4668512772c0cbc64a742c6c630f533c8cc72f62ae833c40bf25842e984bb78bdbf97c0107d55bdb662f5c4e0fab9845cb5148ef7392dd3aaff93ae1e6b667bb3d4247616d4f5ba10d4cfd226de88d39f16fb",
		65537,
		"53339cfdb79fc8466a655c7316aca85c55fd8f6dd898fdaf119517ef4f52e8fd8e258df93fee180fa0e4ab29693cd83b152a553d4ac4d1812b8b9fa5af0e7f55fe7304df41570926f3311f15c4d65a732c483116ee3d3d2d0af3549ad9bf7cbfb78ad884f84d5beb04724dc7369b31def37d0cf539e9cfcdd3de653729ead5d1",
		[]testEncryptOAEPMessage{
			// Example 1.1
			{
				[]byte{0x66, 0x28, 0x19, 0x4e, 0x12, 0x07, 0x3d, 0xb0,
					0x3b, 0xa9, 0x4c, 0xda, 0x9e, 0xf9, 0x53, 0x23, 0x97,
					0xd5, 0x0d, 0xba, 0x79, 0xb9, 0x87, 0x00, 0x4a, 0xfe,
					0xfe, 0x34,
				},
				[]byte{0x18, 0xb7, 0x76, 0xea, 0x21, 0x06, 0x9d, 0x69,
					0x77, 0x6a, 0x33, 0xe9, 0x6b, 0xad, 0x48, 0xe1, 0xdd,
					0xa0, 0xa5, 0xef,
				},
				[]byte{0x35, 0x4f, 0xe6, 0x7b, 0x4a, 0x12, 0x6d, 0x5d,
					0x35, 0xfe, 0x36, 0xc7, 0x77, 0x79, 0x1a, 0x3f, 0x7b,
					0xa1, 0x3d, 0xef, 0x48, 0x4e, 0x2d, 0x39, 0x08, 0xaf,
					0xf7, 0x22, 0xfa, 0xd4, 0x68, 0xfb, 0x21, 0x69, 0x6d,
					0xe9, 0x5d, 0x0b, 0xe9, 0x11, 0xc2, 0xd3, 0x17, 0x4f,
					0x8a, 0xfc, 0xc2, 0x01, 0x03, 0x5f, 0x7b, 0x6d, 0x8e,
					0x69, 0x40, 0x2d, 0xe5, 0x45, 0x16, 0x18, 0xc2, 0x1a,
					0x53, 0x5f, 0xa9, 0xd7, 0xbf, 0xc5, 0xb8, 0xdd, 0x9f,
					0xc2, 0x43, 0xf8, 0xcf, 0x92, 0x7d, 0xb3, 0x13, 0x22,
					0xd6, 0xe8, 0x81, 0xea, 0xa9, 0x1a, 0x99, 0x61, 0x70,
					0xe6, 0x57, 0xa0, 0x5a, 0x26, 0x64, 0x26, 0xd9, 0x8c,
					0x88, 0x00, 0x3f, 0x84, 0x77, 0xc1, 0x22, 0x70, 0x94,
					0xa0, 0xd9, 0xfa, 0x1e, 0x8c, 0x40, 0x24, 0x30, 0x9c,
					0xe1, 0xec, 0xcc, 0xb5, 0x21, 0x00, 0x35, 0xd4, 0x7a,
					0xc7, 0x2e, 0x8a,
				},
			},
			// Example 1.2
			{
				[]byte{0x75, 0x0c, 0x40, 0x47, 0xf5, 0x47, 0xe8, 0xe4,
					0x14, 0x11, 0x85, 0x65, 0x23, 0x29, 0x8a, 0xc9, 0xba,
					0xe2, 0x45, 0xef, 0xaf, 0x13, 0x97, 0xfb, 0xe5, 0x6f,
					0x9d, 0xd5,
				},
				[]byte{0x0c, 0xc7, 0x42, 0xce, 0x4a, 0x9b, 0x7f, 0x32,
					0xf9, 0x51, 0xbc, 0xb2, 0x51, 0xef, 0xd9, 0x25, 0xfe,
					0x4f, 0xe3, 0x5f,
				},
				[]byte{0x64, 0x0d, 0xb1, 0xac, 0xc5, 0x8e, 0x05, 0x68,
					0xfe, 0x54, 0x07, 0xe5, 0xf9, 0xb7, 0x01, 0xdf, 0xf8,
					0xc3, 0xc9, 0x1e, 0x71, 0x6c, 0x53, 0x6f, 0xc7, 0xfc,
					0xec, 0x6c, 0xb5, 0xb7, 0x1c, 0x11, 0x65, 0x98, 0x8d,
					0x4a, 0x27, 0x9e, 0x15, 0x77, 0xd7, 0x30, 0xfc, 0x7a,
					0x29, 0x93, 0x2e, 0x3f, 0x00, 0xc8, 0x15, 0x15, 0x23,
					0x6d, 0x8d, 0x8e, 0x31, 0x01, 0x7a, 0x7a, 0x09, 0xdf,
					0x43, 0x52, 0xd9, 0x04, 0xcd, 0xeb, 0x79, 0xaa, 0x58,
					0x3a, 0xdc, 0xc3, 0x1e, 0xa6, 0x98, 0xa4, 0xc0, 0x52,
					0x83, 0xda, 0xba, 0x90, 0x89, 0xbe, 0x54, 0x91, 0xf6,
					0x7c, 0x1a, 0x4e, 0xe4, 0x8d, 0xc7, 0x4b, 0xbb, 0xe6,
					0x64, 0x3a, 0xef, 0x84, 0x66, 0x79, 0xb4, 0xcb, 0x39,
					0x5a, 0x35, 0x2d, 0x5e, 0xd1, 0x15, 0x91, 0x2d, 0xf6,
					0x96, 0xff, 0xe0, 0x70, 0x29, 0x32, 0x94, 0x6d, 0x71,
					0x49, 0x2b, 0x44,
				},
			},
			// Example 1.3
			{
				[]byte{0xd9, 0x4a, 0xe0, 0x83, 0x2e, 0x64, 0x45, 0xce,
					0x42, 0x33, 0x1c, 0xb0, 0x6d, 0x53, 0x1a, 0x82, 0xb1,
					0xdb, 0x4b, 0xaa, 0xd3, 0x0f, 0x74, 0x6d, 0xc9, 0x16,
					0xdf, 0x24, 0xd4, 0xe3, 0xc2, 0x45, 0x1f, 0xff, 0x59,
					0xa6, 0x42, 0x3e, 0xb0, 0xe1, 0xd0, 0x2d, 0x4f, 0xe6,
					0x46, 0xcf, 0x69, 0x9d, 0xfd, 0x81, 0x8c, 0x6e, 0x97,
					0xb0, 0x51,
				},
				[]byte{0x25, 0x14, 0xdf, 0x46, 0x95, 0x75, 0x5a, 0x67,
					0xb2, 0x88, 0xea, 0xf4, 0x90, 0x5c, 0x36, 0xee, 0xc6,
					0x6f, 0xd2, 0xfd,
				},
				[]byte{0x42, 0x37, 0x36, 0xed, 0x03, 0x5f, 0x60, 0x26,
					0xaf, 0x27, 0x6c, 0x35, 0xc0, 0xb3, 0x74, 0x1b, 0x36,
					0x5e, 0x5f, 0x76, 0xca, 0x09, 0x1b, 0x4e, 0x8c, 0x29,
					0xe2, 0xf0, 0xbe, 0xfe, 0xe6, 0x03, 0x59, 0x5a, 0xa8,
					0x32, 0x2d, 0x60, 0x2d, 0x2e, 0x62, 0x5e, 0x95, 0xeb,
					0x81, 0xb2, 0xf1, 0xc9, 0x72, 0x4e, 0x82, 0x2e, 0xca,
					0x76, 0xdb, 0x86, 0x18, 0xcf, 0x09, 0xc5, 0x34, 0x35,
					0x03, 0xa4, 0x36, 0x08, 0x35, 0xb5, 0x90, 0x3b, 0xc6,
					0x37, 0xe3, 0x87, 0x9f, 0xb0, 0x5e, 0x0e, 0xf3, 0x26,
					0x85, 0xd5, 0xae, 0xc5, 0x06, 0x7c, 0xd7, 0xcc, 0x96,
					0xfe, 0x4b, 0x26, 0x70, 0xb6, 0xea, 0xc3, 0x06, 0x6b,
					0x1f, 0xcf, 0x56, 0x86, 0xb6, 0x85, 0x89, 0xaa, 0xfb,
					0x7d, 0x62, 0x9b, 0x02, 0xd8, 0xf8, 0x62, 0x5c, 0xa3,
					0x83, 0x36, 0x24, 0xd4, 0x80, 0x0f, 0xb0, 0x81, 0xb1,
					0xcf, 0x94, 0xeb,
				},
			},
		},
	},
	// Key 10
	{"ae45ed5601cec6b8cc05f803935c674ddbe0d75c4c09fd7951fc6b0caec313a8df39970c518bffba5ed68f3f0d7f22a4029d413f1ae07e4ebe9e4177ce23e7f5404b569e4ee1bdcf3c1fb03ef113802d4f855eb9b5134b5a7c8085adcae6fa2fa1417ec3763be171b0c62b760ede23c12ad92b980884c641f5a8fac26bdad4a03381a22fe1b754885094c82506d4019a535a286afeb271bb9ba592de18dcf600c2aeeae56e02f7cf79fc14cf3bdc7cd84febbbf950ca90304b2219a7aa063aefa2c3c1980e560cd64afe779585b6107657b957857efde6010988ab7de417fc88d8f384c4e6e72c3f943e0c31c0c4a5cc36f879d8a3ac9d7d59860eaada6b83bb",
		65537,
		"056b04216fe5f354ac77250a4b6b0c8525a85c59b0bd80c56450a22d5f438e596a333aa875e291dd43f48cb88b9d5fc0d499f9fcd1c397f9afc070cd9e398c8d19e61db7c7410a6b2675dfbf5d345b804d201add502d5ce2dfcb091ce9997bbebe57306f383e4d588103f036f7e85d1934d152a323e4a8db451d6f4a5b1b0f102cc150e02feee2b88dea4ad4c1baccb24d84072d14e1d24a6771f7408ee30564fb86d4393a34bcf0b788501d193303f13a2284b001f0f649eaf79328d4ac5c430ab4414920a9460ed1b7bc40ec653e876d09abc509ae45b525190116a0c26101848298509c1c3bf3a483e7274054e15e97075036e989f60932807b5257751e79",
		[]testEncryptOAEPMessage{
			// Example 10.1
			{
				[]byte{0x8b, 0xba, 0x6b, 0xf8, 0x2a, 0x6c, 0x0f, 0x86,
					0xd5, 0xf1, 0x75, 0x6e, 0x97, 0x95, 0x68, 0x70, 0xb0,
					0x89, 0x53, 0xb0, 0x6b, 0x4e, 0xb2, 0x05, 0xbc, 0x16,
					0x94, 0xee,
				},
				[]byte{0x47, 0xe1, 0xab, 0x71, 0x19, 0xfe, 0xe5, 0x6c,
					0x95, 0xee, 0x5e, 0xaa, 0xd8, 0x6f, 0x40, 0xd0, 0xaa,
					0x63, 0xbd, 0x33,
				},
				[]byte{0x53, 0xea, 0x5d, 0xc0, 0x8c, 0xd2, 0x60, 0xfb,
					0x3b, 0x85, 0x85, 0x67, 0x28, 0x7f, 0xa9, 0x15, 0x52,
					0xc3, 0x0b, 0x2f, 0xeb, 0xfb, 0xa2, 0x13, 0xf0, 0xae,
					0x87, 0x70, 0x2d, 0x06, 0x8d, 0x19, 0xba, 0xb0, 0x7f,
					0xe5, 0x74, 0x52, 0x3d, 0xfb, 0x42, 0x13, 0x9d, 0x68,
					0xc3, 0xc5, 0xaf, 0xee, 0xe0, 0xbf, 0xe4, 0xcb, 0x79,
					0x69, 0xcb, 0xf3, 0x82, 0xb8, 0x04, 0xd6, 0xe6, 0x13,
					0x96, 0x14, 0x4e, 0x2d, 0x0e, 0x60, 0x74, 0x1f, 0x89,
					0x93, 0xc3, 0x01, 0x4b, 0x58, 0xb9, 0xb1, 0x95, 0x7a,
					0x8b, 0xab, 0xcd, 0x23, 0xaf, 0x85, 0x4f, 0x4c, 0x35,
					0x6f, 0xb1, 0x66, 0x2a, 0xa7, 0x2b, 0xfc, 0xc7, 0xe5,
					0x86, 0x55, 0x9d, 0xc4, 0x28, 0x0d, 0x16, 0x0c, 0x12,
					0x67, 0x85, 0xa7, 0x23, 0xeb, 0xee, 0xbe, 0xff, 0x71,
					0xf1, 0x15, 0x94, 0x44, 0x0a, 0xae, 0xf8, 0x7d, 0x10,
					0x79, 0x3a, 0x87, 0x74, 0xa2, 0x39, 0xd4, 0xa0, 0x4c,
					0x87, 0xfe, 0x14, 0x67, 0xb9, 0xda, 0xf8, 0x52, 0x08,
					0xec, 0x6c, 0x72, 0x55, 0x79, 0x4a, 0x96, 0xcc, 0x29,
					0x14, 0x2f, 0x9a, 0x8b, 0xd4, 0x18, 0xe3, 0xc1, 0xfd,
					0x67, 0x34, 0x4b, 0x0c, 0xd0, 0x82, 0x9d, 0xf3, 0xb2,
					0xbe, 0xc6, 0x02, 0x53, 0x19, 0x62, 0x93, 0xc6, 0xb3,
					0x4d, 0x3f, 0x75, 0xd3, 0x2f, 0x21, 0x3d, 0xd4, 0x5c,
					0x62, 0x73, 0xd5, 0x05, 0xad, 0xf4, 0xcc, 0xed, 0x10,
					0x57, 0xcb, 0x75, 0x8f, 0xc2, 0x6a, 0xee, 0xfa, 0x44,
					0x12, 0x55, 0xed, 0x4e, 0x64, 0xc1, 0x99, 0xee, 0x07,
					0x5e, 0x7f, 0x16, 0x64, 0x61, 0x82, 0xfd, 0xb4, 0x64,
					0x73, 0x9b, 0x68, 0xab, 0x5d, 0xaf, 0xf0, 0xe6, 0x3e,
					0x95, 0x52, 0x01, 0x68, 0x24, 0xf0, 0x54, 0xbf, 0x4d,
					0x3c, 0x8c, 0x90, 0xa9, 0x7b, 0xb6, 0xb6, 0x55, 0x32,
					0x84, 0xeb, 0x42, 0x9f, 0xcc,
				},
			},
		},
	},
}
