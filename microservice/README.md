# microservice

#### *- Test Environment*
- Golang version higer than 1.6
- Jwt-go
- Godep
- Web browser

#### *- Test Process*
- Download dependent packages with Godep

Godep information
~~~~
https://github.com/tools/godep
~~~~
1. Execute **run.sh** to start services (app, auth, db, router)
2. Set user authentication

	~~~~
	e.g : http://localhost:8001/signup?username=user1
	~~~~
3. Login app

	~~~~
	http://localhost:8000
	~~~~
4. Logout by accessing auth service

	~~~~
	http://localhost:8000/logout
	~~~~

#### *- Other Process*
- Check user authentication information after registering user authentication

	~~~~
	http://localhost:8000/auth
	~~~~
- Check all db data

	~~~~
	http://localhost:8000/db
	~~~~

#### *- Important Functions*
1. *generateToken* in **auth**
	* Parameters : *user name*
	* Return values : *http.Cookie*
 	* Description : *the function to make authentication token*
 		1. With claim information(user name, expiration, issuer), create token from jwt library.
 		2. Make the signed token with specifit key
 		3. Return cookie information with signed token


2. *validate* in **auth**
 	 * Parameters : *http handler function*
 	 * Description : *the function to authenticate the user and connect the following behavior*
 	 	1. Get the cookie information
 	 	2. Parse claims information from signed token
 	 	3. Put claims information into context for the following behavior


3. *protectedData* in **db**
	* Parameters : *http response writer, http request*
	* Description : *the function to get todo list*
		1. Check if the user is authenticated
		2. Get all db data


4. *ServeHTTP* in **util**
	* Parameters : *http response writer, http request*
	* Description : *To check if the request path is accurate, use map container and with valid path, call proper http handlers*

#### *- Unit Test*
1. *TestGenerateToken* in **auth**
	* Description : *check if generateToken function creates signed token successfully by checking token length*
2. *TestGenerateData* in **db**
	* Description : *check if generateData function works properly. Actually checks if reading json file well*
