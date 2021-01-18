package response

// YouzanAppStoreOpenTradeCreateResponse 用于创建应用内购订单
type YouzanAppStoreOpenTradeCreateResponse struct {
	baseResponse
	Data struct {
		// payment_url
		// java.lang.String
		// 有赞应用市场付款地址
		// https://app.youzanyun.com/transaction/payment?orderNo=E20191219195731028700001
		PaymentURL string `json:"payment_url"`
		// goods_name
		// java.lang.String
		// 商品名称
		// 测试内购
		GoodsName string `json:"goods_name"`
		// order_no
		// java.lang.String
		// 老订单Id
		// E20191219195731028700001
		OrderNO string `json:"order_no"`
		// order_status_enum
		// java.lang.String
		// 枚举类型：订单状态
		// WAIT_PAY
		OrderStatusEnum string `json:"order_status_enum"`
		// cashier_url
		// java.lang.String
		// 有赞H5收银台地址
		// https://cashier.youzan.com/pay/preorder?prepay_id=PT1653349182078976&cashier_sign=B9ECC1CDD0C6D506DB85D49B12FD3415&cashier_salt=1576756651843&partner_id=820000000003
		CashierURL string `json:"cashier_url"`
		// pay_amount
		// java.lang.Long
		// 支付金额，单位（分）
		// 1000
		PayAmount int64 `json:"pay_amount"`
	} `json:"data"`
}
