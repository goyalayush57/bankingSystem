# Simple Bank APIs

# API Contract 
https://goyalayush57.github.io/swaggerTemplate/


This repo exposes some basic set of APIs that can be used to perform CRUD operation for employee and customer.
# Features
- Admin role with features
    - Sign in/out as admin
    - Add bank employees
    - Delete employees
- Employee role with feature
    - Sign in/out as an employee
    - Create a customer
    - Create accounts like savings, salary, loan, current account
    - Link customers with accounts
    - Update KYC for a customer
    - Get details of a customer
    - Delete customer
    - Get account balance for an account
    - Transfer money from one account to another
    - Print Account statement of an account for a time range in pdf
    - Calculate interest for the money annually (at 3.5% p.a.) and update the account balance

# Local Setup
### Prerequisite
1. PostGres should be installed
2. Create a user in post Gres
3. Create a DB from that user
4. Mention these configuration in the .env file
5.Postman should be installed if you want to play with the APIS using postman(URLs are under api/postman folder)

### Starting the application
- Just **go build** with the **.env** file changes and now application can be played with

### Default Admin Account
    Username : sa-100001
    Password : 1234

# Table used
[draw.io](Draw.io)

# Tech/Important Library Used
- Go
- PostGres
- JWT for authentication
- GORM for ORM

# Code Design Walkthrough
Every Function is wrapped around Interface to have a clean and loosely coupled code.
Code is organised in three layers
- Controller
    - Authenticates user
    - Parses the Request
    - Passes it to Service layer
- Service
    - Applies the logic if any
    - Passes it to DAL layer
- Data Access Layer
    - Interacts with Databse

Code will start from main.go -> Controller -> Service -> DB
