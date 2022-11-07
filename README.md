# oxpay-go

The go-sdk used to invoke oxpay payments in Singapore

## 开发背景

之前是接了stripe支付，发现stripe的go-sdk很好用。后面接触了OxPay，发现没有go相关的sdk，所以决定写个go-sdk

## 开发版本

- 目前是针对oxpay-api的v5版本的ewallet(电子钱包进行开发)

## 开发目标

- 相关结构体的构造(下单，退款等都将作为一个实体类型来构造)

- 优雅的调用oxpay-api(封装api)

## 功能列表

- 支付功能
  - 银行卡快捷支付
  - 电子钱包支付(支付宝、微信、paynow、grab等)
- 退款功能
  - 由于OxPay建议的退款方式为商户和用户协商，所以线上退款功能仅支持部分的支付方式
    - 线上退款方式：微信
- 账单查询
- 取消订单
  - OxPay目前仅支持部分支付方式的取消订单

