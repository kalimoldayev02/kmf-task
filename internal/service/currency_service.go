package service

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/kalimoldayev02/kmf-task/internal/repository"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
)

var layout = "02-01-2006"

type CurrencyService struct {
	repository repository.Currency
}

type RatesRequest struct {
	Currency []struct {
		FullName    string `xml:"fullname"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Quant       int    `xml:"quant"`
		Index       string `xml:"index"`
	} `xml:"item"`
}

func NewAuthService(r repository.Currency) *CurrencyService {
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
	config, err := config.NewCoifig()
	if err != nil {
		log.Fatalf("error: %s", err)
		return false
	}

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

	rates := new(RatesRequest)
	if err := xml.Unmarshal([]byte(body), rates); err != nil {
		log.Fatalf("error read from notional bank: %s", err)
		return false
	}

	for _, currency := range rates.Currency {
		fmt.Printf("Fullname: %s\n", currency.FullName)
		fmt.Printf("---\n")
	}

	return false
}
