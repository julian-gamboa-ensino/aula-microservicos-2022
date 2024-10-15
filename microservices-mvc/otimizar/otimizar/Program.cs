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

var summaries = new[]
{
    "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
};

// Endpoint GET existente para previsão do tempo
app.MapGet("/weatherforecast", () =>
{
    var forecast =  Enumerable.Range(1, 5).Select(index =>
        new WeatherForecast
        (
            DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
            Random.Shared.Next(-20, 55),
            summaries[Random.Shared.Next(summaries.Length)]
        ))
        .ToArray();
    return forecast;
})
.WithName("GetWeatherForecast")
.WithOpenApi();

// Novo endpoint POST para receber um JSON
app.MapPost("/otimizar", (Item item) =>
{
    if (item == null)
    {
        return Results.BadRequest("Dados inválidos.");
    }

    // Aqui você pode adicionar lógica de otimização ou processamento
    return Results.Ok(new { mensagem = "Dados recebidos com sucesso", item });
})
.WithName("PostOtimizar")
.WithOpenApi();

app.Run();

// Classe record existente para WeatherForecast
record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary)
{
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}

// Modelo de dados para o JSON
public class Item
{
    public int Id { get; set; }
    public string Nome { get; set; }
}
