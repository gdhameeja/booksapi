# BOOKS API

Write a simple BookShelf Web Application that has 2 users, one been admin and
other been regular user; admin user is responsible for adding books and regular user
can subscribe for the books:
Expose APIs for below operations using swagger yaml - https://editor.swagger.io
- Login user - Assume your system already has 2 users [ hardcode this into the DB with roles as admin/regular]
	- username
	- password
- Add book with below information - By admin user
	- book name
	- Author
	- Availability status
- Search book based on name and add filters for author and availability status
- Get book information based on the name
- Bulk API - To add books in bulk, should accept File as input with format as comma separated book info
