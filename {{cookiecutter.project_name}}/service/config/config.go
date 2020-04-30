package config

// Config contains all the configuration variables for this service
type Config struct {
	// LogLevel defines logging severity, can be trace, debug, info, warning, error, fatal and panic
	LogLevel string `validate:"required"`
	// HTTPort is the port on which all http endpoints will listen
	HTTPPort uint16 `validate:"required"`
	{% if cookiecutter.kafka|lower != "no" -%}
	// Kafka contains all configurations belonging to kafka
	Kafka    Kafka
	{%- endif %}
	{% if cookiecutter.postgress|lower != "no" -%}
	// Postgress contains all configurations belonging to Postgress
	Postgress Postgress
	{%- endif %}
}

{% if cookiecutter.kafka|lower != "no" -%}
// Kafka contains the config for kafka
type Kafka struct {
	// Version of the kafka broker
	Version             string   `validate:"required"`
	// Addresses of the kafka brokers
	Brokers             []string `validate:"gt=0,dive,hostname_port"`
	// ClientID is the registration id of this service towards Kafka
	ClientID            string   `validate:"lowercase,printascii"`
	{% if cookiecutter.kafka|lower in ["consumer", "both"] -%}
	// InputTopics contains the Kafka input topics
	InputTopics struct {
		// MyInputTopic string
	} `validate:"required,dive,required,lowercase,printascii"`
	{% endif -%}
	{% if cookiecutter.kafka|lower in ["producer", "both"] -%}
	// OutputTopics contains the Kafka output topics
	OutputTopics struct {
		// MyOutputTopic string   `validate:"lowercase,printascii"`
	} `validate:"required,dive,required,lowercase,printascii"`
	{%- endif %}
}
{% endif -%}

{% if cookiecutter.postgress|lower != "no" -%}
type Postgress struct {
	DBName string `validate:"required"`
	User string `validate:"required"`
	Password string `validate:"required"`
	Host string `validate:"hostname"`
	Port uint16
}
{% endif -%}
