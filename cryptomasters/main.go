package cryptomasters

import (
	"dw/go/crypto/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "LTC", "DOGE", "XRP"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyRate(currency)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyRate(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
