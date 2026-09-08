package f12025 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type WeatherForecastSample

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
