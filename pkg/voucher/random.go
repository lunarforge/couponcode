// + build !coupon-crypto

package voucher

import (
	"math/rand"
)

func randString(n int) string {
	var bytes = make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = symbols[rand.Int()%int(length)]
	}
	return string(bytes)
}
