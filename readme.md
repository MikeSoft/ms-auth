# Firebase Authentication as an auth-layer 

In this repository, I show how to use [Firebease Auth](https://firebase.google.com/docs/auth "Firebease Auth's page") for authenticate requests inside a microservices infraestructure.


For startups or some fast projects , we preffer to delegate the authentication functionalities to another entity like Firebase. for this cases could be necessary to integrate the user data with our business logic. 

![Flow firebase authentication](/flow-diagram.png?raw=true "Optional Title")

For this example, I won't use environment variables. ( For the purists )

I've seen this pattern some projects of big companies and i've used it in personal projects with excelent results. 


## Nginx config
This is optional, and it depends of your architecture, for a MVP in a own project that i use django, i built an middleware. but for this example i'll explain how to di it with nginx , because it's more flexible and scalable solution.  ( and expensive )

We'll use [Nginx's Auth Request module](http://nginx.org/en/docs/http/ngx_http_auth_request_module.html "Nginx's Auth Request module") , this module is used to authenticate request.

```nginx
auth_request /auth;
```

In the file [nginx.conf](/nginx.conf), inside this repository, you can see an example of nginx config using this parameter. 

## Auth microservice
This is the most interesant and simplest thing of this.   
Because it works only with one file, i did an example in flask and tornado, but tornado was faster than Flask so i decided to upload the tornado's example. 
  
I like to experiment with technologies and proof of concept (like this that now it's used in big companies) and i tried to make a functional microservice with golang, but i'll upload an example sooner with several languages and frameworks.

In the main.py file, it's important this.

```python
filejson_route = "adminsdk.json"
```
Here we have to set the firebase's config file. the configuration throught environment variables exists too you can read the [firebase documentation](https://firebase.google.com/docs/admin/setup/#python).

In this solution , as you can see the authentication request is successful, i add the **firebase user id** in the headers. This is for identify each request with its own user. 

The treatment of this data in database and more things is topic about microservice architecture. 