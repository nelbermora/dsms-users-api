# DPMS Users API

GET /users Shows a List of Users
GET /users/{userId} Shows User required if exists
POST /users Creates a new User
PUT /users Updates the user posted
DELETE /user/{userId} Deletes a determined user


Example of Response

```json

{
  "id": 1,
  "name": "John",
  "lastName": "Doe",
  "title": "Marketing Director",
  "branchId": 1,
  "role": 1,
  "status": "active",
  "branch": {
    "id": 1,
    "name": "Testing Branch",
    "address": "",
    "status": "Active"
  }
}
```
