# MVP GBM

### Instalación de Docker
Seguir los siguientes pasos:

1. Instalar la última versión del Docker:
```sh
curl -sSL https://get.docker.com/ | sudo sh
```

2. Verificar la instalación
```sh 
docker --version
```

3. Para usar Docker sin privilegios de root:
```sh
gpasswd -a $USER docker
```
*este comando se tiene que ejecutar como root


### Instalación de Elasticsearch

Seguir los siguientes pasos:

1. Descargar las imagen Elasticsearch:6.3.2 de Docker:
```sh
$ docker pull docker.elastic.co/elasticsearch/elasticsearch:6.3.2
```
2. Correr el contenedor Docker de Elasticsearch:
```sh
$ docker run --restart=always -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2
```
3. Verificar que este corriendo el contenedor de Docker:
```sh
docker ps -a
```
4. Instalar las variables de entorno:
```sh
echo 'export ELASTICSEARCH_HOSTS=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' elasticsearch`\n' >> ~/.bash_profile
echo 'export ELASTICSEARCH_PORT=9200\n' >> ~/.bash_profile
echo 'export ELASTICSEARCH_INDEX=gbm\n' >> ~/.bash_profile
echo 'export ELASTICSEARCH_TYPE=vehicle\n' >> ~/.bash_profile
echo 'export ELASTICSEARCH_ENTRYPOINT=http://$ELASTICSEARCH_HOSTS:$ELASTICSEARCH_PORT\n' >> ~/.bash_profile
echo 'export ELASTICSEARCH_USERNAME=elastic\n' >> ~/.bash_profile
echo 'export ELASTICSEARCH_PASSWORD=changeme\n' >> ~/.bash_profile
source ~/.bash_profile
```
4. Verificar que se este ejecutando el servidor de Elastisearch:
```sh
$ curl -u $ELASTICSEARCH_USERNAME:$ELASTICSEARCH_PASSWORD --header "content-type: application/JSON" -X GET  $ELASTICSEARCH_ENTRYPOINT/"_cluster/health?pretty"
```
Se debe de mostrar una salida como la siguiente:
```sh
{
  "cluster_name" : "docker-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 1,
  "active_shards" : 1,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```

### Instalación de Go

[Las instrucciones las puedes encontrar aquí](https://golang.org/doc/install)

Para instalar las dependencias del proyecto:

```sh
go get ./...
```

Para correr test:

```sh
go test -timeout 99999s ./... -p 1
```

Para compilar ejecutar el programa:

```sh
go run main.go
```

### Endpoints

Para verificar que está funcionando el servicio:

```sh
curl -XGET "http://localhost:8088/"
```

```sh
{
    "data":"Elasticsearch returned with code 200 and version 6.3.2"
}
```

Para registrar y obtener la posición de un vehículo:

```sh
curl -XGET "http://localhost:8088/geolocation/960-bmc"
```

```sh
{
    "latitude":19.4471,
    "longitud":-99.1599
}
```



Para buscar un registro por el identificador de un vehículo(666-xxx):

```sh
curl -XGET "http://localhost:8088/record?q=666-xxx"
```

```sh
{
	"data": [
        {
		    "vehicleid": "666-xxx",
		    "location": {
			    "lat": "19.4471",
			    "lon": "-99.1599"
		    },
		    "date": "1545170682684"
        }
    ]
}
```


Para obtener todos los registros:

```sh
curl -XGET "http://localhost:8088/record"
```

```sh
{
	"data": [
        {
		    "vehicleid": "666-xxx",
		    "location": {
			    "lat": "19.4471",
			    "lon": "-99.1599"
		    },
		    "date": "1545170682684"
        }, 
        {
		    "vehicleid": "960-bmc",
		    "location": {
			    "lat": "19.4471",
			    "lon": "-99.1599"
		    },
		    "date": "1545170688419"
	    }
    ]
}
```

### Referencias:

API de Geolocalización para pruebas:
https://ipdata.co/
https://api.ipdata.co/?api-key=test

