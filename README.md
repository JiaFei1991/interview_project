# Interview web server

### Project Description

A Go web server that listens for POST requests on [endpoint](localhost:3000/postEndpoint) and stores the content included in the request body in a JSON file.
This endpoint expects incoming requests to conform to the following JSON format:

```json
{
  "id": 1,
  "firstName": "Name",
  "lastName": "Name",
  "emailAddress": "test@mail.com",
  "validUser": true
}
```

Since the program only defines the POST method on the endpoint, attemps using other http methods will receive a bad request error message **SERVER HANDLES POST REQUESTS ONLY**.
The program will respond with a bad request error with message **REQUIRED FIELDS UNDEFINED** if any field in the above JSON object is undefined in the request. Fields that do not belong to the JSON object will be ignored.
The data is stored in a file named _data.json_ in the same project directory. The program will create the file if it does not exist.
A JSON object is send back to the user if the request were successful.

### Installation and usage

1. Download a Go binary release specific to your OS on [go_instal](https://go.dev/dl/), follow the installation steps.
2. Clone the git repo of this project to a file directory in your computer.
3. Navigate to the project file directory using a terminal, run command _go run main.go_, approve internet access for the program if prompted.
4. Start sending POST request to the [endpoint] once the text **Server starts listening on port 3000!** appears on the terminal screen.
5. A file named _data.json_ should appear if the request were successful.
6. Press ctrl+C to stop the server and exit the program.
