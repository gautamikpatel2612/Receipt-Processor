# Receipt-Processor

Fetch Rewards: Take Home Assessment - Backend Engineering


***********************************************************************************************************************************


Project Overview


We're building a web service with two endpoints:

POST /receipts/process – Accepts receipt JSON, returns an ID.

GET /receipts/{id}/points – Returns the calculated points for a receipt by its ID.

Use the Go standard library (no frameworks).

Store data in memory using a map[string]Receipt.

Use uuid to generate unique IDs.

Use mux to handle API routes cleanly and use path parameters like {id}.

Use net/http for routing.


***********************************************************************************************************************************


Project Does :


Accepts a JSON receipt with information like store name, date, time, and items.

Calculates points using rules (e.g. based on item descriptions, purchase time, etc.).

Returns a unique ID for the receipt when submitted.

Lets you check how many points a receipt earned using that ID.

Stores all data in memory – no database needed.


	FetchRewards /


		main.go 		 
		handlers.go	
		models.go		
		points.go		  
		go.mod			  
		go.sum			 
		Dockerfile		
		README.md 		


***********************************************************************************************************************************


STEP 1 :  
 
  go.mod , go.sum — Initialize Go Module
 
  go mod init receipt-processor

  go get github.com/google/uuid  : To generate unique IDs for receipts.
 
  go get github.com/gorilla/mux  : To handle API routes cleanly and use path parameters like {id}.


***********************************************************************************************************************************


STEP 2 :  
  
  models.go — Define Data Structures


***********************************************************************************************************************************


STEP 3 : 
 
  handlers.go — Define Handlers


***********************************************************************************************************************************


STEP 4 :
  
  points.go — Calculate Points Logic


***********************************************************************************************************************************


STEP 5 :
  
  main.go — Entry Point


***********************************************************************************************************************************


STEP 6 :
  
  Dockerfile — Container Setup


***********************************************************************************************************************************


STEP 7 : Run the App Locally(optional)
 
  Using Go :
    
    go run main.go handlers.go models.go points.go
    Server will be on: http://localhost:8080


***********************************************************************************************************************************


STEP 8 : 
  
  Using Docker :
    
    docker build -t receipt-processor .
    docker run -p 8080:8080 receipt-processor


***********************************************************************************************************************************


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


***********************************************************************************************************************************


STEP 10.1 : EX: 1  
	
 POST url : http://localhost:8080/receipts/process
	
 Set Headers : 
 
 Key: Content-Type  Value: application/json
 
 Set Body : 
 
 1. Select raw : (select JSON instead of TEXT),
 
 2.  Write a JSON file
 
 3. Click on Send


![image](https://github.com/user-attachments/assets/d8fbee7c-b880-4938-9f6d-a57cdece79df)


GET url : http://localhost:8080/receipts/{id}/points

![image](https://github.com/user-attachments/assets/3c871f79-3c3f-4344-ba91-7c4e53f6f16f)


***********************************************************************************************************************************


STEP 10.2  : EX: 2 

POST url : http://localhost:8080/receipts/process
 
Set Headers :  

Key: Content-Type  Value: application/json
	
Set Body : 

1. Select raw : (select JSON instead of TEXT),

2.  Write a JSON file

3. Click on Send


![image](https://github.com/user-attachments/assets/3e64a14c-8dce-4e57-808a-c8d07883cbab)


GET url : http://localhost:8080/receipts/{id}/points


![image](https://github.com/user-attachments/assets/718c51dc-0e3b-4647-bf8b-c829937eaa0b)


***********************************************************************************************************************************


Points Logic:

+1 point for every alphanumeric character in the retailer name

+50 points if the total is a round dollar amount with no cents

+25 points if the total is multiple of 0.25

+5 points for every two items on the receipt.

For items with trimmed description length divisible by 3
          Add ceil(item.price * 0.2) points
	  
+6 points if the purchase day is odd

+10 points if the purchase time is between 2:00 pm and 4:00pm (exclusive)


***********************************************************************************************************************************


Test Scenario:

Scenario 1 : If we are not adding any retailer then we will get “Invalid receipt: retailer is required”

![image](https://github.com/user-attachments/assets/86325367-8627-4c00-bdd2-8d48e7b935f6)


***********************************************************************************************************************************


Scenario 2 : If we are adding wrong format of purchaseDate then we will get “Invalid receipt: invalid purchase date format (expected YYYY-MM-DD)”

![image](https://github.com/user-attachments/assets/aec9ae21-befd-44bd-9769-73c626d38fbe)


***********************************************************************************************************************************


Scenario 3 : If we are not adding wrong format of purchaseTime then we will get “Invalid receipt: invalid purchase time format (expected HH:MM in 24-hour format)”

![image](https://github.com/user-attachments/assets/b5d27a57-04c1-4f62-973f-53f84e2c96e6)


***********************************************************************************************************************************


Scenario 4 : In all these Scenarios if we do GET REQUEST then we  will get output as “Method not allowed”.

![image](https://github.com/user-attachments/assets/70bfc06f-1f0c-4531-89a4-c6e8cfcbcdc0)
