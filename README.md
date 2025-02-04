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


###### Checking the last 10 transactions of a user
http://localhost:8080/users/1/transactions

