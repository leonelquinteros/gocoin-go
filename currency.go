package gocoin

// CurrencyDetail type
type CurrencyDetail struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Symbol    string   `json:"symbol"`
	IsCrypto  bool     `json:"is_crypto"`
	SortOrder int      `json:"sort_order"`
	Precision int      `json:"precision"`
	Aliases   []string `json:"aliases"`
}
