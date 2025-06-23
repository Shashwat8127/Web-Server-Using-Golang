# Web-Server-Using-Golang
I recently built a simple yet functional web server using Go (Golang) to strengthen my understanding of the net/http package. The server runs on port 8080 and handles multiple routes. The root route (/) displays a welcome message, while /form.html serves a contact form designed with modern HTML and CSS styling. When the form is submitted, the /form route processes the POST request, parses the submitted data using ParseForm(), and dynamically displays the user’s name and address. Additionally, I created a /hello route that handles GET requests and returns a friendly greeting. The server gracefully handles unsupported paths and methods by returning appropriate 404 errors. This project helped me practice Go’s HTTP routing, serving static files with http.ServeFile, handling form data securely, and implementing basic error handling. It’s a great foundation for building more advanced backend applications with Go.

# Simple Go Web Server 🌐

This is a basic web server built using **Golang**. It demonstrates:
- Serving static HTML files (like `form.html`)
- Handling **GET** and **POST** requests
- Parsing form data
- Routing to different URLs

---

## 🚀 Features
✅ Root route `/` — Displays a welcome message  
✅ `/form.html` — Serves a styled contact form  
✅ `/form` — Handles form submissions (POST) and displays submitted data  
✅ `/hello` — Responds to GET requests with `Hello!`  
✅ Graceful 404 handling for unsupported routes  

---

## 📂 Project structure

go-simple-webserver/
├── main.go
└── form.html

## 💻 How to run

go run main.go

3️⃣ Open in browser:

http://localhost:8080/ → Welcome message

http://localhost:8080/form.html → Contact form

http://localhost:8080/hello → Hello route

🌱 What I learned
Setting up a basic Go web server

Creating custom route handlers

Parsing form data using ParseForm()

Serving static files using http.ServeFile

Basic error handling for HTTP servers

DEMO / OUTPUT:
Coming soon / or run locally on localhost:8080.

