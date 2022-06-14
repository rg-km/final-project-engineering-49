# final-project-engineering-49 api specification

## **POST /login**

Login user

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Data Params**

```
 {
   "Email" : "email@mail.com",
   "Password" : "password123"
}
```

## **POST /register**

Register user

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Data Params**

```
 {
   "Name" : "johny",
   "Email" : "email@mail.com",
   "Password" : "password123",
   "ConfirmPassword" : "password123"
}
```

## **GET /user**

Returns user by token.

- **URL Params**  
  None
- **Data Params**  
  None
- **Headers**  
  Content-Type: application/json
  Token: string
- **Success Response:**
- **Code:** 200

## **GET /materi/:id**

Returns the specified materi by id.

- **URL Params**  
  _Required:_ `id=[integer]`
- **Data Params**  
  None
- **Headers**  
  Content-Type: application/json  
  Token: string
- **Success Response:**
- **Code:** 200
- **Error Response:**
  - **Code:** 404  
    **Content:** `{ error : "Materi not found" }`  
    OR
  - **Code:** 403  
    **Content:** `{ error : error : "Foribidden" }`

## \*\*GET /materis

Returns all Materi

- **URL Params**  
  None
- **Data Params**  
  None
- **Headers**  
  Content-Type: application/json  
  Token: string
- **Success Response:**
- **Code:** 200

- **Error Response:**
  - **Code:** 403  
    **Content:** `{ error : error : "Forbidden." }`

## **POST /materi**

Creates a new Materi

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
  Token: string
- **Data Params**

```
 {
    "Title" : "course 2",
   "Contain" : "desc course 2",
   "FileName": "course2.mp4",
   "Status": 1,
   "Creator": "Mr course2"
}
```

- **Success Response:**
- **Code:** 200

## **POST /materi/update/:id**

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
  Token: string
- **Data Params**

```
 {
   "ID" : 1,
   "Title" : "course 2",
   "Contain" : "desc course 2",
   "FileName": "course2.mp4",
   "Status": 1,
   "Creator": "Mr course2"
}
```

- **Success Response:**
- **Code:** 200

## **DELETE /materi/:id**

Deletes the specified user.

- **URL Params**  
  _Required:_ `id=[integer]`
- **Data Params**  
  None
- **Headers**  
  Content-Type: application/json  
  Token: string
