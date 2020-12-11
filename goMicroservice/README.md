# goMicroservice

#### *- Test Environment*
- Golang version higer than 1.6
- Consul
- Godep
- Postman or curl

#### *- Test Process*
- Download dependent packages with Godep

Godep information
~~~~
https://github.com/tools/godep
~~~~
- Run Cunsul

~~~~
consul agent -dev -advertise=127.0.0.1
~~~~
- Execute **run.sh** to start services (app, auth, db, router)
- Use postman
	1. open postman

	2. set method "GET", address "http://localhost:8000", header with token

	3. press send for test

- Use curl

~~~~
curl -X GET -H "authentication_token: 1q2w3e4r" "http://localhost:8000"
~~~~
#### *- Main Functions*
1. *handler* in **router**
  	* Parameters : *http response writer, http request*
  	* Description : *The handler function to get todo data*
  		1. Check if request is correct
  		2. Get the request to load todo data
  		3. Fetch the todo data


2. *GetTodoList* in **app**
	* Parameters : *context, auth service request, db service response*
	* Return values : *error*
 	* Description : *The function to return data*
 		1. Check authentication from auth service
 		2. Get data from db service


3. *Authenticate* in **auth**
 	 * Parameters : *context, auth service request, db service response*
 	 * Return values : *error*
 	 * Description : *The function to check authentication with token*
 	 	1. Check token
 	 	2. Set username


4. *LoadTodoList* in **db**
	* Parameters : *context, auth service response (request), db service response (response)*
	* Return values : *error*
	* Description : *The function to get todo data*
		1. Load data by accessing db service
		2. Set the data format

#### *- Unit Test*
- *TestLoadData* in **db**
	* Description : *Check if loadData function works properly. Actually checks if reading json file well*
