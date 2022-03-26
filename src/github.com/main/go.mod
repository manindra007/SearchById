module github.com/main

require (
github.com/details v0.0.0-1
github.com/search v0.0.0-1
github.com/userops v0.0.0-1
)

replace (
    github.com/details => ../details
    github.com/search => ../search
    
    github.com/userops => ../userops
)
go 1.18
