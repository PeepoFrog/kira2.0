package types

// Transaction structs for `sekaid query tx`

type EventAttributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
}

type Events struct {
	Type       string            `json:"type"`
	Attributes []EventAttributes `json:"attributes"`
}

type LogAttributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Logs struct {
	Type       string          `json:"type"`
	Attributes []LogAttributes `json:"attributes"`
}

type Message struct {
	Type       string `json:"@type"`
	Proposer   string `json:"proposer"`
	Address    string `json:"address"`
	Permission int    `json:"permission"`
}

type Body struct {
	Messages []Message `json:"messages"`
	Memo     string    `json:"memo"`
}

type Tx struct {
	Type string `json:"@type"`
	Body Body   `json:"body"`
}

type TxData struct {
	Height    string      `json:"height"`
	Txhash    string      `json:"txhash"`
	Code      int         `json:"code"`
	Data      string      `json:"data"`
	RawLog    string      `json:"raw_log"`
	Logs      interface{} `json:"logs"` // no data provided, assuming a slice of empty interface
	Info      string      `json:"info"`
	GasWanted string      `json:"gas_wanted"`
	GasUsed   string      `json:"gas_used"`
	Tx        interface{} `json:"tx"` // no data provided, using an empty interface
	Timestamp string      `json:"timestamp"`
	Events    interface{} `json:"events"` // no data provided, assuming a slice of empty interface
}
