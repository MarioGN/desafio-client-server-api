package server

import "errors"

type APIResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type ServiceResponse struct {
	Bid float64 `json:"bid"`
}

type USDBRL struct {
	Code       string  `json:"code"`
	Codein     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high,string"`
	Low        float64 `json:"low,string"`
	VarBid     float64 `json:"varBid,string"`
	PctChange  float64 `json:"pctChange,string"`
	Bid        float64 `json:"bid,string"`
	Ask        float64 `json:"ask,string"`
	Timestamp  int     `json:"timestamp,string"`
	CreateDate string  `json:"create_date"`
}

func (u *USDBRL) Validate() error {
	if u.Bid <= 0 {
		return errors.New("invalid bid")
	}
	return nil
}
