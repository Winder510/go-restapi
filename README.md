# Go MySQL REST API

Dự án REST API chuẩn doanh nghiệp sử dụng Go và MySQL.

## Cấu trúc dự án

```
project/
├── api/                    # API xử lý yêu cầu (handler)
│   ├── handler/           # API handlers
│   ├── middleware/        # API middleware
│   └── router.go          # Router
├── cmd/                   # Start app
│   ├── app/              # Main application
│   └── cli/              # Command line tools
├── config/               # Tệp cấu hình và mã liên quan
├── internal/             # Code riêng dùng nội bộ
│   ├── model/           # Model database
│   ├── repository/      # Truy cập cơ sở dữ liệu
│   ├── service/         # Logic business
│   └── util/            # Utilities
├── migrations/          # Database migrations
├── pkg/                # Thư viện code đã được khai báo sẵn
├── scripts/            # Build, install scripts
├── test/               # Tests
├── web/                # Frontend (nếu có)
├── .env.example        # Environment variables example
├── .gitignore          # Git ignore
├── LICENSE             # Giấy phép dự án
├── README.md           # Mô tả dự án
└── go.mod              # Go modules
```

## Yêu cầu hệ thống

- Go 1.21 hoặc cao hơn
- MySQL 8.0 hoặc cao hơn
- Git

## Cài đặt

### 1. Clone repository

```bash
git clone <repository-url>
cd go_mysql
```

### 2. Cài đặt dependencies

```bash
go mod download
```

### 3. Cấu hình môi trường

Tạo file `.env` từ `.env.example`:

```bash
cp .env.example .env
```

Cập nhật các giá trị trong file `.env`:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080
SERVER_MODE=debug

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=go_mysql_db

# JWT Configuration
JWT_SECRET=your-secret-key-change-this-in-production
JWT_EXPIRY=24h
```

### 4. Tạo database

```bash
mysql -u root -p
CREATE DATABASE go_mysql_db;
```

### 5. Chạy migrations

**Linux/Mac:**

```bash
chmod +x scripts/migrate.sh
./scripts/migrate.sh

mysql -h127.0.0.1 -P3306 -uroot -p 123456
```

**Windows:**
Chạy SQL files trong thư mục `migrations/` thủ công hoặc sử dụng MySQL Workbench.

### 6. Chạy ứng dụng

```bash
go run cmd/app/main.go
```

Server sẽ chạy tại: `http://localhost:8080`

## API Endpoints

### Authentication

#### Register

```bash
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123",
  "full_name": "John Doe"
}
```

#### Login

```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

### Users (Yêu cầu authentication)

#### Get all users

```bash
GET /api/v1/users?page=1&page_size=10
Authorization: Bearer <token>
```

#### Get user by ID

```bash
GET /api/v1/users/:id
Authorization: Bearer <token>
```

#### Update user

```bash
PUT /api/v1/users/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "newusername",
  "email": "newemail@example.com",
  "full_name": "New Name",
  "status": "active"
}
```

#### Delete user

```bash
DELETE /api/v1/users/:id
Authorization: Bearer <token>
```

### Health Check

```bash
GET /health
```

## Build

### Linux/Mac

```bash
chmod +x scripts/build.sh
./scripts/build.sh
```

### Windows

```bash
scripts\build.bat
```

Binary sẽ được tạo trong thư mục `bin/`.

## Testing

```bash
go test ./...
```

## Cấu trúc code

### Clean Architecture

Dự án sử dụng Clean Architecture với các layers:

1. **Handler Layer** (`api/handler`): Xử lý HTTP requests/responses
2. **Service Layer** (`internal/service`): Business logic
3. **Repository Layer** (`internal/repository`): Database operations
4. **Model Layer** (`internal/model`): Data structures

### Dependency Injection

Dependencies được inject từ `main.go`:

```
main -> handler -> service -> repository -> database
```

## Công nghệ sử dụng

- **Framework**: Gin Web Framework
- **Database**: MySQL
- **ORM**: database/sql (native)
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Configuration**: godotenv

## Best Practices

1. **Error Handling**: Tất cả errors được handle và trả về response phù hợp
2. **Validation**: Input validation sử dụng Gin binding tags
3. **Security**:
   - Password được hash bằng bcrypt
   - JWT authentication cho protected routes
   - CORS middleware
4. **Code Organization**: Clean Architecture, separation of concerns
5. **Database**: Connection pooling, prepared statements
6. **Logging**: Request logging middleware

## License

MIT License

## Tác giả

Your Name

## Liên hệ

- Email: your.email@example.com
- GitHub: @yourusername
