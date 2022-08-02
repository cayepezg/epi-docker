
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

## Regeneración de swagger.yaml
Se debe ejecutar el siguiente comando, en caso de que se realicen actualizaciones en el código, que afecten la documentación del mismo
```bash
~/go/bin/swagger generate spec -o ./view/swagger.yaml --scan-models
~/go/bin/swagger mixin \
		./view/init.yml \
		./view/swagger.yaml \
		--format yaml \
		-o ./view/swagger.yaml
```

go mod init personas