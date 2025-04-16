# Simple Album Management API in Go

A basic API to manage a collection of music albums, built with Gin (a Go web framework).  
Includes basic CRUD operations and is ideal for learning REST API fundamentals.

---

## 📌 Features

- **Supported HTTP methods**: GET, POST, PUT, DELETE
- **Endpoints**:
  - Get all albums
  - Get an album by ID
  - Create a new album
  - Update an existing album (full replacement)
  - Delete an album

---

## ⚙️ Requirements

- Go 1.20+ installed
- Dependency: [Gin Web Framework](https://github.com/gin-gonic/gin)

---

## 🛠️ Installation

1. Clone the repository:
```

git clone https://github.com/your-username/your-repo.git

```
2. Install Gin:
```

go get -u github.com/gin-gonic/gin

```
3. Run the API:
```

go run main.go

```

---

## 📡 Endpoints and Examples

### 1. Get all albums
```

GET /albums

```
**Example**:
```

curl http://localhost:8080/albums

```

### 2. Get an album by ID
```

GET /albums/:id

```
**Example**:
```

curl http://localhost:8080/albums/1

```

### 3. Create a new album
```

POST /albums

```
**Example**:
```

curl -X POST -H "Content-Type: application/json" -d '{
"id": "4",
"title": "Thriller",
"artist": "Michael Jackson",
"price": 19.99
}' http://localhost:8080/albums

```

### 4. Update an album (PUT)
```

PUT /albums/:id

```
**Example**:
```

curl -X PUT -H "Content-Type: application/json" -d '{
"id": "1",
"title": "Blue Train (Special Edition)",
"artist": "John Doe",
"price": 25.99
}' http://localhost:8080/albums/1

```

### 5. Delete an album
```

DELETE /albums/:id

```
**Example**:
```

curl -X DELETE http://localhost:8080/albums/2

```

---

## 🧪 Testing

1. Use **curl** or tools like [Postman](https://www.postman.com/).
2. Check HTTP responses:
   - `200 OK`: Successful operation
   - `201 Created`: Resource created
   - `400 Bad Request`: Invalid data
   - `404 Not Found`: Resource not found

---

## 📄 License
MIT License. Educational project.  
**Note**: This is a basic implementation without real persistence (data is stored in memory).

