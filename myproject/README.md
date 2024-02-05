# MyProject

MyProject is a simple HTTP server application written in Go that provides endpoints to manage user data.

## Features

- Retrieve user data by ID
- Update user data

## Usage

1. Run the compiled binary:

    ```bash
    ./myproject
    ```

2. Access the API endpoints using tools like `curl` or [Postman](https://www.postman.com/).

## API Endpoints

### GET /user

Retrieves user data by ID.

#### Parameters

- `id`: The ID of the user to retrieve.

#### Example

```bash
curl http://localhost:8080/user?id=1
```

### POST /user/update

Updates user data.

#### Request Body

```json
{
    "id": 1,
    "name": "Updated Name"
}
```

#### Example

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "name": "Updated Name"}' http://localhost:8080/user/update
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize this template according to your project's specific requirements and add more details as needed!