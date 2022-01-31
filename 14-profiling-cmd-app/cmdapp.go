package main

import (
	"bytes"
	"math"
	"net/http"
	_ "net/http/pprof" // gathering the profile data
	"strconv"
	"sync"
)

func isPrime(no int) bool {
	end := int(math.Sqrt(float64(no)))
	for i := 2; i < end; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func generatePrimes() *[]int {
	start := 2
	end := 100
	//primes := []int{}
	primes := make([]int, 0, 100)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return &primes
}

/*
func formatPrimes(primeNos []int) string {
	str := ""
	for _, no := range primeNos {
		str += strconv.Itoa(no) + "-"
	}
	return str
}
*/

var pool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func formatPrimes(primeNos *[]int) *[]byte {
	result := pool.Get().(*bytes.Buffer)
	for _, primeNo := range *primeNos {
		result.Write([]byte(strconv.Itoa(primeNo)))
		result.WriteRune('-')
	}
	resultBytes := result.Bytes()
	result.Reset()
	pool.Put(result)
	return &resultBytes
}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	formatPrimes(generatePrimes())
}
