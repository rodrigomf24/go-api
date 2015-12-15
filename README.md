# my go learning examples

I took this example from http://thenewstack.io/make-a-restful-json-api-go/, instead of working with a todo list, I decided to use item's list.

Example of POST request to items:

curl -H "Content-Type: application/json" -d '{"name":"New Item"}' http://localhost:8080/api/item


API endpoints:
[
	{
		url: '/api/item/{id}',
		description: 'perform this request to get an item by id',
		method: 'GET'
	},
	{
		url: '/api/item',
		description: 'perform this request to add a new item record',
		method: 'POST'
	},
	{
		url: '/api/items',
		description: 'perform this request to get all items' 
	},
	{
		url: '/',
		description: 'webserver index page' 
	}
]