# Product Service

Take your solution from golang-workshop-tasks/05_types_interfaces_methods/02_product_service and copy it here.
Create this directory as a go module.
Now refactor the services into following structure:

- core
  - domain
    - entities
      - product.go
    - value_objects
      - rating.go
  - ports
    - product_service.go
- infrastructure
  - services
    - real_product_service.go
    - fake_product_service.go
- main.go