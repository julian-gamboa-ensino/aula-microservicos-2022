# Use a imagem base do JMeter
FROM justb4/jmeter:latest

# Copie seus scripts de teste para o container
COPY ./scripts /jmeter/scripts

# Exponha as portas necessárias para RMI e JMeter
EXPOSE 60000 1099

# Comando para iniciar o JMeter em modo não GUI
CMD ["-n", "-t", "/jmeter/scripts/meu_teste.jmx", "-l", "/jmeter/scripts/resultados.jtl", "-Djava.rmi.server.hostname=0.0.0.0", "-Dserver_port=60000", "-Dserver.rmi.localport=1099", "-Dserver.rmi.ssl.disable=true"]