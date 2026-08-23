package f12025 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type WeatherForecastSample struct {
	SessionType            uint8 // 0 = unknown, see appendix
	TimeOffset             uint8 // Time in minutes the forecast is for
	Weather                uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature       int8  // Track temp. in degrees Celsius
	TrackTemperatureChange int8  // Track temp. change – 0 = up, 1 = down, 2 = no change
	AirTemperature         int8  // Air temp. in degrees celsius
	AirTemperatureChange   int8  // Air temp. change – 0 = up, 1 = down, 2 = no change
	RainPercentage         uint8 // Percentage chance of rain (0-100)
}

// GetSessionType returns the SessionType of *WeatherForecastSample
func (data *WeatherForecastSample) GetSessionType() uint8 { return data.SessionType }

// SetSessionType stores the SessionType of *WeatherForecastSample
func (data *WeatherForecastSample) SetSessionType(v uint8) { data.SessionType = v }

// GetTimeOffset returns the TimeOffset of *WeatherForecastSample
func (data *WeatherForecastSample) GetTimeOffset() uint8 { return data.TimeOffset }

// SetTimeOffset stores the TimeOffset of *WeatherForecastSample
func (data *WeatherForecastSample) SetTimeOffset(v uint8) { data.TimeOffset = v }

// GetWeather returns the Weather of *WeatherForecastSample
func (data *WeatherForecastSample) GetWeather() uint8 { return data.Weather }

// SetWeather stores the Weather of *WeatherForecastSample
func (data *WeatherForecastSample) SetWeather(v uint8) { data.Weather = v }

// GetTrackTemperature returns the TrackTemperature of *WeatherForecastSample
func (data *WeatherForecastSample) GetTrackTemperature() int8 { return data.TrackTemperature }

// SetTrackTemperature stores the TrackTemperature of *WeatherForecastSample
func (data *WeatherForecastSample) SetTrackTemperature(v int8) { data.TrackTemperature = v }

// GetTrackTemperatureChange returns the TrackTemperatureChange of *WeatherForecastSample
func (data *WeatherForecastSample) GetTrackTemperatureChange() int8 {
	return data.TrackTemperatureChange
}

// SetTrackTemperatureChange stores the TrackTemperatureChange of *WeatherForecastSample
func (data *WeatherForecastSample) SetTrackTemperatureChange(v int8) { data.TrackTemperatureChange = v }

// GetAirTemperature returns the AirTemperature of *WeatherForecastSample
func (data *WeatherForecastSample) GetAirTemperature() int8 { return data.AirTemperature }

// SetAirTemperature stores the AirTemperature of *WeatherForecastSample
func (data *WeatherForecastSample) SetAirTemperature(v int8) { data.AirTemperature = v }

// GetAirTemperatureChange returns the AirTemperatureChange of *WeatherForecastSample
func (data *WeatherForecastSample) GetAirTemperatureChange() int8 { return data.AirTemperatureChange }

// SetAirTemperatureChange stores the AirTemperatureChange of *WeatherForecastSample
func (data *WeatherForecastSample) SetAirTemperatureChange(v int8) { data.AirTemperatureChange = v }

// GetRainPercentage returns the RainPercentage of *WeatherForecastSample
func (data *WeatherForecastSample) GetRainPercentage() uint8 { return data.RainPercentage }

// SetRainPercentage stores the RainPercentage of *WeatherForecastSample
func (data *WeatherForecastSample) SetRainPercentage(v uint8) { data.RainPercentage = v }

func (data *WeatherForecastSample) Parse(reader *xbinary.LittleEndianReader) {
	data.SessionType = reader.ReadUint8()
	data.TimeOffset = reader.ReadUint8()
	data.Weather = reader.ReadUint8()
	data.TrackTemperature = reader.ReadInt8()
	data.TrackTemperatureChange = reader.ReadInt8()
	data.AirTemperature = reader.ReadInt8()
	data.AirTemperatureChange = reader.ReadInt8()
	data.RainPercentage = reader.ReadUint8()

}
