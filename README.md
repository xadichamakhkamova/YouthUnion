# 🌐 Youth Union Project

A **role-based event management platform** designed for student organizations and youth unions.  
It enables admins, organizers, and students to efficiently manage, participate in, and monitor events, rankings, and notifications within a unified system.

---

## 🚀 Overview

**Youth Union** is built with a **microservice architecture**, ensuring scalability, modularity, and clean separation of responsibilities between services.  
Each user role has access to specific functionalities and panels.

### 🎯 Main Roles
- **Admin**
- **Organizer**
- **Student**

---

## 🧩 Features by Role

### 🧑‍💼 Admin Panel
- Full CRUD for:
  - Users  
  - Roles  
  - Event Types  
- Manage system-wide media uploads  
- Send and receive automatic notifications  
- Configure global settings  

---

### 👨‍🏫 Organizer Panel
- Create, update, and delete events  
- View event statistics and student rankings  
- Send and receive notifications  
- Manage own profile settings  

---

### 👩‍🎓 Student Panel
- View and join available events  
- See personal and general rankings  
- Receive notifications from organizers or admin  
- Edit own profile (except unique ID)  

---

### 🌍 Public Dashboard
- Displays:
  - Active events  
  - Organizer/Team information  
- Provides login entry for authenticated users  

---

## 🧠 Core Concepts

### 🔑 Authentication & Authorization
- Secure login using **Identifier (ID)** and **Password**
- Role-based access control (RBAC)
- JWT-based token authentication handled through API Gateway

### 📊 Rankings
- Rankings are dynamically calculated based on participation and performance in events

### 📢 Notifications
- Real-time or queued notification system
- Admins and organizers can broadcast messages to users

### 🗃️ Database
- Each microservice maintains its own isolated database  
- Communication between services via **gRPC**

---
