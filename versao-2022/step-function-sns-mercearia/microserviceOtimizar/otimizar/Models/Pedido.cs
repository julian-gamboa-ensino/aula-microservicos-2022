namespace Otimizar.Models
{
    public class Pedido
    {
        public int PedidoId { get; set; }
        public List<Produto> Produtos { get; set; }
    }
}
