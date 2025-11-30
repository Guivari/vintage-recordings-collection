# vintage-recordings-collection


Installation step-by-step:
1. go run build
2. go list -f '{{.Target}}'
(if not yet done, adding Go install directory to system's shell path)
LINUX or MAC:
$ export PATH=$PATH:/path/to/your/install/directory
WINDOWS
$ set PATH=%PATH%;C:\path\to\your\install\directory
3. go install



Usage:
$ vintage-recordings-collection {port}
(server uses port localhost:8080 by default)



Endpoints:
- GET   /albums
- GET   /albums/{id}
- POST  /albums


Examples:
$ curl http://localhost:8080/albums
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    ...
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]

$ curl http://localhost:8080/albums/1
{
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
}

$ curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 30 Nov 2025 04:27:15 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}