# Shopping Cart

Hello, ini adalah API RESTfull sederhana buat data product.

Untuk menjalankan APInya pastikan local sudah terinstall golang dan docker.

## How to run

-   Jalankan docker terlebih dahulu dan jalankan docker-compose
    ```bash
    make server-run
    ```
-   Running app
    ```bash
    make local-run
    ```

## Stop docker

-   stop docker server
    ```bash
    make server-stop
    ```
-   remove docker server
    ```bash
    make remove
    ```

# Endpoint

## Create Product

```
localhost:8000/api/v1/products
```

### header

-   content-type : application/json

### request body

```json
{
    "product_name": "polpen hitam",
    "quantity": 20
}
```

### response

```json
{
    "code": 201,
    "message": "success create new products",
    "data": null
}
```

## List Product

```
localhost:8000/api/v1/products
```

### response

```json
{
    "code": 201,
    "message": "success get list product",
    "data": [
        {
            "product_code": "p-1",
            "product_name": "buku",
            "quantity": 10
        },
        {
            "product_code": "p-2",
            "product_name": "polpen",
            "quantity": 20
        }
    ]
}
```

## List Product Filter by Product name

```
localhost:8000/api/v1/products
```

### query params

-   product_name : "nama produk"

### response

```json
{
    "code": 201,
    "message": "success get list product",
    "data": [
        {
            "product_code": "p-1",
            "product_name": "buku",
            "quantity": 10
        }
    ]
}
```

## Update Product Kuantitas

```
localhost:8000/api/v1/products/{id}
```

### Params

-   id : product_code

### header

-   content-type : application/json

### request body

```json
{
    "quantity": 20
}
```

### response

```json
{
    "code": 201,
    "message": "success update products",
    "data": null
}
```

## Delete By product code

```
localhost:8000/api/v1/products/{id}
```

### Params

-   id : product_code

### response

```json
{
    "code": 201,
    "message": "success delete product",
    "data": ""
}
```
