```shell
$ jo user_id=123 order_items=$(jo -a $(jo product_code=prod quantity=4 unit_price=12))

{"user_id":123,"order_items":[{"product_code":"prod","quantity":4,"unit_price":12}]}

```

```shell
$ grpcurl -d '{"user_id":123,"order_items":[{"product_code":"prod","quantity":4,"unit_price":12}]}' -plaintext localhost:3000 Order/Create
{
  "orderId": "1"
}
$ grpcurl -d '{"user_id":123,"order_items":[{"product_code":"prod","quantity":4,"unit_price":12}]}' -plaintext localhost:3000 Order/Create
{
  "orderId": "2"
}
```
