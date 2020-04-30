# {{cookiecutter.project_name}}
> Created using cookiecutter [template](https://github.com/Polygens/go-template)

![Tests](https://github.com/Polygens/{{cookiecutter.project_name}}/workflows/Test/badge.svg)
![Build](https://github.com/Polygens/{{cookiecutter.project_name}}/workflows/Build/badge.svg)
![Release](https://github.com/Polygens/{{cookiecutter.project_name}}/workflows/Test,%20Build%20and%20Publish/badge.svg)

{{cookiecutter.project_short_description}}

## Configuration

<!-- START CONFIG -->
<!-- END CONFIG -->

{% if cookiecutter.kafka|lower != "no" -%}
## Kafka

{% if cookiecutter.kafka|lower in ["consumer", "both"] -%}
### Input

<!-- START KAFKA_INPUT -->
<!-- END KAFKA_INPUT -->
{%- endif -%}

{% if cookiecutter.kafka|lower in ["producer", "both"] -%}
### Output

<!-- START KAFKA_OUTPUT -->
<!-- END KAFKA_OUTPUT -->
{%- endif -%}

{%- endif -%}

{% if cookiecutter.postgress|lower != "no" -%}
## Postgress

### Data

{% endif %}
