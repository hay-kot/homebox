# Quick Start

## Docker Run

Great for testing out the application, but not recommended for stable use. Checkout the docker-compose for the recommended deployment.

```sh
docker run --name=homebox \
    --restart=always \
    --publish=3100:7745 \
    ghcr.io/hay-kot/homebox:nightly
```

## Docker-Compose

```yml
version: "3.4"
 services:
   homebox:
     image: ghcr.io/hay-kot/homebox:nightly
     container_name: homebox
     restart: always
     environment:
      - HBOX_LOG_LEVEL=info
      - HBOX_LOG_FORMAT=text
      - HBOX_WEB_MAX_UPLOAD_SIZE=10
     volumes:
       - homebox-data:/data/
     ports:
       - 3100:7745

volumes:
   homebox-data:
     driver: local
```

## Env Variables & Configuration

| Variable                 | Default                | Description                                                                        |
| ------------------------ | ---------------------- | ---------------------------------------------------------------------------------- |
| HBOX_MODE                | production             | application mode used for runtime behavior  can be one of: development, production |
| HBOX_WEB_PORT            | 7745                   | port to run the web server on, in you're using docker do not change this           |
| HBOX_WEB_HOST            |                        | host to run the web server on, in you're using docker do not change this           |
| HBOX_STORAGE_DATA        | /data/                 | path to the data directory, do not change this if you're using docker              |
| HBOX_STORAGE_SQLITE_URL  | /data/homebox.db?_fk=1 | sqlite database url, in you're using docker do not change this                     |
| HBOX_WEB_MAX_UPLOAD_SIZE | 10                     | maximum file upload size supported in MB                                           |
| HBOX_LOG_LEVEL           | info                   | log level to use, can be one of: trace, debug, info, warn, error, critical         |
| HBOX_LOG_FORMAT          | text                   | log format to use, can be one of: text, json                                       |
| HBOX_MAILER_HOST         |                        | email host to use, if not set no email provider will be used                       |
| HBOX_MAILER_PORT         | 587                    | email port to use                                                                  |
| HBOX_MAILER_USERNAME     |                        | email user to use                                                                  |
| HBOX_MAILER_PASSWORD     |                        | email password to use                                                              |
| HBOX_MAILER_FROM         |                        | email from address to use                                                          |
| HBOX_SWAGGER_HOST        | 7745                   | swagger host to use, if not set swagger will be disabled                           |
| HBOX_SWAGGER_SCHEMA      | http                   | swagger schema to use, can be one of: http, https                                  |

!!! tip "CLI Arguments"
      If you're deploying without docker you can use command line arguments to configure the application. Run `homebox --help` for more information.

      ```sh
      Usage: api [options] [arguments]

      OPTIONS
        --mode/$HBOX_MODE                                <string>  (default: development)
        --web-port/$HBOX_WEB_PORT                        <string>  (default: 7745)
        --web-host/$HBOX_WEB_HOST                        <string>
        --web-max-upload-size/$HBOX_WEB_MAX_UPLOAD_SIZE  <int>     (default: 10)
        --storage-data/$HBOX_STORAGE_DATA                <string>  (default: ./homebox-data)
        --storage-sqlite-url/$HBOX_STORAGE_SQLITE_URL    <string>  (default: ./homebox-data/homebox.db?_fk=1)
        --log-level/$HBOX_LOG_LEVEL                      <string>  (default: info)
        --log-format/$HBOX_LOG_FORMAT                    <string>  (default: text)
        --mailer-host/$HBOX_MAILER_HOST                  <string>
        --mailer-port/$HBOX_MAILER_PORT                  <int>
        --mailer-username/$HBOX_MAILER_USERNAME          <string>
        --mailer-password/$HBOX_MAILER_PASSWORD          <string>
        --mailer-from/$HBOX_MAILER_FROM                  <string>
        --swagger-host/$HBOX_SWAGGER_HOST                <string>  (default: localhost:7745)
        --swagger-scheme/$HBOX_SWAGGER_SCHEME            <string>  (default: http)
        --help/-h
        display this help message
      ```