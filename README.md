** POC- API Endpoints Testing Service **

This a service that tests all the api endpoints present in a service using  Newman Commands and Concises the results into a simple readable table.

This takes the File path of the service that needs to be tested and analyzes all the enpoints present in Service using "go/parser"

There is a executable run file present in the service that we are testing - up's the container, pulls the newman and runs the commands and generates the newman report.

The executable run file present in api_testing service will up the docker contianer, copies the All the files of the service we are testing and runs the executable run file present in service we are testing and then generates table with API endpoints status- success, failure, Not Analyzed

This api_testing service will give a fine printed table with the All the API endpoints tested, which can save the time of a developer to check the status of API Endpoints.
