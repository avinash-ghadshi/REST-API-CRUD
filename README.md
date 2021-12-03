# REST-API-CRUD
Simple REST API application using GOLANG [/net/http &amp; mux]

#### Available URLs:

| URL | METHOD |
| - | - |
| http://127.0.0.1:8888/ | GET |
| http://127.0.0.1:8888/list/ | GET |
| http://127.0.0.1:8888/retrieve/{id} | GET |
| http://127.0.0.1:8888/add/ | POST |
| http://127.0.0.1:8888/delete/{id} | DELETE |
| http://127.0.0.1:8888/update/ | PUT |


> Note: `{id} refers to the aadharNumber in this application. Do not put bracket in URL.`
E.g. http://127.0.0.1:8888/retrieve/123456781290`

**To change IP and PORT in URL**
- *Edit main.go file*
Change ```http.ListenAndServe(":8888", router)```
 to ```http.ListenAndServe("YOUR_IP:YOUR_PORT", router)```
- *Save & Build the code*
```go build main.go router.go```
