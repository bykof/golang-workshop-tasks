# Product Database

Use gorm with an sqlite connection for the database backend:
https://gorm.io/index.html

As you see, there are already some structs and interfaces defined.

There is the core module, where our product domain exists.
In there, we have our interface `ProductRepository`, which should be implemented 
by the `SqLiteProductRepository` in infrastructure/repositories/sqlite_product_repository.go.

Entities in the `infrastructure/entities` should have the gorm model.

Mappers should map the gorm models from `infrastructure/entities` to the domain models in core/domain/entities.

In the end, the `main.go` file should work.