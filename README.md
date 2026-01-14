# k8s-cost-analyzer
##  MVP

# Structure 


k8s-cost-analyzer/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   └── auth_handler.go
│   ├── middleware/
│   │   └── auth_middleware.go
│   ├── models/
│   │   └── user.go
│   ├── repositories/
│   │   └── user_repo.go
│   └── services/
│       └── auth_service.go
│
└── pkg/
    └── k8s/
