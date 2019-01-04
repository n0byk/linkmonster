/*
package main

import (
        "fmt"

        "github.com/omar-h/csrf"
)

func main() {
        const secret = "erHUnxuhBMRIsVB1LfqmiWCgB83ZEerH"
        CSRF := csrf.New(csrf.Config{
                // Secret should persist over program restart.
                Secret: secret,
                SaltLen: 16,
        })

        salt := CSRF.GenerateSalt()
        token := CSRF.GenerateToken(salt)

        // Print the secret, a random salt and the token generated from them.
        fmt.Println("Secret: ", secret)
        fmt.Println("Salt: ", salt)
        fmt.Println("Token: ", token)

        // Returns true
        CSRF.Verify(token)
}
*/

package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

// Toolset holds the configuration and all the utility tools.
// The utility tools use this to access the secret and saltLen.
type Toolset struct {
	c Config
}

// Config is the configuration for the package and toolset.
type Config struct {
	Secret  string
	SaltLen int
}

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	idxBits = 6
	idxMask = 1<<idxBits - 1
)

// GenerateSalt generates a random string of specified length.
func (f *Toolset) GenerateSalt() string {
	result := make([]byte, f.c.SaltLen)
	bufferSize := int(float64(f.c.SaltLen) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < f.c.SaltLen; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%f.c.SaltLen] & idxMask); idx < len(chars) {
			result[i] = chars[idx]
			i++
		}
	}

	return string(result)
}

// GenerateToken generates a secure token from a secret and salt.
func (f *Toolset) GenerateToken(salt string) string {
	return salt + hash(salt+"-"+f.c.Secret)
}

// Verify verifies if a token is valid.
// It takes in the salt length and secret used to create the token.
func (f *Toolset) Verify(token string) bool {
	salt := token[0:f.c.SaltLen]
	return salt+hash(salt+"-"+f.c.Secret) == token
}

// New returns a new Toolset, it takes in a type Options.
// The toolset will use the options.
func New(c Config) *Toolset {
	return &Toolset{
		c: c,
	}
}

// hash hashes a string using sha256 and returns a
// base64 encoded string, which is URL safe.
func hash(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)

	hashedString := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
	return hashedString
}

// secureRandomBytes generates secure random bytes of a specified
// length.
func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	rand.Read(randomBytes)
	return randomBytes
}
