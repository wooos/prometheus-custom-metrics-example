colon := :

.PHONY: build
build:
	go build -o prometheus-custom-metrics-example main.go
	docker build -t wooos/prometheus-custom-metrics-example$(colon)1.0.0 ./
	docker push wooos/prometheus-custom-metrics-example$(colon)1.0.0