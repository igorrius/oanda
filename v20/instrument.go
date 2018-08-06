package v20

import "time"

type CandlesRequest struct {
	// Name of the Instrument [required]
	Instrument string
	// The Price component(s) to get candlestick data for.
	// Can contain any combination of the characters “M” (midpoint candles) “B” (bid candles) and “A” (ask candles).
	// [default=M]
	Price string
	// The granularity of the candlesticks to fetch [default=S5].
	// Look CandlestickGranularity_X constants.
	granularity string
	// The number of candlesticks to return in the reponse.
	// Count should not be specified if both the start and end parameters are provided,
	// as the time range combined with the graularity will determine the number of candlesticks to return.
	// [default=500, maximum=5000]
	count int
	// The start of the time range to fetch candlesticks for.
	from time.Time
	// The end of the time range to fetch candlesticks for.
	to time.Time
	// A flag that controls whether the candlestick is “smoothed” or not.
	// A smoothed candlestick uses the previous candle’s close price as its open price,
	// while an unsmoothed candlestick uses the first price from its time range as its open price.
	// [default=False]
	smooth bool
	// A flag that controls whether the candlestick that is covered by the from time should be included in the results.
	// This flag enables clients to use the timestamp of the last completed candlestick received to poll
	// for future candlesticks but avoid receiving the previous candlestick repeatedly.
	// [default=True]
	includeFirst bool
	// The hour of the day (in the specified timezone) to use for granularities that have daily alignments.
	// [default=17, minimum=0, maximum=23]
	dailyAlignment bool
	// The timezone to use for the dailyAlignment parameter.
	// Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone.
	// Note that the returned times will still be represented in UTC.
	// [default=America/New_York]
	alignmentTimezone string
	// The day of the week used for granularities that have weekly alignment. [default=Friday]
	weeklyAlignment string
}

type CandlesResponse struct {
	Instrument  string `json:"instrument"`
	Granularity string `json:"granularity"`
	Candles     []struct {
		Complete bool      `json:"complete"`
		Volume   int       `json:"volume"`
		Time     time.Time `json:"time"`
		Mid      struct {
			O string `json:"o"`
			H string `json:"h"`
			L string `json:"l"`
			C string `json:"c"`
		} `json:"mid"`
	} `json:"candles"`
}
