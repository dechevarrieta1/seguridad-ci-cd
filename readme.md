## Pasos de ejecución del código

##### 1.1 Se realiza la ejecución del main.
 
##### 1.2 Se llama a la ejecución de todos los handlers hacia la url definida por variables de entorno.
 
##### 1.3 Se loggean todas las respuestas ya sean de éxito u error.
 
##### 1.4 Finaliza el programa.

## Pasos de ejecución del CICD

##### 1.1 "Checkout code": Este paso utiliza la acción actions/checkout@v3 para clonar el repositorio en la máquina virtual de GitHub Actions.
 
##### 1.2 "Set up Go": Este paso utiliza la acción actions/setup-go@v3 para configurar el entorno de Go en la máquina virtual. Se especifica la versión de Go a utilizar (en este caso, la versión 1.22).
 
##### 1.3 "Install dependencies": Este paso ejecuta el comando go mod tidy para instalar las dependencias del proyecto. El comando go mod tidy analiza el archivo go.mod y descarga las dependencias necesarias.
 
##### 1.4 "Build": Este paso ejecuta el comando go build -v ./... para compilar el código del proyecto. El comando go build compila el código fuente y genera un ejecutable.
 
##### 1.5 "Test": Este paso ejecuta el comando go test -cover ./... para ejecutar las pruebas unitarias del proyecto. El comando go test busca y ejecuta las pruebas en los archivos con el sufijo _test.go.
 
##### 1.6 "Run server": Este paso ejecuta el comando go run main.go para iniciar el servidor. También se establece una variable de entorno llamada SEGURIDAD_URL que se obtiene de los secretos de GitHub. El servidor se ejecuta en segundo plano utilizando nohup y se redirige la salida a un archivo llamado server.log.
 
##### 1.7 "Check server log": Este paso muestra el contenido del archivo server.log en la salida de GitHub Actions.