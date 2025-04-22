using Otimizar.Services; // Certifique-se de usar o namespace correto para o RabbitMQConsumerService

var rabbitMQConsumerService = new RabbitMQConsumerService("rabbitmq");
rabbitMQConsumerService.StartConsuming();

await Task.Delay(-1); // Mantém a aplicação em execução


