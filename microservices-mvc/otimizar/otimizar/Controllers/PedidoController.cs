using Microsoft.AspNetCore.Mvc;
using Otimizar.Models;
using Otimizar.Services;

namespace Otimizar.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class PedidoController : ControllerBase
    {
        private readonly AlocacaoCaixaService _alocacaoCaixaService;

        public PedidoController(AlocacaoCaixaService alocacaoCaixaService)
        {
            _alocacaoCaixaService = alocacaoCaixaService;
        }

        [HttpPost("processar-pedidos")]
        public IActionResult ProcessarPedidos([FromBody] EntradaPedidos entrada)
        {
            if (entrada == null || entrada.Pedidos == null || !entrada.Pedidos.Any())
            {
                return BadRequest("Entrada inválida ou vazia.");
            }

            var resultado = entrada.Pedidos
                .Select(pedido => new
                {
                    PedidoId = pedido.PedidoId,
                    Caixas = _alocacaoCaixaService.ProcessarPedido(pedido)
                })
                .ToList();

            return Ok(new
            {
                Mensagem = "Pedidos processados com sucesso",
                Pedidos = resultado
            });
        }
    }
}
