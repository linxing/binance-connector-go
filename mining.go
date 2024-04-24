package binance_connector

import (
	"context"
	"encoding/json"
)

const (
	miningEndpoint = "/sapi/v1/mining/payment/list"
)

type MiningPaymentListService struct {
	c         *Client
	algo      string
	userName  string
	coin      string
	pageIndex int
	pageSize  int
}

type MiningPaymentItem struct {
	Time           int64   `json:"time"`
	Type           int     `json:"type"`
	HashTransfer   int64   `json:"hashTransfer"`
	TransferAmount float64 `json:"transferAmount"`
	DayHashRate    float64 `json:"dayHashRate"`
	ProfitAmount   float64 `json:"profitAmount"`
	CoinName       string  `json:"coinName"`
	Status         int     `json:"status"`
}

type AccountProfitsData struct {
	AccountProfits []MiningPaymentItem `json:"accountProfits"`
}

type MiningPaymentListResponse struct {
	Code     int                `json:"code"`
	Msg      string             `json:"msg"`
	Data     AccountProfitsData `json:"data"`
	TotalNum int                `json:"totalNum"`
	PageSize int                `json:"pageSize"`
}

func (s *MiningPaymentListService) PageIndex(pageIndex int) *MiningPaymentListService {
	s.pageIndex = pageIndex
	return s
}

func (s *MiningPaymentListService) PageSize(pageSize int) *MiningPaymentListService {
	s.pageSize = pageSize
	return s
}

func (s *MiningPaymentListService) Algo(algo string) *MiningPaymentListService {
	s.algo = algo
	return s
}

func (s *MiningPaymentListService) UserName(userName string) *MiningPaymentListService {
	s.userName = userName
	return s
}

func (s *MiningPaymentListService) Coin(coin string) *MiningPaymentListService {
	s.coin = coin
	return s
}

func (s *MiningPaymentListService) Do(ctx context.Context, opts ...RequestOption) (res *MiningPaymentListResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: miningEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("algo", s.algo)
	r.setParam("userName", s.userName)
	r.setParam("coin", s.coin)
	r.setParam("pageSize", s.pageSize)
	r.setParam("pageIndex", s.pageIndex)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MiningPaymentListResponse{}, err
	}

	res = new(MiningPaymentListResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return &MiningPaymentListResponse{}, err
	}
	return res, nil
}
