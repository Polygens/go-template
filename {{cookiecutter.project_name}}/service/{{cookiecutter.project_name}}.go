package service

import (
	"github.com/gorilla/mux"

	"github.com/polygens/{{cookiecutter.project_name}}/service/config"
	"github.com/polygens/{{cookiecutter.project_name}}/service/http"
	{% if cookiecutter.kafka|lower != "no" -%}
	"github.com/polygens/{{cookiecutter.project_name}}/service/kafka"
	{%- endif %}
)

// Service is the service data container
type Service struct {
	cfg    *config.Config
	http 	*http.HTTP
	{% if cookiecutter.kafka|lower != "no" -%}
	kafka 	*kafka.Kafka
	{%- endif %}
}

// Setup creates and starts the service
func Setup(r *mux.Router, cfg *config.Config) *Service {
	svc := &Service{cfg, nil{% if cookiecutter.kafka|lower != "no"  -%}, nil{% endif %}}

	svc.http = http.Setup(r)
	{% if cookiecutter.kafka|lower != "no" -%}

	var err error
	svc.kafka, err = kafka.Setup(&cfg.Kafka)
	if err != nil {
		log.Fatalf("Kafka setup failed: %s", err)
	}
	{%- endif %}

	return svc
}

// Close is used to handle a gracefull shutdown of the service
func (svc *Service) Close() {
	{% if cookiecutter.kafka|lower != "no" -%}
	svc.kafka.Close()
	{%- endif %}
}
