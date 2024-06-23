package auth

import "github.com/gofiber/fiber/v2"

func AuthRouters(app fiber.Router) {
	app.Post("/register", Register)
	app.Post("/login", Login)
}

// {
// 	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkxNTgyNzYsImp0aSI6IjZlZDNlNjYwLTZiODgtNDUxNi1hOWY1LTkwYTUyZWY0ZGYxYSIsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyIjo2fQ.TqbtHM8VIDreTYCcz07gVUFeJM-Fm1uDS4ZQdVwRTUA",
// 	"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk0MTM4NzYsImp0aSI6IjFiYTgzMWQwLTRjZmYtNGE2OS05ZGQ3LWNmZTM2YjEwMjRhOCIsInR5cGUiOiJyZWZyZXNoIiwidXNlciI6Nn0.Lqh1RaXgDExN-iE0miRFg2hbx5x4Suj_pxHM79ftk0M"
//   }
