## boilerplate go

### Login
req login, https://domain/login
```
{
  "Username":"",
  "Password":""
}
```

### Auth Token = Bearer JWT

### Create User
Method Post, https://domain/users
```
{
  "badge":"12345",
  "nama":"hahahah",
  "password":"123",
  "email":"tko@email.com",
  "role_id":1
}
```

### Get User
Method Get, https://domain/users?page=1&limit=10&badge=

### Services Includes
Rate Limiter,
Cors,


### Database Connection 
SqlServer
