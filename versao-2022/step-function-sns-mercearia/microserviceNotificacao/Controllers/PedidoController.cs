using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using System.Text.Json;

namespace Notificacao.Services
{
    public class RabbitMQConsumerService
    {
        private readonly IConnection _connection;
        private readonly IModel _channel;

        public RabbitMQConsumerService(string hostname = "rabbitmq")
        {
            var factory = new ConnectionFactory() { HostName = hostname };
            _connection = factory.CreateConnection();
            _channel = _connection.CreateModel();
            _channel.QueueDeclare(queue: "fila_resultado", durable: false, exclusive: false, autoDelete: false, arguments: null);
        }

        public void StartConsuming()
        {
            var consumer = new EventingBasicConsumer(_channel);
            consumer.Received += (model, ea) =>
            {
                var body = ea.Body.ToArray();
                var message = Encoding.UTF8.GetString(body);
                var resultado = JsonSerializer.Deserialize<List<ResultadoCaixa>>(message);

                Console.WriteLine("Resultado recebido para notificação:");
                foreach (var caixa in resultado)
                {
                    Console.WriteLine($"Caixa: {caixa.CaixaId}, Produtos: {string.Join(", ", caixa.Produtos)}");
                }

                _channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
            };

            _channel.BasicConsume(queue: "fila_resultado", autoAck: false, consumer: consumer);
            Console.WriteLine("Aguardando mensagens na fila 'fila_resultado'...");
        }
    }
}
