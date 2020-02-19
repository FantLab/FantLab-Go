FROM perl:5.28
WORKDIR /app
COPY deps.txt .
RUN cat deps.txt | xargs cpanm --no-wget -n
