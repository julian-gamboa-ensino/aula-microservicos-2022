using Otimizar.Models;
using Otimizar.Business;
using System.Collections.Generic;

namespace Otimizar.Services
{
    public class AlocacaoCaixaService
    {
        private readonly RegrasDeNegocio _regrasDeNegocio;
        private readonly List<Caixa> _caixasDisponiveis;

        public AlocacaoCaixaService(RegrasDeNegocio regrasDeNegocio)
        {
            _regrasDeNegocio = regrasDeNegocio;
            _caixasDisponiveis = new List<Caixa>
            {
                new Caixa { Id = 1, Altura = 30, Largura = 40, Comprimento = 80 },
                new Caixa { Id = 2, Altura = 80, Largura = 50, Comprimento = 40 },
                new Caixa { Id = 3, Altura = 50, Largura = 80, Comprimento = 60 }
            };
        }

        public List<ResultadoCaixa> ProcessarPedido(Pedido pedido)
        {
            var produtosNoPedido = new List<ResultadoCaixa>();

            foreach (var produto in pedido.Produtos)
            {
                var menorCaixa = _regrasDeNegocio.EncontraMenorCaixa(produto, _caixasDisponiveis);
                if (menorCaixa != null)
                {
                    produtosNoPedido.Add(new ResultadoCaixa
                    {
                        CaixaId = $"Caixa {menorCaixa.Id}",
                        Produtos = new List<string> { produto.ProdutoId }
                    });
                }
                else
                {
                    produtosNoPedido.Add(new ResultadoCaixa
                    {
                        CaixaId = "Nenhuma",
                        Produtos = new List<string> { produto.ProdutoId },
                        Observacao = "Produto não cabe em nenhuma caixa disponível."
                    });
                }
            }

            return produtosNoPedido;
        }
    }
}
