package utils

import (
	"crypto/rand"
	"math/big"
	"net"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0" // Fallback
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "0.0.0.0"
}

func GenerateToken() string {
	tokenLength := 32
	token := make([]byte, tokenLength)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range token {
		randInt, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			panic("Failed to generate token")
		}
		token[i] = charset[randInt.Int64()]
	}

	return string(token)
}
