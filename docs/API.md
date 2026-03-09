# API Documentation

## Postman Collection Examples

### 1. Register User

**Endpoint:** `POST /api/v1/auth/register`

**Request:**

```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123",
  "full_name": "John Doe"
}
```

**Success Response (201):**

```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "uuid-here",
    "username": "johndoe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "status": "active",
    "created_at": "2026-03-09T10:00:00Z",
    "updated_at": "2026-03-09T10:00:00Z"
  }
}
```

### 2. Login

**Endpoint:** `POST /api/v1/auth/login`

**Request:**

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (200):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "uuid-here",
      "username": "johndoe",
      "email": "john@example.com",
      "full_name": "John Doe",
      "status": "active",
      "created_at": "2026-03-09T10:00:00Z",
      "updated_at": "2026-03-09T10:00:00Z"
    }
  }
}
```

### 3. Get All Users (Protected)

**Endpoint:** `GET /api/v1/users?page=1&page_size=10`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Success Response (200):**

```json
{
  "success": true,
  "data": [
    {
      "id": "uuid-here",
      "username": "johndoe",
      "email": "john@example.com",
      "full_name": "John Doe",
      "status": "active",
      "created_at": "2026-03-09T10:00:00Z",
      "updated_at": "2026-03-09T10:00:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 10,
    "total_items": 1,
    "total_pages": 1
  }
}
```

### 4. Get User by ID (Protected)

**Endpoint:** `GET /api/v1/users/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Success Response (200):**

```json
{
  "success": true,
  "data": {
    "id": "uuid-here",
    "username": "johndoe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "status": "active",
    "created_at": "2026-03-09T10:00:00Z",
    "updated_at": "2026-03-09T10:00:00Z"
  }
}
```

### 5. Update User (Protected)

**Endpoint:** `PUT /api/v1/users/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Request:**

```json
{
  "username": "newusername",
  "full_name": "New Full Name",
  "status": "inactive"
}
```

**Success Response (200):**

```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": "uuid-here",
    "username": "newusername",
    "email": "john@example.com",
    "full_name": "New Full Name",
    "status": "inactive",
    "created_at": "2026-03-09T10:00:00Z",
    "updated_at": "2026-03-09T10:10:00Z"
  }
}
```

### 6. Delete User (Protected)

**Endpoint:** `DELETE /api/v1/users/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Success Response (200):**

```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

## Error Responses

### 400 Bad Request

```json
{
  "success": false,
  "error": "email already exists"
}
```

### 401 Unauthorized

```json
{
  "success": false,
  "error": "Authorization header required"
}
```

### 404 Not Found

```json
{
  "success": false,
  "error": "user not found"
}
```

### 500 Internal Server Error

```json
{
  "success": false,
  "error": "internal server error"
}
```
