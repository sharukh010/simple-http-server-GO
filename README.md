# Go Web Server Example

This is a basic Go web server that serves static files and handles simple HTTP routes for demonstration purposes.

## Features

- Serves static files from the `./static` directory.
- Handles a GET request at `/hello`.
- Handles a POST request at `/form` and prints submitted form data.

## Endpoints

### 1. `/` (Static Files)
Serves files from the `static` directory. Place your `index.html`, CSS, JS, or images in `./static`.

### 2. `/hello` (GET)
Returns a greeting message.

**Example:**
```bash
curl http://localhost:8081/hello
````

### 3. `/form` (POST)

Parses and responds with form values `name` and `address`.

**Example using `curl`:**

```bash
curl -X POST -d "name=John&address=NYC" http://localhost:8081/form
```

## How to Run

1. Make sure you have [Go installed](https://golang.org/doc/install).
2. Place any static files (like `index.html`) inside a folder named `static`.
3. Run the server:

```bash
go run main.go
```

4. Open your browser and go to:
   [http://localhost:8081](http://localhost:8081)

## Example Output

**GET /hello**

```
hello! Sharukh Khan
```

**POST /form**

```
POST request successful
Name = John
Address = NYC
```

## Notes

* The server listens on port **8081**.
* If an invalid route or method is used, it returns a 404 or an appropriate error.

## License

This project is open source and available under the [MIT License](LICENSE).

