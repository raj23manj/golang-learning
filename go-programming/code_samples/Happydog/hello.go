package hello

// https://go.dev/blog/using-go-modules
import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}
