export SERVER_NAME="Portofolio Rest API Service"
export SERVER_URL="127.0.0.1:3008"
export SERVER_READ_TIMEOUT=60

# JWT Environment
export JWT_SECRET_KEY="XL67xKuuTrBsPmECeEkMdN"
export JWT_EXPIRE_MINUTES=240

# Environment Database Connection
# export DB_SERVER_URL="host=localhost port=5432 user=postgres password=localhost dbname=kemendagrip3dn sslmode=disable"
export DB_SERVER_URL="host=localhost port=5432 user=postgres password=localhost dbname=p3dn sslmode=disable"

# Run app
go mod tidy
go run main.go