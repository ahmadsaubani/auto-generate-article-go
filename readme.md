#### STRUCTURE PROJECT
```sh
project-root/
├── src/                             # Source Code
│   ├── configs/
│   │   └── database/
│   │       └── database.go
│   │
│   ├── controllers/
│   │   └── api/v1/
│   │       ├── customers/
│   │       ├── product_flavours/
│   │       ├── product_sizes/
│   │       ├── product_types/
│   │       ├── products/
│   │       └── transactions/
│   │
│   ├── dtos/
│   │   ├── customer_dtos/
│   │   ├── product_dtos/
│   │   ├── product_flavour_dtos/
│   │   ├── product_size_dtos/
│   │   ├── product_type_dtos/
│   │   ├── redeem_dtos/
│   │   ├── transaction_dtos/
│   │   └── transaction_item_dtos/
│   │
│   ├── entities/
│   │   ├── customers/
│   │   ├── product_flavours/
│   │   ├── product_sizes/
│   │   ├── product_types/
│   │   ├── products/
│   │   ├── transaction_items/
│   │   └── transactions/
│   │
│   ├── helpers/
│   │   ├── helper.go
│   │   └── response.go
│   │
│   ├── repositories/
│   │   ├── customer_repositories/
│   │   ├── product_flavour_repositories/
│   │   ├── product_repositories/
│   │   ├── product_size_repositories/
│   │   ├── product_type_repositories/
│   │   └── transaction_repositories/
│   │
│   ├── routes/
│   │   └── route.go
│   │
│   ├── seeders/
│   │   ├── product_flavour_seeders/
│   │   ├── product_size_seeders/
│   │   ├── product_type_seeders/
│   │   └── seeder.go
│   │
│   ├── services/
│   │   ├── customer_services/
│   │   ├── product_flavour_services/
│   │   ├── product_services/
│   │   ├── product_size_services/
│   │   ├── product_type_services/
│   │   └── transaction_services/
│   │
│   ├── traits/
│   │   ├── generate_uuid.go
│   │   └── utils/filters/
│   │       └── filter.go
│
├── main.go                          # Main entry point
├── .env                             # Environment variables
├── go.mod                           # Go modules definition
├── go.sum                           # Go dependencies lock file
└── README.md


```

## Getting Started
1. **Clone the repository**:
```bash
git clone https://github.com/ahmadsaubani/bsnake-go.git
```

2. **How To Run Development**:
```sh
Requirements:
- go > 1.24.4
- postgre


# Salin file .env default
cp .env.example .env
# akan create migration dan seeder untuk master data
go run main.go 

```
3. **Filter Usage**:
```sh
/api/v1/endpoint?field[like]=john&field[moreThan]=18&order_by=id,desc

Example :
1. /api/v1/product/types?order_by=id,desc
```