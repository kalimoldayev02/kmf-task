package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/kalimoldayev02/kmf-task/app/repository"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
)

var layout = "02-01-2006"

type CurrencyService struct {
	repository repository.Currency
}

func NewAuthSerive(r repository.Currency) *CurrencyService {
	return &CurrencyService{
		repository: r,
	}
}

func (s *CurrencyService) Save(date string) bool {
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("error parsing date: %s", err)
		return false
	}

	config := config.LoadConfig()

	baseUrl := fmt.Sprintf("https://%s/rss/get_rates.cfm", config.ServiceHosts.NotionalBank)
	query := url.Values{}
	query.Set("fdate", parsedDate.Format(layout))
	url := baseUrl + "?" + query.Encode()

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("error on get: %s", err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("error on get: %s", response.Status)
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error read from notional bank: %s", err)
		return false
	}

	fmt.Println(string(body))

	return false
}
