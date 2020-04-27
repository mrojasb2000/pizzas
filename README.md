# PIZZA Restaurant
Pizza restaurant API

Here's a small implementation of a pizza restaurant API with three endpoints:

* List all pizzas on the menu: GET /pizzas
* Make a simple pizza order: POST /orders
* List all orders in the system: GET /orders

curl -X POST localhost:8080/orders -d '{"pizza_id":1,"quantity":0}' 

curl -X GET localhost:8080/orders
