package tdex

import "fmt"

type Tdex interface {
	ProductCoins() (*ProductCoins, error)

	UserInfo() (*UserInfo, error)

	WalletBalances(bqs WalletBalancesRequest) (*WalletBalances, error)
	WalletBalance(bq WalletBalanceRequest) (*WalletBalance, error)
	WalletWithdraw(wq WalletWithdrawRequest) (*WalletWithdraw, error)
	WalletSwitch(sq WalletSwitchRequest) (*WalletSwitch, error)

	FuturesOpen(fr FuturesOpenRequest) (*FuturesOpen, error)
	FuturesClose(fcr []FuturesCloseRequest) (*FuturesClose, error)
	FuturesCloseAll(fcar []FuturesCloseAllRequest) (*FuturesCloseAll, error)
	FuturesCancel(fcr []FuturesCancelRequest) (*FuturesCancel, error)
	FuturesSetsl(sr FuturesSetslRequest) (*FuturesSetsl, error)
	FuturesSettp(sr FuturesSettpRequest) (*FuturesSettp, error)
	FuturesMerge(mr FuturesMergeRequest) (*FuturesMerge, error)
	FuturesSplit(sr FuturesSplitRequest) (*FuturesSplit, error)
	FuturesSetup(sr FuturesSetupRequest) (*FuturesSetup, error)
	FuturesSchemeSet(ssr FuturesSchemeSetRequest) (*FuturesSchemeSet, error)
	FuturesSchemeGet(sgr FuturesSchemeGetRequest) (*FuturesSchemeGet, error)
	FuturesOrders() (*FuturesOrders, error)
	FuturesPosition() (*FuturesPosition, error)
	FuturesHistory(ghr FuturesHistoryRequest) (*FuturesHistory, error)
	FuturesContract(gcr FuturesContractRequest) (*FuturesContract, error)

	SpotBuy(sbr SpotBuyRequest) (*SpotBuy, error)
	SpotSell(ssr SpotSellRequest) (*SpotSell, error)
	SpotCancel(scr SpotCancelRequest) (*SpotCancel, error)
	SpotOrders() (*SpotOrders, error)
	SpotHistory(shr SpotHistoryRequest) (*SpotHistory, error)
}
type tdex struct {
	Service Service
}

// Error represents Binance error structure with error code and message.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// Error returns formatted error message.
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
func NewTdex(service Service) Tdex {
	return &tdex{
		Service: service,
	}
}

type Data struct {
}
type ProductCoins struct {
	Status uint32
	Data   ProductCoinsData
	Errmsg string
}
type ProductCoinsData struct {
	List []*ProductCoinsList
}
type ProductCoinsList struct {
	ID            int64
	Code          string
	Symbol        string
	Family        string
	MinTrade      float64
	MinCash       float64
	Fee           float64
	FeeType       string
	Digits        uint32
	ShowDigits    uint32
	Confirmations uint32
	Op            uint32
	Types         uint32
	Exchange      ProductCoinsExchange
}
type ProductCoinsExchange struct {
	Max uint32
	Min float64
	To  []uint32
}

func (t *tdex) ProductCoins() (*ProductCoins, error) {
	return t.Service.ProductCoins()
}

type UserInfo struct {
	Status uint32
	Data   UserInfoData
	Errmsg string
}
type UserInfoData struct {
	AccountID     int64
	UID           string
	Vip           uint32
	GgAuth        bool
	Email         string
	Name          string
	Mobile        string
	Lang          string
	LastLoginTime int64
	LastLoginIP   string
}

func (t *tdex) UserInfo() (*UserInfo, error) {
	return t.Service.UserInfo()
}

type WalletBalances struct {
	Status uint32
	Data   WalletBalancesData
	Errmsg string
}
type WalletBalancesData struct {
	List []*WalletBalancesList
}
type WalletBalancesList struct {
	Type      uint32
	Currency  uint32
	Quantity  float64
	Withdraw  float64
	Transfer  float64
	Deposit   float64
	Lock      uint32
	Available uint32
	Status    uint32
}

func (t *tdex) WalletBalances(bq WalletBalancesRequest) (*WalletBalances, error) {
	return t.Service.WalletBalances(bq)
}

type WalletBalance struct {
	Status uint32
	Data   WalletBalanceData
	Errmsg string
}
type WalletBalanceData struct {
	Type      uint32
	Currency  uint32
	Quantity  float64
	Withdraw  float64
	Transfer  float64
	Deposit   float64
	Lock      uint32
	Available uint32
	Status    uint32
}

func (t *tdex) WalletBalance(bq WalletBalanceRequest) (*WalletBalance, error) {
	return t.Service.WalletBalance(bq)
}

type WalletWithdraw struct {
	Status uint32
	Data   WithdrawData
	Errmsg string
}
type WithdrawData struct {
	Currency uint32
	Quantity float64
	Withdraw float64
	Transfer float64
	Deposit  float64
	OrderID  string
}

func (t *tdex) WalletWithdraw(wq WalletWithdrawRequest) (*WalletWithdraw, error) {
	return t.Service.WalletWithdraw(wq)
}

type WalletSwitch struct {
	Status uint32
	Data   SwitchData
	Errmsg string
}
type SwitchData struct {
	Currency  uint32
	Quantity  float64
	Unconfirm float64
	Withdraw  float64
	Transfer  float64
	Deposit   float64
	OrderId   string
}

func (t *tdex) WalletSwitch(wq WalletSwitchRequest) (*WalletSwitch, error) {
	return t.Service.WalletSwitch(wq)
}

type FuturesOpen struct {
	Status uint32
	Data   FuturesOpenData
	Errmsg string
}
type FuturesOpenData struct {
}

func (t *tdex) FuturesOpen(wq FuturesOpenRequest) (*FuturesOpen, error) {
	return t.Service.FuturesOpen(wq)
}

type FuturesClose struct {
	Status uint32
	Data   FuturesCloseData
	Errmsg string
}
type FuturesCloseData struct {
	Result FuturesCloseResult
}
type FuturesCloseResult struct {
	Code  uint32
	Error string
}

func (t *tdex) FuturesClose(wq []FuturesCloseRequest) (*FuturesClose, error) {
	return t.Service.FuturesClose(wq)
}

type FuturesCloseAll struct {
	Status uint32
	Data   FuturesCloseAllData
	Errmsg string
}
type FuturesCloseAllData struct {
}

func (t *tdex) FuturesCloseAll(wq []FuturesCloseAllRequest) (*FuturesCloseAll, error) {
	return t.Service.FuturesCloseAll(wq)
}

type FuturesCancel struct {
	Status uint32
	Data   FuturesCancelData
	Errmsg string
}
type FuturesCancelData struct {
}

func (t *tdex) FuturesCancel(wq []FuturesCancelRequest) (*FuturesCancel, error) {
	return t.Service.FuturesCancel(wq)
}

type FuturesSetsl struct {
	Status uint32
	Data   SetslData
	Errmsg string
}
type SetslData struct {
}

func (t *tdex) FuturesSetsl(wq FuturesSetslRequest) (*FuturesSetsl, error) {
	return t.Service.FuturesSetsl(wq)
}

type FuturesSettp struct {
	Status uint32
	Data   SettpData
	Errmsg string
}
type SettpData struct {
}

func (t *tdex) FuturesSettp(wq FuturesSettpRequest) (*FuturesSettp, error) {
	return t.Service.FuturesSettp(wq)
}

type FuturesMerge struct {
	Status uint32
	Data   MergeData
	Errmsg string
}
type MergeData struct {
}

func (t *tdex) FuturesMerge(wq FuturesMergeRequest) (*FuturesMerge, error) {
	return t.Service.FuturesMerge(wq)
}

type FuturesSplit struct {
	Status uint32
	Data   SplitData
	Errmsg string
}
type SplitData struct {
}

func (t *tdex) FuturesSplit(wq FuturesSplitRequest) (*FuturesSplit, error) {
	return t.Service.FuturesSplit(wq)
}

type FuturesSetup struct {
	Status uint32
	Data   SetupData
	Errmsg string
}
type SetupData struct {
}

func (t *tdex) FuturesSetup(wq FuturesSetupRequest) (*FuturesSetup, error) {
	return t.Service.FuturesSetup(wq)
}

type FuturesSchemeGet struct {
	Status uint32
	Data   SchemeGetData
	Errmsg string
}
type SchemeGetData struct {
	Shared bool
	Merged bool
}

func (t *tdex) FuturesSchemeGet(wq FuturesSchemeGetRequest) (*FuturesSchemeGet, error) {
	return t.Service.FuturesSchemeGet(wq)
}

type FuturesSchemeSet struct {
	Status uint32
	Data   SchemeSetData
	Errmsg string
}
type SchemeSetData struct {
}

func (t *tdex) FuturesSchemeSet(wq FuturesSchemeSetRequest) (*FuturesSchemeSet, error) {
	return t.Service.FuturesSchemeSet(wq)
}

type FuturesOrders struct {
	Status uint32
	Data   GetOrdersData
	Errmsg string
}
type GetOrdersData struct {
	List []GetOrdersList
}
type GetOrdersList struct {
	ID          int64
	Pid         int64
	Cid         uint32
	Rid         int64
	Side        uint32
	Kind        uint32
	Attempt     uint32
	Scale       float64
	Volume      uint32
	Visible     uint32
	Filled      uint32
	Distance    uint32
	Price       float64
	Shared      uint32
	Timely      uint32
	TimelyParam uint32
	Deadline    uint32
	Passive     uint32
	Better      uint32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Relative    uint32
	Vertex      int64
	Activated   int64
	Triggered   int64
	State       uint32
	SerialID    uint32
	Notes       string
	CreatedAt   string
	UpdatedAt   string
}

func (t *tdex) FuturesOrders() (*FuturesOrders, error) {
	return t.Service.FuturesOrders()
}

type FuturesPosition struct {
	Status uint32
	Data   GetPositionData
	Errmsg string
}
type GetPositionData struct {
	List []GetPositionList
}
type GetPositionList struct {
	ID        int64
	Cid       uint32
	Shared    uint32
	Side      uint32
	Price     float64
	Volume    uint32
	Scale     float64
	Margin    float64
	Repay     float64
	Force     float64
	Fee       float64
	Notes     string
	Completed uint32
	Bankrupt  uint32
	SerialID  uint32
	Index     uint32
	CreatedAt string
	UpdatedAt string
}

func (t *tdex) FuturesPosition() (*FuturesPosition, error) {
	return t.Service.FuturesPosition()
}

type FuturesHistory struct {
	Status uint32
	Data   GetHistoryData
	Errmsg string
}
type GetHistoryData struct {
	List      []GetHistoryList
	Total     int32
	PageCount int32
	PageSize  int32
	Page      int32
}
type GetHistoryList struct {
	ID          int64
	Lid         int64
	Cid         uint32
	Pid         int64
	Rid         int64
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Relative    uint32
	Vertex      int64
	Attempt     uint32
	Side        uint32
	Kind        uint32
	Scale       float64
	Volume      uint32
	Filled      uint32
	Distance    uint32
	Price       float64
	Shared      uint32
	Timely      uint32
	TimelyParam uint32
	Deadline    uint32
	Passive     uint32
	Visible     uint32
	Better      uint32
	Activated   int64
	Triggered   int64
	State       uint32
	SerialID    uint32
	Origin      int64
	Reason      int32
	Fee         float64
	Rote        uint32
	FinalPrice  float64
	XVolume     int32
	Action      uint32
	Notes       string
	CreatedAt   string
}

func (t *tdex) FuturesHistory(wq FuturesHistoryRequest) (*FuturesHistory, error) {
	return t.Service.FuturesHistory(wq)
}

type FuturesContract struct {
	Status uint32
	Data   GetContractData
	Errmsg string
}
type GetContractData struct {
	MarkPrice             string
	FairBasis             string
	FairBasisRate         string
	OpenInterest          string
	Turnover24h           string
	FundingRate           string
	FundingTimestamp      string
	IndicativeFundingRate string
}

func (t *tdex) FuturesContract(wq FuturesContractRequest) (*FuturesContract, error) {
	return t.Service.FuturesContract(wq)
}

type SpotBuy struct {
	Status uint32
	Data   SpotBuyData
	Errmsg string
}
type SpotBuyData struct {
	ID int
}

func (t *tdex) SpotBuy(sq SpotBuyRequest) (*SpotBuy, error) {
	return t.Service.SpotBuy(sq)
}

type SpotSell struct {
	Status uint32
	Data   SpotSellData
	Errmsg string
}
type SpotSellData struct {
	ID int
}

func (t *tdex) SpotSell(sq SpotSellRequest) (*SpotSell, error) {
	return t.Service.SpotSell(sq)
}

type SpotCancel struct {
	Status uint32
	Data   SpotCancelData
	Errmsg string
}
type SpotCancelData struct {
}

func (t *tdex) SpotCancel(sq SpotCancelRequest) (*SpotCancel, error) {
	return t.Service.SpotCancel(sq)
}

type SpotOrders struct {
	Status uint32
	Data   SpotOrdersData
	Errmsg string
}
type SpotOrdersData struct {
	List []SpotOrdersList
}
type SpotOrdersList struct {
	ID         int
	Symbol     string
	Currency   int
	Rid        int
	Direction  int
	Quantity   float64
	Completed  float64
	Price      float64
	Reality    float64
	Commission float64
	State      int
	CreateTime string
}

func (t *tdex) SpotOrders() (*SpotOrders, error) {
	return t.Service.SpotOrders()
}

type SpotHistory struct {
	Status uint64
	Data   SpotHistoryData
	Errmsg string
}
type SpotHistoryData struct {
	List      []SpotHistoryList
	Total     int32
	PageCount int32
	PageSize  int32
	Page      int32
}
type SpotHistoryList struct {
	ID         int
	Symbol     string
	Currency   int
	Rid        int
	Direction  int
	Quantity   float64
	Completed  float64
	Price      float64
	Reality    float64
	Commission float64
	State      int
	CreateTime string
}

func (t *tdex) SpotHistory(sq SpotHistoryRequest) (*SpotHistory, error) {
	return t.Service.SpotHistory(sq)
}

type WalletBalancesRequest struct {
	Type uint32
}
type WalletBalanceRequest struct {
	Currency uint32
	Type     uint32
}
type WalletSwitchRequest struct {
	Currency  uint32
	Direction uint32
	Amount    float64
}

type WalletWithdrawRequest struct {
	Currency uint32
	Address  string
	Amount   float64
}

type FuturesOpenRequest struct {
	Cid         int64
	Side        uint32
	Scale       float64
	Volume      uint32
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Passive     bool
	Visible     int32
	Strategy    uint32
	Better      bool
	Variable    uint32
	Constant    float64
	Sl          FuturesOpenRequestSl
	Tp          FuturesOpenRequestSTp
}
type FuturesOpenRequestSl struct {
	Distance bool
	Param    float64
}
type FuturesOpenRequestSTp struct {
	Distance bool
	Param    float64
}

type FuturesCloseRequest struct {
	Cid         int64
	ID          uint64
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Passive     bool
	Visible     int32
	Better      bool
}
type FuturesCloseAllRequest uint64

type FuturesCancelRequest struct {
	Cid int64
	ID  uint64
}
type FuturesSetslRequest struct {
	Cid         int64
	ID          uint64
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Passive     bool
	Visible     int32
	Better      bool
}
type FuturesSettpRequest struct {
	Cid         int64
	ID          uint64
	Distance    bool
	Price       float64
	Timely      uint32
	TimelyParam int32
	Strategy    uint32
	Variable    uint32
	Constant    float64
	Passive     bool
	Visible     int32
	Better      bool
}
type FuturesMergeRequest struct {
	Cid  int64
	List []uint64
}
type FuturesSplitRequest struct {
	Cid    int64
	ID     uint64
	Volume uint64
}
type FuturesSetupRequest struct {
	Cid    int64
	ID     uint64
	Target uint32
}
type FuturesSchemeSetRequest struct {
	Cid     uint32
	Options FuturesSchemeSetRequestOptions
}
type FuturesSchemeSetRequestOptions struct {
	Shared bool
	Merged bool
}
type FuturesSchemeGetRequest struct {
	Cid uint32
}
type FuturesHistoryRequest struct {
	PageSize int32
	Page     int32
}
type FuturesContractRequest struct {
	Symbol string
}
type SpotBuyRequest struct {
	Amount float64
	Price  float64
	Symbol string
}
type SpotSellRequest struct {
	Amount float64
	Price  float64
	Symbol string
}
type SpotCancelRequest struct {
	ID     int
	Symbol string
}
type SpotHistoryRequest struct {
	BeginTime string
	EndTime   string
	PageSize  int32
	Page      int32
}
