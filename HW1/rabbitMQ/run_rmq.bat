@echo We are running rabbitMQ in a docker deamon...
@echo ................................................... 
@echo To use rabbitMQ try to connect to localhost:5672 !
@echo Also you must have docker engine installed!
@docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3