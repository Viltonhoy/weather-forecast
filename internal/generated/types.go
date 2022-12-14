package generated

type WeatherResult struct {
	Error    *errorInfo
	Location *LocationInfo
	Current  *CurrentInfo
}

type errorInfo struct {
	Code    int64
	Message string
}

type LocationInfo struct {
	Name            string
	Region          string
	Country         string
	Lat             float32
	Lon             float32
	Tz_id           string
	Localtime_epoch int64
	Localtime       string
}

type CurrentInfo struct {
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
