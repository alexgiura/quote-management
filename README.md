# **Quote Management API - Setup Guide**

This guide explains how to set up and run the **Quote Management API** locally using **Docker Compose**.

---

## **Project Setup Steps**

### **1. Run the API Using Docker Compose**

1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```

2. Build and run the application using **Docker Compose**:
   ```bash
   docker compose up -d --build
   ```

---

### **2. Send a POST Request to Create a Quote**

Use **Postman** (or another API client) to send a **POST request** to:

```
http://localhost:8080/api/quotes
```

---

### **Example Request Body (JSON)**

```json
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
```

---

### **3. Verify Response (Example)**

Expect a **201 Created** response with the created quote details.

---

### **Useful Commands**

- **Stop Containers:**
  ```bash
  docker compose down
  ```

- **View Docker Logs:**
  ```bash
  docker logs <container_name>
  ```

---

### **Project Overview**

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Containerization:** Docker, Docker Compose
- **API Testing:** Postman, HTTP Clients

**Project Structure:**

```
quote-management/
  ├── backend/            # Source code, Dockerfile, and Go modules
  ├── cmd/                # Main application entry point
  └── internal/           # Core business logic
```

---

Let me know if you need further improvements! 🚀

