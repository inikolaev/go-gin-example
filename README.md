# Go Service with Gin

Endpoints:

* `GET /ping`
* `POST /payments`
  ```bash
  curl --request POST \
    --url http://localhost:8080/payments \
    --header 'Content-Type: application/json' \
    --data '{
      "amount": {
        "value": 1000,
        "currency": "EUR"
      },
      "payment_method": {
        "type": "card",
        "id": "card-1"
      }
    }'
  ```
