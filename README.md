# echo-quote
quote server using labstack/echo (Golang Web Framework Minimalist)
## Dev Mode
. install gin

`go get github.com/codegangsta/gin`
. run with dev mode

`gin -p 8080`

## Deploy to heroku
. Install heroku-cli,
. Create heroku in this directory

`heroku create`

. Set environment for mongodb
`heroku config:set MLAB_URL=yourmlaburl`
