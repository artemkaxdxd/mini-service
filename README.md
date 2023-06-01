# Image service

Endpoints:
- POST /login
```json
{"username": "test", "password": 123456}
```
Response: 
```json
{"token": "jwt string"}
```
User can log in (assuming that the user already has a record in database) and recieve a jwt token that is valid for 12 hours. 

All logins are saved in table **users**:
```
[id, username, password_hash]
```
- POST /upload-picture 

User can upload a photo, save it in a file, and the url is saved to table **images**
```
[id, user_id, image_path, image_url]
```
- GET /images
User can get an array with all of his images

Service has a middleware which validates a JWT token which was recieved by user after authorization. Token is passed in header Authorization: Bearer {jwt token here}
