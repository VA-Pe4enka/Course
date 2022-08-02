
## Установка программы migrate  
`
$go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
`
## Установка программы sqlc  
`
$go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
`
  
## Установка программы pgAdmin4  
`
$curl https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add
`  
`
$sudo sh -c 'echo "deb https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update'
`  
#### Затем запустите установку pgAdmin4:  
`
$sudo apt install pgadmin4
`
