using Notificacao.Services; // Certifique-se de que este é o namespace correto


var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Adiciona o RabbitMQConsumerService como um serviço no contêiner de injeção de dependência

//builder.Services.AddSingleton<RabbitMQConsumerService>();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

// Inicia o consumo de mensagens ao iniciar a aplicação
//var rabbitMQConsumerService = app.Services.GetRequiredService<RabbitMQConsumerService>();
//rabbitMQConsumerService.StartConsuming();

app.Run();
