namespace Otimizar.Models
{
    public class Caixa
    {
        public int Id { get; set; }
        public int Altura { get; set; }
        public int Largura { get; set; }
        public int Comprimento { get; set; }

        public int Volume => Altura * Largura * Comprimento;
    }
}
