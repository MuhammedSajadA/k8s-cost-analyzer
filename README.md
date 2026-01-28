
# Kubernetes Cost Analyzer 

A Go-based backend service for securely onboarding Kubernetes clusters using kubeconfig files.  
This project serves as the foundation of a multi-tenant SaaS for analyzing Kubernetes resource usage by.
- Connect to a Kubernetes cluster
- Analyze resource usage
- Detect waste
- Show optimization suggestions

---
## 🚀 MVP Features

- JWT-based user authentication
- Secure cluster connection via kubeconfig upload
- Namespace discovery
- Per-namespace CPU and memory usage visibility
- Resource waste indicator
  - Requested vs actual usage
- Lightweight dashboard UI

## 🏗️ Architecture

**Architecture Flow**

- React Frontend
  - ↓
- Nginx Ingress
  - ↓
- Go API Service → PostgreSQL
  - ↓
- Kubernetes Client
  - ↓
- Target K8s Cluster

## ✨ Features (Current)

- JWT-based user authentication
- Secure Kubernetes cluster onboarding via kubeconfig
- Cluster connectivity validation using `client-go`
- Multi-tenant cluster ownership enforcement
- Clean layered architecture (handler / service / repository)

> ℹ️ This repository currently covers functionality **up to cluster onboarding**.  
> Namespace listing, metrics collection, and cost analysis are planned next.

---



## 🔐 Security Notes

- All cluster operations are JWT-protected
- Cluster ownership is enforced (multi-tenant safe)
- kubeconfig is validated before persistence
- JWT claim parsing is handled safely (no runtime panics)

> ⚠️ **Note:** kubeconfig encryption at rest will be added in a future PR.

## 🧭 Next Steps (Planned)

- Namespace listing per cluster
- CPU & memory metrics via metrics.k8s.io
- Waste calculation & cost estimation
- Frontend dashboard
- kubeconfig encryption & RBAC hardening
- CI/CD & Helm deployment

## 🧑‍💻 Author

Built as a hands-on platform engineering project to demonstrate:
- Go backend design
- Kubernetes API integration
- Secure SaaS-style multi-tenancy
- Production-grade debugging & error handling
