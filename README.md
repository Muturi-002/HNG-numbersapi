# HNG-numbersapi

## Introduction
This is a program that takes an integer input and provides an output of its interesting properties, and a fun fact about it.

The output/response is provided in a JSON format, for both '200 OK' and '400 Bad Request' Responses.

## Requirements
### Technology stack
- Programming language: Go
- Must be deployed to a publicly accessible endpoint
- Must handle CORS (Cross-Origin Resource Sharing)
- Must return responses in JSON format


#### For a JSON Required Format (200 OK)
``` json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,  // sum of its digits
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```
#### For a JSON Required Format (400 Bad Request)
``` json 
{
    "number": "alphabet",
    "error": true
}
```