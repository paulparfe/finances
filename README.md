## Starting the application and database containers
```shell
make run
```


This will create 2 users and a small transaction history.

Table `users`

| id  | name  | balance |
|-----|-------|---------|
| 1   | Adam  | 111.11  |
| 2   | Diana | 12.34   |

Table `transactions`

| id  | user_id | recipient_id | amount | transaction_type | created_at          |
|-----|---------|--------------|--------|------------------|---------------------|
| 1   | 1       | NULL         | 123.45 | deposit          | 2025-02-03 11:11:11 |
| 2   | 1       | 2            | 12.34  | transfer         | 2025-02-04 12:12:12 |


## Operations

#### Checking the last 10 transactions of a user
Request: 
http://localhost:8080/users/1/transactions

Returns:
```json
{
  "data": [
    {
      "ID": 2,
      "UserID": 1,
      "RecipientID": 2,
      "Amount": "12.34",
      "TransactionType": "transfer",
      "CreatedAt": "2025-02-04T12:12:12Z"
    },
    {
      "ID": 1,
      "UserID": 1,
      "RecipientID": null,
      "Amount": "123.45",
      "TransactionType": "deposit",
      "CreatedAt": "2025-02-03T11:11:11Z"
    }
  ]
}
```


#### Refill Adam's account for 12.34
Request:
```shell
curl -X POST http://localhost:8080/users/1/deposit -H "Content-Type: application/json" -d "{\"amount\": 12.34}"
```
Returns: `{"data":{"id":1,"name":"Adam","balance":"123.45"}}`

The new values will be as follows:

| id  | name  | balance |
|-----|-------|---------|
| 1   | Adam  | 123.34  |
| 2   | Diana | 12.34   |

| id | user_id | recipient_id | amount | transaction_type | created_at          |
|----|---------|--------------|--------|------------------|---------------------|
| 1  | 1       | NULL         | 123.45 | deposit          | 2025-02-03 11:11:11 |
| 2  | 1       | 2            | 12.34  | transfer         | 2025-02-04 12:12:12 |
| 3  | 1       | NULL         | 12.34  | deposit          | 2025-02-04 13:13:13 |


#### Diana is transferring 12.33 to Adam 
Request:
```shell
curl -X POST http://localhost:8080/users/2/transfer -H "Content-Type: application/json" -d "{\"recipient_user_id\": 1,\"amount\": 12.33}"
```
Returns: `{"data":{"id":2,"name":"Diana","balance":"0.01"}}`

The new values will be as follows:

| id  | name  | balance |
|-----|-------|---------|
| 1   | Adam  | 135.78  |
| 2   | Diana | 0.01    |

| id | user_id | recipient_id | amount | transaction_type | created_at          |
|----|---------|--------------|--------|------------------|---------------------|
| 1  | 1       | NULL         | 123.45 | deposit          | 2025-02-03 11:11:11 |
| 2  | 1       | 2            | 12.34  | transfer         | 2025-02-04 12:12:12 |
| 3  | 1       | NULL         | 12.34  | deposit          | 2025-02-04 13:13:13 |
| 4  | 2       | 1            | 12.33  | transfer         | 2025-02-04 14:14:14 |

## Run tests

```shell
go test -v ./internal/domain/service
```
