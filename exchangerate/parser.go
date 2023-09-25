package exchangerate

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

const userAgent = "Mozilla/5.0 (Windows NT 6.3; rv:36.0) Gecko/20100101 Firefox/36.0"

const fetchUrl string = "https://freecurrencyrates.com/en/%s-exchange-rate-calculator"
const htmlRateSelector string = "input#rate-iso-%s"

func getDocument(link string) (*html.Node, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", link, nil)
	request.Header.Add("User-Agent", userAgent)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if 200 <= response.StatusCode && response.StatusCode < 300 {
		return html.Parse(response.Body)
	}
	return nil, fmt.Errorf("status code: %d", response.StatusCode)
}

func getExchangeRate(doc *html.Node, iso string) (float64, error) {
	selector, _ := css.Parse(fmt.Sprintf(htmlRateSelector, iso))
	matches := selector.Select(doc)
	if len(matches) == 0 {
		return 0, fmt.Errorf("unknown currency: %s", iso)
	}
	match := matches[0]
	for _, attr := range match.Attr {
		if attr.Key == "value" {
			return strconv.ParseFloat(attr.Val, 64)
		}
	}
	return 0, fmt.Errorf("attribute not found")
}

// CurrencyExists checks whether provided currency exists via a HTTP request.
func CurrencyExists(iso string) bool {
	url := fmt.Sprintf(fetchUrl, strings.ToUpper(iso))
	_, err := getDocument(url)
	return err == nil
}

// GetExchangeRate returns an exchange rate from one currency to another.
// Formula is as follows: 1 <isoFrom> = <rate> <isoTo>
func GetExchangeRate(isoFrom, isoTo string) (*ExchangeRate, error) {
	rate := ExchangeRate{From: strings.ToUpper(isoFrom), To: strings.ToUpper(isoTo)}
	url := fmt.Sprintf(fetchUrl, rate.From)
	document, err := getDocument(url)
	if err != nil {
		return nil, err
	}
	rateValue, err := getExchangeRate(document, rate.To)
	if err != nil {
		return nil, err
	}
	rate.Rate = rateValue
	return &rate, nil
}
