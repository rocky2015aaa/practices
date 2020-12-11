# MERGE INTEGER SERVER

#### *- Process*
1. Parse url from each query parameter
2. Check that parsed url is valid
3. Make request for each valid url
4. Merge each result from url responses
5. Sort result with ascending order
6. Print result in json format

#### *- Functions*
1. MergeIntegers
  	* Parameters : *http response writer  **,** http request*
  	* Description : *The function to make result*
   	       - From request, parse "u" query
 	       - With url parse, add url port number to call
 	       - Asynchronized merge operation with go routine and sync.WaitGroup
2. isValidUrl
	* Parameters : *url*
	* Return values : *bool*	
 	* Description : *The helper function to check url is valid*
 		- Use regular expression to check url is valid
 3. mergeUrlInt
 	 * Parameters : *pointer of map to remove duplication of final result **,** pointer of final result **,** url path*
 	 * Description : *The helper function to send request to url and merge numbers with handling timeout*
 	 	- To handle timeout, use go routine and time.After function
 	 	- use map to remove result number duplication
4. callUrl
	* Parameters : *http get function **,** url **,** channel for result*
	* Description : *The helper function to handle http request for getting result*
		- get function parameter to apply mock unit test
		- use channel to hanle asynchronized merge operation

#### *- Unit Test*
1. TestisValidUrl
	* Description : *check if isValidUrl function works properly*
2. TestCallUrl
	* Description : *check if making http request to url is successful*
		- use mock http get request function
3. TestCallUrlTimeout
	* Description : *check if channel to handle timeout is working well with making http request*
		- use mock http get request function

#### *- Note*
1. To test url like "example.com" or "foobar.com" on local machine, set those domains with localhost in host file
2. The query example for the server doesn't use port number so add port number with domain before making http request
3. just print log information for error handling, not use http error to return because to return empty list when error is occured