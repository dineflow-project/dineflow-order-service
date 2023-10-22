# dineflow-order-services

Open postman -> new grpc request
URL = 0.0.0.0:8080
Service Definition -> Using server reflection
Choose Service -> Enter Message for field -> Invoke

Example message for create order
{
"order_menus": [
{
"menu_id": "80000",
"price": 50,
"amount": 1,
"request": "Fast"
},
{
"menu_id": "80001",
"price": 30,
"amount": 2
}
],
"user_id": "6330654321",
"vendor_id": "50000"
}
