# Golang Online Shop REST API

This project is an example of a Golang backend REST API implementation for an online shop. It uses the **Gin Framework** for routing and **Gorm ORM** for database interaction.

## Getting Started

### Setup Environment Variables

Before running the application, you need to set up environment variables to configure the connection to the PostgreSQL database.

1. Create a `.env` file in the root directory of the project.
2. Add the following lines to the `.env` file:

    ```env
    DB_PASSWORD=your_db_password   # Set this to the password for your PostgreSQL database
    DB_NAME=your_db_name           # Set this to the name of your PostgreSQL database
    ```

   - **DB_PASSWORD**: Set this to the password you want to use for your PostgreSQL database. Ensure it is strong and secure.
   - **DB_NAME**: Set this to the name of the PostgreSQL database that the API will connect to.

### Running the Application

Once you've set up the `.env` file with the correct values, you can start the application using Docker Compose.

### Docker Lifecycle Commands

To build and start the containers:

```sh
$ docker-compose up -d

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
