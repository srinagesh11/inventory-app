FROM mysql:latest

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./db/*.sql /docker-entrypoint-initdb.d/