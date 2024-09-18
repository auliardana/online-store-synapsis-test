# Golang Online Shop REST API

This project is an example of a Golang backend REST API implementation for an online shop. It uses the **Gin Framework** for routing and **Gorm ORM** for database interaction.

### Docker Lifecycle Commands

To build and start the containers:

```sh
$ docker-compose up -d
```

To access API documentation:

```sh
http://localhost:8080/swagger/index.html#/
```

## Endpoint

| **Nama**        | **Route**                  | **Method** |
| --------------- | -------------------------- | ---------- |
| **auth**        |                            |            |
|                 | */api/v1/register*         | *POST*     |
|                 | */api/v1/login*            | *POST*     |
| **Product**     |                            |            |
|                 | */api/v1/auth/product*     | *POST*     |
|                 | */api/v1/auth/product*     | *GET*      |
| **Cart**        |                            |            |
|                 | */api/v1/auth/cart*        | *POST*     |
|                 | */api/v1/auth/cart*        | *GET*      |
|                 | */api/v1/auth/cart/:id*    | *DELETE*   |
| **Category**    |                            |            |
|                 | */api/v1/auth/category*    | *POST*     |
|                 | */api/v1/auth/category*    | *GET*      |
| **order**       |                            |            |
|                 | */api/v1/order*            | *POST*     |
|                 | */api/v1/order*            | *GET*      |



## Feature
• Customer can view product list by product category
• Customer can add product to shopping cart
• Customers can see a list of products that have been added to the shopping cart
• Customer can delete product list in shopping cart
• Customers can checkout and make payment transactions
• Login and register customers


https://bgiri-gcloud.medium.com/designing-the-database-schema-for-a-new-e-commerce-platform-and-considering-factors-like-ec28d4fb81db
