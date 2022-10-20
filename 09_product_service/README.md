# Product Service

Now let's build a simple product service, where we should be able to interact via REST calls.
Use the gin framework: https://gin-gonic.com/docs/

Copy over your infrastructure result from the task `08_database`.

Now implement the controller in `interfaces/product_controller.go`, where each function of the 
product_repository should be exposed as a REST endpoint:

- GET /products -> List()
- POST /products -> Add()
- DELETE /products/1 -> Remove()
- PATCH /products/1 -> Save()

Try to create dtos in `interfaces/controllers/dtos` to create the structs for the POST and PUT mappings.

In the end, you should be able to map each function from the `ProductController` to a gin route:

```go
r := gin.Default()
r.GET("/products", productController.List)
```