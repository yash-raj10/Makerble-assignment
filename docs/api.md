# Clinic API Documentation (Swagger Style)

## Authentication

### POST /login

- **Description:** Login for both receptionist and doctor.
- **Request Body:**

```
{
  "username": "string",
  "password": "string"
}
```

- **Response:**

```
{
  "token": "JWT"
}
```

### POST /register

- **Description:** Register a new user (receptionist or doctor).
- **Request Body:**

```
{
  "username": "string",
  "password": "string",
  "role": "receptionist" | "doctor"
}
```

- **Response:**

```
{
  "message": "User registered"
}
```

---

## Patients (JWT required in `Authorization: Bearer <token>` header)

### POST /patients

- **Role:** receptionist
- **Description:** Register a new patient.
- **Request Body:**

```
{
  "firstName": "string",
  "lastName": "string",
  "age": int,
  "gender": "string",
  "address": "string",
  "phone": "string",
  "details": "string-{current patient detail - eg. inProgress/treated/done}"
}
```

- **Response:** Patient object (with ID, timestamps, etc.)

### GET /patients

- **Role:** any (doctor or receptionist)
- **Description:** List all patients.
- **Response:**

```
[
  { Patient }, ...
]
```

### GET /patients/{id}

- **Role:** any
- **Description:** Get patient by ID.
- **Response:** Patient object

### PUT /patients/{id}

- **Role:** doctor
- **Description:** Update patient details.
- **Request Body:** (same as POST /patients)
- **Response:** Updated patient object

### DELETE /patients/{id}

- **Role:** receptionist
- **Description:** Delete a patient by ID.
- **Response:**

```
{
  "message": "Patient deleted"
}
```

---

## Notes

- All protected endpoints require a valid JWT in the `Authorization` header.
- Timestamps (`createdAt`, `updatedAt`) are managed by the backend.(GORM)

---

## Example Patient Object

```
{
  "id": 1,
  "firstName": "yash",
  "lastName": "raj",
  "age": 20,
  "gender": "male",
  "address": "smi",
  "phone": "1234567890",
  "details": "Healthy",
  "updatedBy": 2,
  "createdAt": "2025-06-15T12:00:00Z",
  "updatedAt": "2025-06-15T12:00:00Z"
}
```
