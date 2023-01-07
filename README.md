# :rocket: Token-based authentication using JWT and PASETO
## A completely pluggable, well tested and documented Go package for token-based authentication.

---

# Features:
- [x] `JWT` tokens generation and verification.
- [x] `PASETO` tokens generation and verification.
- [x] Protected from using the weak/broken algorithms (ex. `ECDSA`, `RSA` with `PKCSv1.5`) incase of using `JWT`.
- [x] `Oracle` attack protection.
- [x] `Invalid-Curve` attack protection.
- [x] `Trivial Forgery` attack protection.
- [x] Strong algorithms.

---

# Examples
## PASETO Token
```go
package main

import (
	"github.com/dsha256/token-go/token"
	"log"
	"time"
)

func main() {
	maker, err := token.NewPasetoMaker("KEY MUST BE EXACTLY 32 CHARS....")
	if err != nil {
		log.Fatalf("Failed to create PASETO maker: %v", err)
	}

	username := "YOUR USERNAME HERE"
	// Token duration
	duration := time.Minute

	tkn, payload, err := maker.CreateToken(username, duration)
	if err != nil {
		log.Fatalf("Failed to create PASETO tkn: %v", err)
	}

	log.Println(tkn) // Output #1
	log.Println(payload) // Output #2

	payload, err = maker.VerifyToken(tkn)
	if err != nil {
		log.Fatalf("Invalid token: %v", err)
	}
	log.Println(payload) // Output #3
}
```

Output #1
```txt
v2.local.rrfTw0gcJH_0iwaPeVUsqgLdrwf4wDfhaPVHBIyB3i8-8-GP-a1XOPhNkrY3yEZnZkJebiNyivGaRdZvpyZsUzPJzArJCwz3rt2dmiCDPtyBstlZ5fTBYKU7dN51aRl5zHR6cXnvmHsK21KTBZ3BB65nc5eBuRXAHz9IKB3TIwjYeWQk6cu9nutAfiEZ8oQ2la4Q-pruIRYv_1niwilmwZWQVTnUmEiBhjoRAlkcilp-buDycHs_LwwS1bVt_lEbmhgyKc3CYFJSxg-q6tKD6ZkU0k-Dw84e.bnVsbA
```

Output #2
```txt
&{669c9c24-44c2-4a2b-9005-dbc51457c16b YOUR USERNAME HERE 2023-01-07 14:47:36.55678518 +0400 +04 m=+0.000188001 2023-01-07 14:48:36.55678518 +0400 +04 m=+60.000188101}
```

Output #3
```txt
&{669c9c24-44c2-4a2b-9005-dbc51457c16b YOUR USERNAME HERE 2023-01-07 14:47:36.55678518 +0400 +04 2023-01-07 14:48:36.55678518 +0400 +04}
```

## JWT Token
```go
package main

import (
	"github.com/dsha256/token-go/token"
	"log"
	"time"
)

func main() {
	maker, err := token.NewJwtMaker("KEY MUST BE EXACTLY 32 CHARS....")
	if err != nil {
		log.Fatalf("Failed to create PASETO maker: %v", err)
	}

	username := "YOUR USERNAME HERE"
	// Token duration
	duration := time.Minute

	tkn, payload, err := maker.CreateToken(username, duration)
	if err != nil {
		log.Fatalf("Failed to create PASETO tkn: %v", err)
	}

	log.Println(tkn)
	log.Println(payload)

	payload, err = maker.VerifyToken(tkn)
	if err != nil {
		log.Fatalf("Invalid token: %v", err)
	}
	log.Println(payload)
}
```

Output #1
```txt
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImRkN2MzMjYxLTc2MWQtNGRkYi1hYTcxLTUxNTU3MzEyNjdiNyIsInVzZXJuYW1lIjoiWU9VUiBVU0VSTkFNRSBIRVJFIiwiaXNzdWVkX2F0IjoiMjAy
My0wMS0wN1QxNTowNjo0OS4zMDA0NjQ2ODcrMDQ6MDAiLCJleHBpcmVkX2F0IjoiMjAyMy0wMS0wN1QxNTowNzo0OS4zMDA0NjQ3ODcrMDQ6MDAifQ.IYJFRd7yhxoZbXuis2ptDdjidhY4SMzIzCMvMxiZboE
```

Output #2
```txt
&{dd7c3261-761d-4ddb-aa71-5155731267b7 YOUR USERNAME HERE 2023-01-07 15:06:49.300464687 +0400 +04 m=+0.016301401 2023-01-07 15:07:49.300464787 +0400 +04 m=+60.01
6301401}
```

Output #3
```txt
&{dd7c3261-761d-4ddb-aa71-5155731267b7 YOUR USERNAME HERE 2023-01-07 15:06:49.300464687 +0400 +04 2023-01-07 15:07:49.300464787 +0400 +04}
```
