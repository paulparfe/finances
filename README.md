###### Starting the application and database containers
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


###### Operations

- Checking the last 10 transactions of a user
http://localhost:8080/users/1/transactions

```json
{
  "data": [
    {
      "ID": 0,
      "UserID": 1,
      "RecipientID": 2,
      "Amount": "12.34",
      "TransactionType": "transfer",
      "CreatedAt": "2025-02-04T12:12:12Z"
    },
    {
      "ID": 0,
      "UserID": 1,
      "RecipientID": null,
      "Amount": "123.45",
      "TransactionType": "deposit",
      "CreatedAt": "2025-02-03T11:11:11Z"
    }
  ]
}
```


- Refill Adam's account for 12.34
```shell
curl -X POST http://localhost:8080/users/1/deposit -H "Content-Type: application/json" -d "{\"amount\": 12.34}"
```

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

