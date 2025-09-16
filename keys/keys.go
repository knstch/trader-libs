package keys

import "fmt"

const (
	OneMinute   = "1m"
	FiveMinutes = "5m"
	OneHour     = "1h"
	EightHours  = "8h"
)

func GetLastOiKey(symbol string) string {
	return fmt.Sprintf("mkt:last:oi:%s", symbol)
}

func GetLastFundingKey(symbol string) string {
	return fmt.Sprintf("mkt:last:funding:%s", symbol)
}

func GetFundingKey(symbol string, suffix string) string {
	return fmt.Sprintf("mkt:funding:%s:%s", suffix, symbol)
}

func GetOiKey(symbol string, suffix string) string {
	return fmt.Sprintf("mkt:oi:%s:%s", suffix, symbol)
}

func GetKlineKey(symbol string, suffix string) string {
	return fmt.Sprintf("mkt:kline:%s:%s", suffix, symbol)
}

func GetLastKlineKey(symbol string, suffix string) string {
	return fmt.Sprintf("mkt:last:kline:%s:%s", suffix, symbol)
}

func GetTsID(ts int64) string {
	return fmt.Sprintf("%d-0", ts)
}

func GetLiqRawKey(symbol string) string {
	return fmt.Sprintf("mkt:liqs:raw:%s", symbol)
}

func GetLiq1mKey(symbol string) string {
	return fmt.Sprintf("mkt:liqs:1m:%s", symbol)
}

func GetLastLiq1mKey(symbol string) string {
	return fmt.Sprintf("mkt:last:liqs:1m:%s", symbol)
}

func GetCoinGeckoKey() string {
	return "mkt:map:cg-token-id"
}
