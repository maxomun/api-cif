package request

import "github.com/shopspring/decimal"

type PersonaRequest struct {
	Nombre     string          `json:"nombre"`
	ApPaterno  string          `json:"apPaterno"`
	ApMaterno  string          `json:"apMaterno"`
	Mail       string          `json:"mail"`
	NumCelular int             `json:"numCelular"`
	Edad       int             `json:"edad"`
	Renta      int             `json:"renta"`
	Impuesto   decimal.Decimal `json:"impuesto"`
}
