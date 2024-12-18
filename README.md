
Follow these steps to set up the project:
1.Open the "backend" folder and run the following command:
docker compose up -d --build

2. Use postman to send a POST request to: http://localhost:8080/api/quotes

Example of body request:

{
  "quote_code": "Q123456",
  "customer_id": 1,
  "status": "PENDING",
  "total_amount": 2500.00,
  "currency": "USD",
  "products": [
    {
      "product_id": 1,
      "quantity": 2,
      "price_per_unit": 1000.00,
      "tax_rate": 10.0,
      "total_price": 2200.00
    },
    {
      "product_id": 2,
      "quantity": 2,
      "price_per_unit": 50.00,
      "tax_rate": 0.0,
      "total_price": 100.00
    }
  ]
}
