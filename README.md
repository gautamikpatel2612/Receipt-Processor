# Receipt-Processor
Fetch Rewards: Take Home Assessment - Backend Engineering


Project Overview


We're building a web service with two endpoints:
POST /receipts/process – Accepts receipt JSON, returns an ID.
GET /receipts/{id}/points – Returns the calculated points for a receipt by its ID.
Use the Go standard library (no frameworks).
Store data in memory using a map[string]Receipt.
Use uuid to generate unique IDs.
Use mux to handle API routes cleanly and use path parameters like {id}.
Use net/http for routing.

Project Does :
Accepts a JSON receipt with information like store name, date, time, and items.
Calculates points using rules (e.g. based on item descriptions, purchase time, etc.).
Returns a unique ID for the receipt when submitted.
Lets you check how many points a receipt earned using that ID.
Stores all data in memory – no database needed.

Receipt-processor /
main.go 		  Start the server
handlers.go		Request handler logic
models.go		  Points calculation logic
points.go		  Data structures
go.mod			  Go modules
go.sum			  Go modules
Dockerfile		for containerization
README.md 		This file 



STEP 1 :  
  go.mod , go.sum — Initialize Go Module
  go mod init receipt-processor
  go get github.com/google/uuid  : To generate unique IDs for receipts.
  go get github.com/gorilla/mux  : To handle API routes cleanly and use path parameters like {id}.


STEP 2 :  
  models.go — Define Data Structures


STEP 3 : 
  handlers.go — Define Handlers



STEP 4 :
  points.go — Calculate Points Logic



STEP 5 :
  main.go — Entry Point



STEP 6 :
  Dockerfile — Container Setup



STEP 7 : Run the App Locally(optional)
  Using Go :
    go run main.go handlers.go models.go points.go
    Server will be on: http://localhost:8080



STEP 8 : 
  Using Docker :
    docker build -t receipt-processor .
    docker run -p 8080:8080 receipt-processor



STEP 9 : Test API with Postman (GUI)
  1. Submit Receipt – POST
    Method: POST
	  URL: http://localhost:8080/receipts/process
	  Headers:
		  Key: Content-Type | Value: application/json
	  Body: Choose raw → JSON and paste:



  Click Send
	Response: You’ll get { "id": "some-id" }

  2. Get Points – GET
    Method: GET
	  URL: http://localhost:8080/receipts/YOUR-ID-HERE/points
	  Replace YOUR-ID-HERE with the ID from the first call.
    Click Send
	  Response: { "points": 123 }


STEP 10 : EX: 1  
	POST url : http://localhost:8080/receipts/process
		Set Headers :  Key: Content-Type  Value: application/json
		Set Body :  1. Select raw : (select JSON instead of TEXT),
     	2.  Write a JSON file
     	3. Click on Send

	 https://github.com/gautamikpatel2612/Receipt-Processor/blob/main/1.png?raw=true
