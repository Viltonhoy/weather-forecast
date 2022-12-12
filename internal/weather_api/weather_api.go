package weatherapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

var ErrApi = errors.New("")

type ApiClient struct {
	apiKey string
	Client *http.Client
}

type weatherResult struct {
	Error    *errorInfo
	Location *locationInfo
	Current  *currentInfo
}

type errorInfo struct {
	Code    int64
	Message string
}

type locationInfo struct {
	Name            string
	Region          string
	Country         string
	Lat             float32
	Lon             float32
	Tz_id           string
	Localtime_epoch int64
	Localtime       string
}

type currentInfo struct {
	Last_updated_epoch int64
	Last_updated       string
	Temp_c             float32
	Temp_f             float32
	Is_day             int64
	Condition          *conditionInfo
	Wind_mph           float32
	Wind_kph           float32
	Wind_degree        int64
	Wind_dir           string
	Pressure_mb        float32
	Pressure_in        float32
	Precip_mm          float32
	Precip_in          float32
	Humidity           int64
	Cloud              int64
	Feelslike_c        float32
	Feelslike_f        float32
	Vis_km             float32
	Vis_miles          float32
	Uv                 float32
	Gust_mph           float32
	Gust_kph           float32
}

type conditionInfo struct {
	Text string
}

const envApiKey = "API_KEY"

func New() *ApiClient {
	return &ApiClient{
		apiKey: os.Getenv(envApiKey),
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (a *ApiClient) CallAt(logger *zap.Logger, loc chan string, f func(*zap.Logger, string) (weatherResult, error)) error {
	// loc, err := time.LoadLocation("Local")
	// if err != nil {
	// 	return err
	// }

	// now := time.Now().Local()
	//  firstCallTime := time.Date(
	// 	now.Year(),
	// 	now.Month(),
	// 	now.Day(),
	// 	hour,
	// 	min,
	// 	sec,
	// 	0,
	// 	loc,
	// )

	// firstCallTime := time.Now().Local()
	// if firstCallTime.Before(now) {
	// 	firstCallTime = firstCallTime.Add(time.Hour * 24)
	// }

	// duration := firstCallTime.Sub(time.Now().Local())

	// wg.Add(1)

	go func() {
		var city = "Moscow"
		t := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-t.C:
				f(logger, city)
			case newCity := <-loc:
				city = newCity
			}
		}
	}()
	// defer wg.Done()
	return nil
}

func (a *ApiClient) WeatherRates(logger *zap.Logger, city string) (weatherResult, error) {
	logger.Debug("starting weather rates")

	var w *weatherResult

	url := fmt.Sprintf(`http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no`, envApiKey, city)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logger.Error("bad request error", zap.Error(err))
		return weatherResult{}, err
	}
	res, err := a.Client.Do(req)
	if err != nil {
		return weatherResult{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weatherResult{}, err
	}
	err = json.Unmarshal(body, &w)
	if err != nil {
		return weatherResult{}, err
	}
	if w.Error != nil {
		logger.Error(w.Error.Message, zap.Error(ErrApi))
		return weatherResult{}, ErrApi
	}

	fmt.Println(*w.Current, *w.Location)

	return *w, nil
}

// func main() {
// 	logger, _ := zap.NewDevelopment()
// 	defer logger.Sync()

// 	// var wg sync.WaitGroup

// 	err := callAt(logger, 0, 0, 0, New().WeatherRates)
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}

// 	sigint := make(chan os.Signal, 1)
// 	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
// 	<-sigint

// 	w, err := New().WeatherRates(logger, "Moscow")
// 	if err != nil {
// 		return
// 	}

// 	fmt.Printf("%#v\n", w)
// }
