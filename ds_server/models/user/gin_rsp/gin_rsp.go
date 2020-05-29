package gin_rsp

import mygormdl "ds_server/models/user/gorm_mysql"

type DsUserWalletExchangeRate_rsp struct {
    ErrorCode  int    `json:"error_code"`
    Resultcode string `json:"resultcode"`
    Reason     string `json:"reason"`
    Result     []struct {
        Data1 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data1"`
        Data2 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data2"`
        Data3 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data3"`
        Data4 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data4"`
        Data5 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data5"`
        Data6 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data6"`
        Data7 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data7"`
        Data8 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data8"`
        Data9 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data9"`
        Data10 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data10"`
        Data11 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data11"`
        Data12 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data12"`
        Data13 struct {
            BankConversionPri string      `json:"bankConversionPri"`
            Date              string      `json:"date"`
            FBuyPri           string      `json:"fBuyPri"`
            FSellPri          string      `json:"fSellPri"`
            MBuyPri           interface{} `json:"mBuyPri"`
            MSellPri          string      `json:"mSellPri"`
            Name              string      `json:"name"`
            Time              string      `json:"time"`
        } `json:"data13"`
        Data14 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data14"`
        Data15 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data15"`
        Data16 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data16"`
        Data17 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data17"`
        Data18 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data18"`
        Data19 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data19"`
        Data20 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data20"`
        Data21 struct {
            BankConversionPri string      `json:"bankConversionPri"`
            Date              string      `json:"date"`
            FBuyPri           interface{} `json:"fBuyPri"`
            FSellPri          string      `json:"fSellPri"`
            MBuyPri           string      `json:"mBuyPri"`
            MSellPri          string      `json:"mSellPri"`
            Name              string      `json:"name"`
            Time              string      `json:"time"`
        } `json:"data21"`
        Data22 struct {
            BankConversionPri string `json:"bankConversionPri"`
            Date              string `json:"date"`
            FBuyPri           string `json:"fBuyPri"`
            FSellPri          string `json:"fSellPri"`
            MBuyPri           string `json:"mBuyPri"`
            MSellPri          string `json:"mSellPri"`
            Name              string `json:"name"`
            Time              string `json:"time"`
        } `json:"data22"`
    } `json:"result"`
}
type Login_rsp struct {
    UUIDSelf     string  `json:"uuid_self"`     // 新注册的用户ID号
    MobileSelf   string  `json:"mobile_self"`   // 新注册的用户手机号
    NickName     string  `json:"nick_name"`     //
    InvcodeSelf  string  `json:"invcode_self"`  // 新注册的用户自身邀请码
    UUIDAgent    string  `json:"uuid_agent"`    // 代理的ID号
    MobileAgent  string  `json:"mobile_agent"`  // 代理的手机号
    InvcodeAgent string  `json:"invcode_agent"` // 代理的邀请码
    MemclassSelf int8    `json:"memclass_self"` // 会员等级
    MemberTag    string  `json:"member_tag"`    // 会员标识
    MemberName   string  `json:"member_name"`   // 会员名称
    AgentClass   int8    `json:"agent_class"`   // 代理等级
    AgentTag     string  `json:"agent_tag"`     // 代理标识
    AgentName    string  `json:"agent_name"`    // 代理名称
    Balance      float64 `json:"balance"`       // 余额
    Ispwd        int8    `json:"ispwd"`         // 密码是否为空
    Token        string  `json:"token"`
}
type MemAgent_rsp struct {
    UUIDSelf     string  `json:"uuid_self"`     // 新注册的用户ID号
    MobileSelf   string  `json:"mobile_self"`   // 新注册的用户手机号
    InvcodeSelf  string  `json:"invcode_self"`  // 新注册的用户自身邀请码
    UUIDAgent    string  `json:"uuid_agent"`    // 代理的ID号
    MobileAgent  string  `json:"mobile_agent"`  // 代理的手机号
    InvcodeAgent string  `json:"invcode_agent"` // 代理的邀请码
    MemclassSelf int8    `json:"memclass_self"` // 会员等级
    MemberTag    string  `json:"member_tag"`    // 会员标识
    MemberName   string  `json:"member_name"`   // 会员名称
    AgentClass   int8    `json:"agent_class"`   // 代理等级
    AgentTag     string  `json:"agent_tag"`     // 代理标识
    AgentName    string  `json:"agent_name"`    // 代理名称
    Balance      float64 `json:"balance"`       // 余额
    PayPwd       string  `json:"paypwd"`        // 支付密码
    Ispwd        int8    `json:"ispwd"`         // 密码是否为空
    Salt         string  `json:"salt"`
}

type MemberDepositLog_rsp struct {
    Memdephistory []*mygormdl.DsUserMemberDepositHistory `json:"deposit_history"`
    Num           int                                    `json:"num"` //总数
}

type MobileCode_rsp struct { //短信验证码
    ErrorCode int64  `json:"error_code"`
    Reason    string `json:"reason"`
    Result    struct {
        Sid string `json:"sid"`
    } `json:"result"`
}

type SuccessOrFail_rsp struct {
    ErrorCode  int    `json:"error_code"`
    Resultcode string `json:"resultcode"`
    Reason     string `json:"reason"`
}

type RechargeCallback_rsp struct {
    Sign            string `json:"sign"`
    ChainName       string `json:"chain_name"`
    CoinName        string `json:"coin_name"`
    Alias           string `json:"alias"`
    TradType        string `json:"trad_type"`
    BlockHeight     string `json:"block_height"`
    TransactionId   string `json:"transaction_id"`
    Trxn            string `json:"trx_n"`
    Confirmations   string `json:"confirmations"`
    FromAddress     string `json:"from_address"`
    ToAddress       string `json:"to_address"`
    Memo            string `json:"memo"`
    Amount          string `json:"amount"`
    Fee             string `json:"fee"`
    ContractAddress string `json:"contract_address"`
    OuterOrderNo    string `json:"outer_order_no"`
    ConfirmTime     string `json:"confirm_time"`
    Message         string `json:"message"`
}