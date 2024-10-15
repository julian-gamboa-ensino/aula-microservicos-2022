using Otimizar.Models;
using System.Collections.Generic;
using System.Linq;

namespace Otimizar.Business
{
    public class RegrasDeNegocio
    {
        public bool CabeNaCaixa(Produto produto, Caixa caixa)
        {
            var dimensoesProduto = new[] { produto.Dimensoes.Altura, produto.Dimensoes.Largura, produto.Dimensoes.Comprimento };
            var rotacoes = Permutations(dimensoesProduto);
            return rotacoes.Any(rotacao =>
                rotacao[0] <= caixa.Altura && rotacao[1] <= caixa.Largura && rotacao[2] <= caixa.Comprimento);
        }

        public Caixa EncontraMenorCaixa(Produto produto, List<Caixa> caixas)
        {
            var caixasOrdenadas = caixas.OrderBy(c => c.Volume).ToList();
            return caixasOrdenadas.FirstOrDefault(caixa => CabeNaCaixa(produto, caixa));
        }

        private IEnumerable<int[]> Permutations(int[] dimensoes)
        {
            return new List<int[]>
            {
                new int[] {dimensoes[0], dimensoes[1], dimensoes[2]},
                new int[] {dimensoes[0], dimensoes[2], dimensoes[1]},
                new int[] {dimensoes[1], dimensoes[0], dimensoes[2]},
                new int[] {dimensoes[1], dimensoes[2], dimensoes[0]},
                new int[] {dimensoes[2], dimensoes[0], dimensoes[1]},
                new int[] {dimensoes[2], dimensoes[1], dimensoes[0]},
            };
        }
    }
}
