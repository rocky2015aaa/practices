# xml Conversion Server (to json)

**1. Exceptional Field Data type (not correct data type)**

    houseNumber : string - The source(xml) value is CDATA. CustomInteger type is used for int value.

    houseNumberAddtion : string - The source value is empty.

    minprice : string - The source value is CDATA. CustomInteger type is used for int value.
    
    Projectnaam : string - The source value is empty.
    
    WoningtypeInProject : string - The source value is empty.
    
    pdf : string - The source value is empty.
    
    contractLentgh_months : string - The source value is empty. And the source has a typo. (Lentgh -> Length) CustomInteger type is used for int value.
    
    minContractLentgh_months : string - The source value is empty. And the source has a typo. (Lentgh -> Length) CustomInteger type is used for int value.
    
    buildYear : string - The source value is empty. CustomInteger type is used for int value.
    
    gardenLigging : string - The source value is empty.
    
    roofTerrassLigging : string - The source value is empty.
    
    BalconyLigging : string - The source value is empty.
    
**2. Functions**

    init : set logs and call helper importXMl function to import xml in advance
    
    GetServer : set router for server api and return it
    
    MainApi : convert xml data to json. set ok status code and print json and info log
              
    importXml : open xml file, read it and unmarshal it
    
    UnmarshalXML(method - CustomTime) : parse date format "yyyy-mm-dd"
    
    UnmarshalXML(method - CustomInteger) : convert cdata to int or set 0 for empty value

    ServeHTTP(method) : run the api by path and if it is not found, 
                        then print proper code and warning log
    
    
**3. Server API Information**

    Port number : 8080	

    / : return json from xml data


