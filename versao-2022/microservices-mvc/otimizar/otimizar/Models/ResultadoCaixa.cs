namespace Otimizar.Models
{
    public class ResultadoCaixa
    {
        public string CaixaId { get; set; }
        public List<string> Produtos { get; set; }
        public string Observacao { get; set; } = string.Empty;
    }
}
