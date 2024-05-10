package binance_connector

import (
	"context"
	"encoding/json"
)

const (
	miningEndpoint = "/sapi/v1/mining"
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
	TotalNum       int                 `json:"totalNum"`
	PageSize       int                 `json:"pageSize"`
}

type MiningPaymentListResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data AccountProfitsData `json:"data"`
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
		endpoint: miningEndpoint + "/payment/list",
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

type MiningWorkerListService struct {
	c            *Client
	algo         string
	userName     string
	pageIndex    int
	sort         int
	sortColumn   int
	workerStatus int
}

type MiningWorkerDataItem struct {
	WorkerID      string  `json:"workerId"`
	WorkerName    string  `json:"workerName"`
	Status        int     `json:"status"`
	DayHashRate   float64 `json:"dayHashRate"`
	HashRate      float64 `json:"hashRate"`
	RejectRate    float64 `json:"rejectRate"`
	LastShareTime int64   `json:"lastShareTime"`
}

type MiningWorkerData struct {
	WorkDatas []MiningWorkerDataItem `json:"workerDatas"`
	TotalNum  int                    `json:"totalNum"`
	PageSize  int                    `json:"pageSize"`
}

type MiningWorkerListResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data MiningWorkerData `json:"data"`
}

func (s *MiningWorkerListService) PageIndex(pageIndex int) *MiningWorkerListService {
	s.pageIndex = pageIndex
	return s
}

func (s *MiningWorkerListService) Algo(algo string) *MiningWorkerListService {
	s.algo = algo
	return s
}

func (s *MiningWorkerListService) UserName(userName string) *MiningWorkerListService {
	s.userName = userName
	return s
}

func (s *MiningWorkerListService) Sort(sort int) *MiningWorkerListService {
	s.sort = sort
	return s
}

func (s *MiningWorkerListService) SortColumn(sortColumn int) *MiningWorkerListService {
	s.sortColumn = sortColumn
	return s
}
func (s *MiningWorkerListService) WorkerStatus(workerStatus int) *MiningWorkerListService {
	s.workerStatus = workerStatus
	return s
}

func (s *MiningWorkerListService) Do(ctx context.Context, opts ...RequestOption) (res *MiningWorkerListResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: miningEndpoint + "/worker/list",
		secType:  secTypeSigned,
	}
	r.setParam("algo", s.algo)
	r.setParam("userName", s.userName)
	r.setParam("pageIndex", s.pageIndex)
	r.setParam("sort", s.sort)
	r.setParam("sortColumn", s.sortColumn)
	r.setParam("workerStatus", s.workerStatus)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MiningWorkerListResponse{}, err
	}

	res = new(MiningWorkerListResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return &MiningWorkerListResponse{}, err
	}
	return res, nil
}

type MiningStatisticsUserStatusService struct {
	c        *Client
	algo     string
	userName string
}

type MiningStatisticsUserStatusData struct {
	FifteenMinHashRate string `json:"fifteenMinHashRate"`
	DayHashRate        string `json:"dayHashRate"`
	ValidNum           int    `json:"validNum"`
	InvalidNum         int    `json:"invalidNum"`
	ProfitToday        struct {
		Btc string `json:"BTC"`
		Bsv string `json:"BSV"`
		Bch string `json:"BCH"`
	} `json:"profitToday"`
	ProfitYesterday struct {
		Btc string `json:"BTC"`
		Bsv string `json:"BSV"`
		Bch string `json:"BCH"`
	} `json:"profitYesterday"`
	UserName string `json:"userName"`
	Unit     string `json:"unit"`
	Algo     string `json:"algo"`
}

type MiningStatisticsUserStatusResponse struct {
	Code     int                            `json:"code"`
	Msg      string                         `json:"msg"`
	Data     MiningStatisticsUserStatusData `json:"data"`
	TotalNum int                            `json:"totalNum"`
	PageSize int                            `json:"pageSize"`
}

func (s *MiningStatisticsUserStatusService) Algo(algo string) *MiningStatisticsUserStatusService {
	s.algo = algo
	return s
}

func (s *MiningStatisticsUserStatusService) UserName(userName string) *MiningStatisticsUserStatusService {
	s.userName = userName
	return s
}

func (s *MiningStatisticsUserStatusService) Do(ctx context.Context, opts ...RequestOption) (res *MiningStatisticsUserStatusResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: miningEndpoint + "/statistics/user/status",
		secType:  secTypeSigned,
	}
	r.setParam("algo", s.algo)
	r.setParam("userName", s.userName)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MiningStatisticsUserStatusResponse{}, err
	}

	res = new(MiningStatisticsUserStatusResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return &MiningStatisticsUserStatusResponse{}, err
	}
	return res, nil
}
