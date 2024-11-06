using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using System.Text.Json;
using Otimizar.Models; 

namespace Otimizar.Services
{
    public class RabbitMQConsumerService
    {
        private readonly IConnection _connection;
        private readonly IModel _channel;

        public RabbitMQConsumerService(string hostname = "rabbitmq")
        {
            // Cria a conexão com o RabbitMQ
            var factory = new ConnectionFactory() { HostName = hostname };
            _connection = factory.CreateConnection();
            _channel = _connection.CreateModel();

            // Declara a fila de onde vamos consumir as mensagens
            _channel.QueueDeclare(queue: "fila_otimizar", durable: false, exclusive: false, autoDelete: false, arguments: null);
        }

        public void StartConsuming()
        {
            // Cria o consumidor que irá escutar as mensagens da fila
            var consumer = new EventingBasicConsumer(_channel);

            consumer.Received += (model, ea) =>
            {
                try
                {
                    var body = ea.Body.ToArray();
                    var message = Encoding.UTF8.GetString(body);

                    // Exibe a mensagem completa no console para depuração
                    Console.WriteLine("----------- Mensagem Recebida ------------");
                    Console.WriteLine(message);

                    // Tenta desserializar a mensagem em um objeto EntradaPedidos
                    var entradaPedidos = JsonSerializer.Deserialize<EntradaPedidos>(message);

                    // Verifica se a desserialização foi bem-sucedida
                    if (entradaPedidos != null && entradaPedidos.Pedidos != null)
                    {
                        foreach (var pedido in entradaPedidos.Pedidos)
                        {
                            Console.WriteLine("Pedido recebido para processamento:");

                            
                            Console.WriteLine($"PedidoId: {pedido.PedidoId}");
                            foreach (var produto in pedido.Produtos)
                            {
                                Console.WriteLine($"ProdutoId: {produto.ProdutoId}, Dimensões: {produto.Dimensoes.Altura}x{produto.Dimensoes.Largura}x{produto.Dimensoes.Comprimento}");
                            }
                        }
                    }
                    else
                    {
                        Console.WriteLine("Erro: A mensagem recebida não pôde ser desserializada em um objeto EntradaPedidos.");
                    }

                    // Confirma que a mensagem foi processada com sucesso
                    _channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                }
                catch (Exception ex)
                {
                    // Log de erro no caso de falha de processamento
                    Console.WriteLine("Erro ao processar a mensagem:");
                    Console.WriteLine(ex.Message);

                    // Opcional: NACK se quiser que a mensagem retorne à fila em caso de erro
                    //_channel.BasicNack(deliveryTag: ea.DeliveryTag, multiple: false, requeue: true);
                }
            };

            // Consome da fila com autoAck desativado para confirmação manual
            _channel.BasicConsume(queue: "fila_otimizar", autoAck: false, consumer: consumer);
            Console.WriteLine("Aguardando mensagens na fila 'fila_otimizar'...");
        }
    }
}
