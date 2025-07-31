# Tasks for Strata
## Task 1
* Reviewed JSON file and added correction under "Corrected version.json" to make it valid

## Task 2
* Setup a HTTP GO server with a GET and POST endpoint
* There is a check to make sure the HTTP request is valid (not allowing PUT on POST for example)

**How to Run**  
1. Download main.go (https://github.com/ggnngit/Strata/blob/main/Task%202/httpserver/main.go)
2. In your terminal, run the following command
   > go run main.go
3. You can run the following on port 8080
   > curl --location --request GET 'localhost:8080/echo' \
      --header 'Content-Type: application/json' \
      --data '{
        "name": "test"
      }'
   
   >curl --location 'localhost:8080/ping'
4. CTRL + C will stop the program
## Task 3
* Setup an GO API client that retrieves a set of users

**How to Run**  
1. Download API Client.go (https://github.com/ggnngit/Strata/blob/main/Task%203/API%20Client.go)
2. In your terminal, run the following command
   > go run API\ Client.go
   
3. The program will return a JSON response

## Task 4
1. Install Maverics Orchestrator (https://docs.strata.io/docs/installation-overview)
2. Clone this repo: https://github.com/ggnngit/Strata to yours
3. Create your Github storage: https://docs.strata.io/docs/github-3
4. Start the orchestrator
5. Navigate to https://localhost/anything
6. You'll get a redirect to Okta for authentication
7. It will send you to https://httppbin.org/anything post authentication
