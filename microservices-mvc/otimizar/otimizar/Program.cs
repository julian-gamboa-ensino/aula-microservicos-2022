var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

// Dimensões das caixas disponíveis
var caixasDisponiveis = new List<Caixa>
{
    new Caixa { Id = 1, Altura = 30, Largura = 40, Comprimento = 80 },
    new Caixa { Id = 2, Altura = 80, Largura = 50, Comprimento = 40 },
    new Caixa { Id = 3, Altura = 50, Largura = 80, Comprimento = 60 }
};

// Novo endpoint POST para processar pedidos com alocação de caixas
app.MapPost("/processar-pedidos", (EntradaPedidos entrada) =>
{
    if (entrada == null || entrada.Pedidos == null || !entrada.Pedidos.Any())
    {
        return Results.BadRequest("Entrada inválida ou vazia.");
    }

    var resultado = new List<object>();

    foreach (var pedido in entrada.Pedidos)
    {
        var produtosNoPedido = new List<ResultadoCaixa>();

        foreach (var produto in pedido.Produtos)
        {
            var menorCaixa = RegraDeNegocio.EncontraMenorCaixa(produto, caixasDisponiveis);

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

        resultado.Add(new
        {
            PedidoId = pedido.PedidoId,
            Caixas = produtosNoPedido
        });
    }

    return Results.Ok(new
    {
        Mensagem = "Pedidos processados com sucesso",
        Pedidos = resultado
    });
})
.WithName("PostProcessarPedidos")
.WithOpenApi();



app.Run();


public class Dimensoes
{
    public int Altura { get; set; }
    public int Largura { get; set; }
    public int Comprimento { get; set; }
}

public class Produto
{
    public string ProdutoId { get; set; }
    public Dimensoes Dimensoes { get; set; }
}

public class Pedido
{
    public int PedidoId { get; set; }
    public List<Produto> Produtos { get; set; }
}

public class EntradaPedidos
{
    public List<Pedido> Pedidos { get; set; }
}

public class Caixa
{
    public int Id { get; set; }
    public int Altura { get; set; }
    public int Largura { get; set; }
    public int Comprimento { get; set; }

    public int Volume => Altura * Largura * Comprimento;
}

public class ResultadoCaixa
{
    public string CaixaId { get; set; }
    public List<string> Produtos { get; set; }
    public string Observacao { get; set; } = string.Empty;
}

public static class RegraDeNegocio
{
    public static bool CabeNaCaixa(Produto produto, Caixa caixa)
    {
        var dimensoesProduto = new[] { produto.Dimensoes.Altura, produto.Dimensoes.Largura, produto.Dimensoes.Comprimento };
        var rotacoes = Permutations(dimensoesProduto);
        foreach (var rotacao in rotacoes)
        {
            if (rotacao[0] <= caixa.Altura && rotacao[1] <= caixa.Largura && rotacao[2] <= caixa.Comprimento)
            {
                return true;
            }
        }
        return false;
    }

    public static Caixa EncontraMenorCaixa(Produto produto, List<Caixa> caixas)
    {
        var caixasOrdenadas = caixas.OrderBy(c => c.Volume).ToList();
        foreach (var caixa in caixasOrdenadas)
        {
            if (CabeNaCaixa(produto, caixa))
            {
                return caixa;
            }
        }
        return null;
    }

    public static IEnumerable<int[]> Permutations(int[] dimensoes)
    {
        return new List<int[]> {
            new int[] {dimensoes[0], dimensoes[1], dimensoes[2]},
            new int[] {dimensoes[0], dimensoes[2], dimensoes[1]},
            new int[] {dimensoes[1], dimensoes[0], dimensoes[2]},
            new int[] {dimensoes[1], dimensoes[2], dimensoes[0]},
            new int[] {dimensoes[2], dimensoes[0], dimensoes[1]},
            new int[] {dimensoes[2], dimensoes[1], dimensoes[0]},
        };
    }
}
