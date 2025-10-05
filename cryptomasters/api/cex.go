package api

import (
	"dw/go/crypto/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("error: status code %d", res.StatusCode)
	}
	rate := datatypes.Rate{
		Currency: upCurrency,
		Price:    float64(response.Bid),
	}
	return &rate, nil
}
