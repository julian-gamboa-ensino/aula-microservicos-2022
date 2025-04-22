curl -X POST http://localhost:8081/api/products -H "Content-Type: application/json" -d '{
  "name": "Product 3",
  "category": "Category 3",
  "systemCode": "SC003",
  "customCode": "CC003",
  "costPrice": 30.0,
  "salePrice": 35.0,
  "measure": "pcs",
  "amount": 300,
  "timeToPrepare": "20 min",
  "description": "Description for product 3",
  "availability": true,
  "mobileApp": true,
  "active": true
}'

curl -X GET http://localhost:8081/api/products