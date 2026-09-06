module github.com/luckybaraka/snippetbox

go 1.26.4

//a go.mode file make my project a module
//but what is a module? It makes it much easier to manage third party dependancies, and avoid supply-chain attacks and ensure reproducable builds of my application in future.

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/go-sql-driver/mysql v1.10.0 // indirect
)
