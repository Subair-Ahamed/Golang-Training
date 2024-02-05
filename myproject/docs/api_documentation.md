# API Documentation

## Endpoints

### GET /user
- Description: Get user information by ID.
- Parameters:
  - id (integer): The ID of the user to retrieve.
- Response:
  - 200 OK: User data retrieved successfully.
  - 404 Not Found: User with the specified ID not found.

### POST /user/update
- Description: Update user information.
- Parameters:
  - id (integer): The ID of the user to update (in the query string).
  - Request Body: User object with updated information.
- Response:
  - 200 OK: User updated successfully.
  - 400 Bad Request: Invalid request body format.
  - 404 Not Found: User with the specified ID not found.
