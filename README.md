# Fetch Rewards Assesment
This is a public repository containing the source code for Fetch Rewards. 
The task involves developing a server using Go and implementing two endpoints:

1. The /receipts/process endpoint accepts a POST request with a Receipt JSON object. The server parses the receipt and calculates the points based on a predefined set of rules. It responds with a JSON object containing the ID of the submitted receipt for future reference.

2. The /receipts/{id}/points endpoint retrieves the points associated with a previously submitted receipt. The receipt ID is provided as a parameter in the endpoint's path. The endpoint returns a JSON object containing the total points earned by the receipt.

The server is designed to handle various error scenarios, such as missing attributes in the JSON body of the POST request, errors in parsing date and time attributes, and internal server errors.

The code can be found in the fetch.go file.
The testing code can be found in fetch_test.go file.

How to reproduce the results and evaluate:

1. Clone the project repository using the command: git clone https://github.com/deadpanther/fetch_rewards_assesment
2. Open a terminal and navigate to the repository's directory.
3. Build the Docker image using the command: docker build -t fetch . This will create a Docker image with the tag receipt-api.
4. Run the Docker container from the image using the command: docker run -p 5001:5000 fetch
5. The Go application will now be running inside the Docker container. You can access and test the API using tools like POSTMAN or cURL by sending requests to http://localhost:5001.
6. For the /receipts/process endpoint, use a POST request to http://localhost:5001/receipts/process.
7. For the /receipts/{id}/points endpoint, use a GET request to http://localhost:5001/receipts/{id}/points.
8. To run the unit test we can use the terminal and run the following command: go test