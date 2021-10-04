# Microservices as Firebase Authentication  

This is a POOC using [Firebease Auth](https://firebase.google.com/docs/auth "Firebease Auth's page") to authenticate request for own microservice infrastructure.

For startups or some fast projects , we prefer to delegate the authentication logic to a third-service to integrate social networks easier and maintain a microservice infrastructure. 

![Flow firebase authentication](/flow-diagram.png?raw=true "Optional Title")

For this example, I won't use environment variables.
I've seen this pattern in projects of big companies, and I've used it in personal projects with excelent results. 

## Nginx config
This is optional, and it depends on your architecture

We'll use [Nginx's Auth Request module](http://nginx.org/en/docs/http/ngx_http_auth_request_module.html "Nginx's Auth Request module") , this module is used to authenticate request.

```nginx
auth_request /auth;
```

In the file [nginx.conf](/nginx.conf), inside this repository, you can see an example of nginx config using this parameter. 

with this Nginx configuration, if the auth-ms respond a 4xx response, it denied the request, and if the auth-ms respond 2xx, it re-direct the request with our customs headers to the right microservice.

## Auth microservice
This is the most interesant and simplest thing of this.   
Because it works only with one file, I did an example in flask and tornado, but tornado was faster than Flask so i decided to upload the tornado's example. 

In the main.py file, it's important this.

```python
filejson_route = "adminsdk.json"
```
Here we have to set the firebase's config file. the configuration throught environment variables exists too you can read the [firebase documentation](https://firebase.google.com/docs/admin/setup/#python).

In this solution , as you can see the authentication request is successful, i add the **firebase user id** in the headers. This is for identify each request with its own user and identify in another microservices.

The treatment of this data in database and more things is topic about microservice architecture. 