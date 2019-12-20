**Get all books**
----

* **URL**

  /library/books/

* **Method:**
  `GET`
  
* **Success Response:**
  
    Returns list of all Books in library

  * **Code:** 200 OK <br />
  * **Content:** `
    [
        {
            "Id": 0,
            "Name": "The Hobbit",
            "CurrentUser": null,
            "Issued": false,
            "IssuedTimes": 0
        },
        {
            "Id": 1,
            "Name": "The Fellowship of the Ring",
            "CurrentUser": null,
            "Issued": false,
            "IssuedTimes": 0
        }
        ]`
 
**Create book**
----

* **URL**

  /library/books/

* **Method:**
  `POST`
  
*  **JSON Body**
   `{   "Id": 0,
        "Name": "The Hobbit"
    }`

* **Success Response:**
  * **Code:** 201 Created <br />
 
**Delete book**
----

* **URL**

  /library/books/{id}/

* **Method:**
  `DELETE`

* **Success Response:**
  * **Code:** 200 OK <br />
 
* **Error Response:**
  * **Code:** 404 Not found <br />

**Get book**
----

* **URL**

  /library/books/{id}/

* **Method:**
  `GET`

* **Success Response:**
  * **Code:** 200 OK <br />
 
* **Error Response:**
  * **Code:** 404 Not found <br />
  
**Get most issued book**
  ----
  
  * **URL**
  
    /library/books/issue/most/
  
  * **Method:**
    `GET`
  
  * **Success Response:**
    * **Code:** 200 OK <br />
    * **Content:** 
    `{
         "Id": 3,
         "Name": "The Return of the King",
         "CurrentUser": {
             "Id": 2
         },
         "Issued": true,
         "IssuedTimes": 1
     }`
   
  * **Error Response:**

  **Issue book to user**
  ----
  
  * **URL**
  
    /library/books/{bookId}/issue/{userId}/
  
  * **Method:**
    `POST`
  
  * **Success Response:**
    * **Code:** 200 OK <br />
   
  * **Error Response:**
    * **Code:** 404 Not found <br />
    * **Code:** 409 Conflict <br />
    
  **Return book to library**
  ----
  
  * **URL**
  
    /library/books/{bookId}/return/
  
  * **Method:**
    `POST`
  
  * **Success Response:**
    * **Code:** 200 OK <br />
   
  * **Error Response:**
    * **Code:** 404 Not found <br />
    * **Code:** 409 Conflict <br />


  **Check if book is issued**
  ----
  
  * **URL**
  
    /library/books/{bookId}/issued/
  
  * **Method:**
    `GET`
  
  * **Success Response:**
    * **Code:** 200 OK <br />
    * **Content:** 
    `
        {
            "issued": false
        }
    `
   
  * **Error Response:**
    * **Code:** 404 Not found <br />

  **Get all users**
  ----
  
  * **URL**
  
    /library/users/
  
  * **Method:**
    `GET`
  
  * **Success Response:**
    * **Code:** 200 OK <br />
    * **Content:** 
    `
        [
            {
                "Id": 2
            },
            {
                "Id": 0
            }
        ]
    `
    
  **Create user**
  ----
  
  * **URL**
  
    /library/users/
  
  * **Method:**
    `POST`
  
  *  **JSON Body**
    `
        {
                "Id": 1
        }
    `
  
  * **Success Response:**
    * **Code:** 201 CREATED <br />
    
  **Get user**
  ----
  
  * **URL**
  
    /library/users/{userId}/
  
  * **Method:**
    `GET`
  
  * **Success Response:**
    * **Code:** 201 CREATED <br />
    * **Content:** 
        `
            {
                 "Id": 2
            }
        `
        
  * **Error Response:**
    * **Code:** 404 Not found <br />