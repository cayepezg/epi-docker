
## Ejecución a partir del código
Ajustando los valores correctos de las variables, ejecutar el siguiente comando:
```bash
API_SERVER_PORT=puerto_a_servir \
	DB_HOST=ip_de_bd \
	DB_PORT=puerto_de_bd \
	DB_USER=usuario_de_bd \
	DB_PASSWORD=clave_super_segura \
	DB_NAME=nombre_de_base_de_datos \
	DB_SSL_MODE=disable \
	DB_MAXIDLECONNS=10 \
	DB_MAXOPENCONNS=10 \
	go run cmd/main.go
```

go mod init personas