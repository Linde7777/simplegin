# simplegin
simplegin imitating Gin

# Start
start from reading simplegin/simplegin.go first

`context` seal the `http.ResponseWriter` and `*http.Request`, 
provide a direct way to access path, method, status code, 
simplify the usage of original `http` package's function

`Router`

main.go is a example to illustrate how to use it.