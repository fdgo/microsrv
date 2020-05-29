package dto

import "github.com/shopspring/decimal"

type GetBalanceRequest struct {
	ClientID string `url:"client_id"`
	CoinName string `url:"coin_name"`
	Sign     string `url:"sign,omitempty"`
}

type GetBalanceResponse struct {
	Response
	Data struct {
		Token_name string `json:"token_name"`
		Balance    string `json:"balance"`
		Frozen     string `json:"frozen"`
	} `json:"data"`
}

type Response struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GetAddressRequest struct {
	ClientID string `url:"client_id"`
	CoinName string `url:"coin_name"`
	Num      int    `url:"num"`
	Sign     string `url:"sign,omitempty"`
}

type GetAddressResponse struct {
	Response
	Data []string `json:"data"`
}

type GetWithdrawRequest struct {
	ClientID        string          `url:"client_id"`
	CoinName        string          `url:"coin_name"`
	TokenName       string          `url:"token_name"`       //是	代币名称（实际提现的币种名，如：USDT、USDT-ERC20等），提现时必须填写
	OrderID         string          `url:"order_id"`         //是	订单号，唯一，最长48位长度
	Amount          decimal.Decimal `url:"amount"`           //是	提现金额
	Fee             decimal.Decimal `url:"fee"`              //否	费用
	ToAddress       string          `url:"to_address"`       //是	提现地址
	Memo            string          `url:"memo"`             //否	Memo，Tag，Payment ID等
	ContractAddress string          `url:"contract_address"` //否	代币合约地址，只有转出代币时，才需要填写
	SendNum         int             `url:"send_num"`         //是	提现序号，提现币种最后一次提现记录的值+1，即为该字段的值（如果没有提现值为1），Hoo会对改值进行验证
	Sign            string          `url:"sign,omitempty"`
}

type GetWithdrawResponse struct {
	Response
	Data struct {
		OrderNo string `json:"order_no"`
	} `json:"data"`
}

type GetPushOrderRequest struct {
	ClientID  string `url:"client_id"`
	CoinName  string `url:"coin_name"`
	Txid      string `url:"txid"`
	Sign      string `url:"sign,omitempty"`
	ToAddress string `url:"to_address"` //发送目标地址，可选
}

type GetPushOrderResponse struct {
	Response
	Data []GetPushOrderResponse2 `json:"data"`
}

type GetPushOrderResponse2 struct {
	OuterOrderNo  string `json:"outer_order_no"` // 订单号
	TradeType     string `json:"trade_type"`     // 交易类型：1发送，2接收
	CoinName      string `json:"coin_name"`      // 提现币种
	ChainName     string `json:"chain_name"`     // 主链名称
	TransactionID string `json:"transaction_id"` // TXID
	BlockHeight   string `json:"block_height"`   // 区块高度
	Confirmations string `json:"confirmations"`  // 确认数（Hoo）
	FromAddress   string `json:"from_address"`   // 发送地址
	ToAddress     string `json:"to_address"`     // 接收地址
	Amount        string `json:"amount"`         // 发送金额
	Fee           string `json:"fee"`            // 费用
	Status        string `json:"status"`         // 状态，success(完成)，confirming(确认中)，rollback(订单回滚)
	CreateAt      string `json:"create_at"`      // 创建时间
	ProcessAt     string `json:"process_at"`     // 完成时间
}

type GetOrdersRequest struct {
	ClientID  string `url:"client_id"`
	CoinName  string `url:"coin_name"`
	TradeType int    `url:"trade_type"` // 否	交易类型：1发送，2接收
	StartAt   int    `url:"start_at"`   //	否	开始时间
	EndAt     int    `url:"end_at"`     //否	截止时间
	PageNum   int    `url:"pagenum"`    //页码，默认1
	PageSize  int    `url:"pagesize"`   //条数，默认100，最大100
	Sign      string `url:"sign,omitempty"`
}

type GetOrdersResponse struct {
	Response
	Data struct {
		Count   int                  `json:"count"`
		Pagenum int                  `json:"pagenum"`
		Records []GetOrdersResponse2 `json:"records"`
	} `json:"data"`
}

type GetOrdersResponse2 struct {
	OuterOrderNo  string `json:"outer_order_no"` // 订单号
	TradeType     string `json:"trade_type"`     // 交易类型：1发送，2接收
	CoinName      string `json:"coin_name"`      // 提现币种
	ChainName     string `json:"chain_name"`     // 主链名称
	TransactionID string `json:"transaction_id"` // TXID
	BlockHeight   string `json:"block_height"`   // 区块高度
	Confirmations string `json:"confirmations"`  // 确认数（Hoo）
	FromAddress   string `json:"from_address"`   // 发送地址
	ToAddress     string `json:"to_address"`     // 接收地址
	Amount        string `json:"amount"`         // 发送金额
	Fee           string `json:"fee"`            // 费用
	Status        string `json:"status"`         // 状态，success(完成)，confirming(确认中)，rollback(订单回滚)
	CreateAt      string `json:"create_at"`      // 创建时间
	ProcessAt     string `json:"process_at"`     // 完成时间
}
