using Login.Services;
using RabbitMQ.Client;

var builder = WebApplication.CreateBuilder(args);

// Configura a URL do servidor para escutar na porta 80
builder.WebHost.UseUrls("http://*:80");

// Adiciona os controladores e serviços
builder.Services.AddControllers(); // Esta linha é necessária para registrar os controladores

// Configura Swagger para documentação da API
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Registro do RabbitMQService
builder.Services.AddScoped<RabbitMQService>(sp => new RabbitMQService("localhost"));

// Configuração da aplicação
var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

// Roteia as requisições para os controladores
app.MapControllers();


// Iniciar a aplicação
app.Run();

