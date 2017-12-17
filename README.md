# install the dependency locally
go get ./...

# Run the program
'''
go run restApiExample.go
'''

# Hit the GET api given below with key
'''
http://localhost:8000/<key>
 eg. http://localhost:8000/Nandeshwar
 output: not found
'' 
 
 
# Hit the PUT api given below with key and body with value
# if key is present, value will be updated, otherwise key value pair will be created.
'''
http://localhost:8000/<key>
   body - 1000
eg. in postman - PUT
  http://localhost:8000/Nandeshwar
    body - 1000
    
 ouput: 
    created
'''
    
