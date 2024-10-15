using Otimizar.Business;
using Otimizar.Services;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Adiciona os controladores
builder.Services.AddControllers(); // Esta linha é necessária para registrar os controladores


builder.Services.AddScoped<RegrasDeNegocio>();
builder.Services.AddScoped<AlocacaoCaixaService>();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.MapControllers();
app.Run();
