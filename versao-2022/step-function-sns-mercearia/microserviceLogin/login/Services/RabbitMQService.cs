using RabbitMQ.Client;
using System;
using System.Text;
using System.Threading;

namespace Login.Services
{
    public class RabbitMQService : IDisposable
    {
        private readonly IConnection _connection;
        private IModel _channel;
        private const int MaxRetries = 5; // Número máximo de tentativas de reconexão
        private const int DelayBetweenRetries = 2000; // Atraso entre as tentativas (milissegundos)

        public RabbitMQService(string hostname)
        {
            var factory = new ConnectionFactory() { HostName = "rabbitmq" };
            _connection = ConnectWithRetry(factory);
            _channel = _connection.CreateModel();
            _channel.QueueDeclare(queue: "fila_otimizar", durable: false, exclusive: false, autoDelete: false, arguments: null);

            Console.WriteLine("Conexão com RabbitMQ estabelecida com sucesso.");
        }

        public void SendMessage(string queue, string message)
        {
            var channel = GetChannel();
            var body = Encoding.UTF8.GetBytes(message);

            channel.QueueDeclare(queue: queue,
                                 durable: false,
                                 exclusive: false,
                                 autoDelete: false,
                                 arguments: null);

            channel.BasicPublish(exchange: "",
                                 routingKey: queue,
                                 basicProperties: null,
                                 body: body);

            Console.WriteLine($"[x] Mensagem enviada para a fila {queue}");
        }


        private IConnection ConnectWithRetry(ConnectionFactory factory)
        {
            int attempts = 0;
            while (attempts < MaxRetries)
            {
                try
                {
                    var connection = factory.CreateConnection();
                    Console.WriteLine("Conectado ao RabbitMQ com sucesso.");
                    return connection;
                }
                catch (RabbitMQ.Client.Exceptions.BrokerUnreachableException ex)
                {
                    attempts++;
                    Console.WriteLine($"Falha ao conectar ao RabbitMQ. Tentativa {attempts}/{MaxRetries}. Erro: {ex.Message}");

                    if (attempts >= MaxRetries)
                    {
                        throw new Exception("Máximo de tentativas de conexão ao RabbitMQ atingido. Verifique se o RabbitMQ está em execução.");
                    }

                    Thread.Sleep(DelayBetweenRetries);
                }
            }
            throw new Exception("Falha ao conectar ao RabbitMQ após várias tentativas.");
        }

        public IModel GetChannel()
        {
            return _channel;
        }

        public void Dispose()
        {
            _channel?.Close();
            _connection?.Close();
        }
    }

}