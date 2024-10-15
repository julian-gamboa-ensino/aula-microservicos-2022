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

app.MapPost("/processar-pedidos", (EntradaPedidos entrada) =>
{
    if (entrada == null || entrada.Pedidos == null || !entrada.Pedidos.Any())
    {
        return Results.BadRequest("Entrada inválida ou vazia.");
    }

    foreach (var pedido in entrada.Pedidos)
    {
        Console.WriteLine($"Processando Pedido ID: {pedido.PedidoId}");
        foreach (var produto in pedido.Produtos)
        {
            Console.WriteLine($"Produto: {produto.ProdutoId}");
            Console.WriteLine($"Dimensões - Altura: {produto.Dimensoes.Altura}, Largura: {produto.Dimensoes.Largura}, Comprimento: {produto.Dimensoes.Comprimento}");
        }
    }

    return Results.Ok(new { mensagem = "Pedidos processados com sucesso", quantidadePedidos = entrada.Pedidos.Count });
})
.WithName("PostProcessarPedidos")
.WithOpenApi();


app.Run();

// Classe record existente para WeatherForecast
record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary)
{
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}



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
