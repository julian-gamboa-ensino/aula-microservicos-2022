using Microsoft.AspNetCore.Mvc;
using Login.Models;
using Login.Services;
using System.Text.Json;

namespace Login.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class PedidoController : ControllerBase
    {
        private readonly RabbitMQService _rabbitMQService;

        public PedidoController(RabbitMQService rabbitMQService)
        {
            _rabbitMQService = rabbitMQService;
        }

        [HttpPost("processar-pedidos")]
        public IActionResult ProcessarPedidos([FromBody] EntradaPedidos entrada)
        {
            if (entrada == null || entrada.Pedidos == null || !entrada.Pedidos.Any())
            {
                return BadRequest("Entrada inválida ou vazia.");
            }

            try
            {
                // Serializa o pedido para enviar ao RabbitMQ
                var mensagem = JsonSerializer.Serialize(entrada);
                _rabbitMQService.SendMessage("fila_otimizar", mensagem);

                return Ok(new
                {
                    Mensagem = "Pedido enviado para processamento",
                    Status = "Sucesso"
                });
            }
            catch (Exception ex)
            {
                return StatusCode(500, new { Mensagem = "Erro ao enviar pedido", Erro = ex.Message });
            }
        }
    }
}
