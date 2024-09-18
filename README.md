# Golang Online shop RESTAPI

Example Golang API backend rest implementation mini project online using Gin Framework and Gorm ORM Database.

## Command

- ### App Lifecyle

```sh
$ go run main.go
```

- ### Docker Lifecycle

```sh
$ docker-compose up -d --build
```
## Endpoint

| **Nama**        | **Route**                  | **Method** |
| --------------- | -------------------------- | ---------- |
| **User**        |                            |            |
|                 | */api/v1/auth/register*    | *POST*     |
|                 | */api/v1/auth/login*       | *POST*     |
| **Product**     |                            |            |
|                 | */api/v1/product*          | *POST*     |
|                 | */api/v1/product*          | *GET*      |
|                 | */api/v1/product/:id*      | *GET*      |
|                 | */api/v1/product/:id*      | *DELETE*   |
|                 | */api/v1/product/:id*      | *UPDATE*   |
| **Cart**        |                            |            |
|                 | */api/v1/cart*             | *POST*     |
|                 | */api/v1/cart*             | *GET*      |
|                 | */api/v1/cart/:id*         | *GET*      |
|                 | */api/v1/cart/:id*         | *DELETE*   |
|                 | */api/v1/cart/:id*         | *PUT*      |
| **checkout**    |                            |            |
|                 | */api/v1/checkout*         | *POST*     |
|                 | */api/v1/checkout*         | *GET*      |
|                 | */api/v1/checkout/:id*     | *GET*      |
|                 | */api/v1/checkout/:id*     | *DELETE*   |
|                 | */api/v1/checkout/:id*     | *PUT*      |


## Feature
• Customer can view product list by product category (medium) (done) 
• Customer can add product to shopping cart (easy) (done)
• Customers can see a list of products that have been added to the shopping cart (medium)(done)
• Customer can delete product list in shopping cart (easy) (done)
• Customers can checkout and make payment transactions (hard)
• Login and register customers (DONE)


https://bgiri-gcloud.medium.com/designing-the-database-schema-for-a-new-e-commerce-platform-and-considering-factors-like-ec28d4fb81db



buat order terlebih dahulu kemudian set statusnya jadi unpaid, kemudian setelah user melakukan payment kemudian dilakukan pengecekan jika payment berhasil/sudah dilakukan maka akan mengubah status dari order menjadi paid