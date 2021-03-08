// + build coupon-crypto

package voucher

import (
	"crypto/rand"
)

func randStringCrypto(n int) string {
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = symbols[b%byte(length)]
	}
	return string(bytes)
}
