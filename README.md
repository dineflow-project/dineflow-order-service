# dineflow-order-services

Add app.env

How to run:
air

Open postman -> new grpc request
URL = 0.0.0.0:8080
Service Definition -> Using server reflection
Choose Service -> Enter Message for field -> Invoke

example message for CreateOrder:
{
"MenuId": "70000",
"Status": "pending",
"UserId": "12345678",
"VenderId": "50000",
"Price" : 30
}
