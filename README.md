# WebApp in Go(lang)

## What is this

I am learning Go, in particular I am following along Trevor Sawler's lessons to get an idea of:

- how to setup a webserver
- how to install packages
- how to use/create middleware (using chi for routing)

## Why did you push?

- I want to have a small reference app with the basics of a web server developed in go

- Trevor uses the Repository pattern to keep data that needs to be used across the app (and I never came across it)

- I have left a lot of comments in the code for future reference

## Technologies

- Go 1.15
- [Chi](https://github.com/go-chi/chi) router
- [SCS](https://github.com/alexedwards/scs/v2) Session Manager
- [Nosurf](https://github.com/justinas/nosurf) for CSRF protection
