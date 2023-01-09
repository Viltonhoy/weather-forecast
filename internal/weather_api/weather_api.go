package weatherapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"weather-forecast/internal/generated"

	"go.uber.org/zap"
)

type ApiClient struct {
	apiKey   string
	Client   *http.Client
	Servicer Servicer
}

var ErrApi = errors.New("")

const (
	envApiKey = "API_KEY"
	mainCity  = "Moscow"
)

func New() *ApiClient {
	return &ApiClient{
		apiKey: os.Getenv(envApiKey),
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (a *ApiClient) CallAt(logger *zap.Logger, ctx context.Context, loc <-chan string, f func(*zap.Logger, string) (generated.WeatherResult, error)) error {
	logger.Debug("")

	go func() {
		var city = mainCity
		t := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-t.C:
				result, _ := f(logger, city)
				a.Servicer.AddWeatherInfo(logger, ctx, *result.Current, city)
			case newCity := <-loc:
				city = newCity
				logger.Debug("Change city to", zap.String("loc", city))
			}
		}
	}()
	return nil
}

func (a *ApiClient) WeatherRates(logger *zap.Logger, city string) (generated.WeatherResult, error) {
	logger.Debug("starting weather rates")

	var w *generated.WeatherResult

	url := fmt.Sprintf(`http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no`, a.apiKey, city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("bad request error", zap.Error(err))
		return generated.WeatherResult{}, err
	}

	res, err := a.Client.Do(req)
	if err != nil {
		return generated.WeatherResult{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return generated.WeatherResult{}, err
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		return generated.WeatherResult{}, err
	}

	if w.Error != nil {
		logger.Error(w.Error.Message, zap.Error(ErrApi))
		return generated.WeatherResult{}, ErrApi
	}

	fmt.Println(*w.Current, *w.Location)

	return *w, nil
}
