# REST API for ONLINE SHOP

## In this project:
- Developing Web Application following to design of REST API.
- Working with framework gin-gonic/gin.
- The Clean Architecture approach to building an application structure. Dependency implementation techniques.
- Working with MongoDB.
- Configuring application with help of library spf13/viper. Working with environment variables.
- Registration and authentication. JWT. Middleware.
- Writing NoSQL queries.
- Graceful Shutdown.
- Docker


### To launch application:

```
make run
```

### Endpoints:
- localhost:8001/auth/sign-up
```
JSON [example] (POST) - To sign-up
{
    "name": "Vladimir",
    "username": "qwerty",
    "password": "qwerty"
}
```
- localhost:8001/auth/sign-in
```
JSON [example] (POST) - To sign-in
  {
  "username": "qwerty",
  "password": "qwerty"
  }
```
### To use endpoints below, you need to set a bearer token, what was generated in response from sign-in endpoint.
- localhost:8001/api/product
```
JSON [example] (POST) - To create a product
{
    "name": "iPhone",
    "description": "11 Pro Max",
    "price": 899,
    "stock": 22
}
```
- localhost:8001/api/product/
````
To get products (GET)
````
- localhost:8001/api/product/`PRODUCTID`
````
To get product by id (GET)
````
- localhost:8001/api/product/`PRODUCTID`
````
To delete product by id (DELETE)
````
- localhost:8001/api/product/`PRODUCTID`
````
To update product by id (PUT)
{
    "name": "iPhone",
    "description": "12 Pro"
}
````
- localhost:8001/api/cart/
````
JSON [example] (POST) - To add a product in cart
{ 
    "product_id": "PRODUCT_ID",
    "quantity": 1
}
````
- localhost:8001/api/cart
````
JSON [example] (DELETE) - To delete a product from cart
{
    "product_id": "PRODUCT_ID"
}
````
- localhost:8001/api/cart
````
To get your cart (GET)
````
- localhost:8001/api/order
````
To place an order from your cart (GET)
````
- localhost:8001/api/order/`ORDER_ID`
````
To delete order (DELETE)
````
