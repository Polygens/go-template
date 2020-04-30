module github.com/polygens/{{cookiecutter.project_name}}

go 1.14

require (
	{% if cookiecutter.kafka|lower != "no" -%}
	github.com/Shopify/sarama v1.26.1
	{%- endif %}
	{% if cookiecutter.postgress|lower != "no" -%}
	github.com/go-pg/pg/v9 v9.1.6
	{%- endif %}
	github.com/go-playground/validator/v10 v10.2.0
	github.com/gorilla/mux v1.7.4
	github.com/prometheus/client_golang v1.5.1
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/viper v1.6.3
)
