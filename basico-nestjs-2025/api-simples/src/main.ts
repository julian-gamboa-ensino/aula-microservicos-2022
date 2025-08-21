import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // Versão da API (URL: /v1/...)
  app.setGlobalPrefix('v1');

  // Validação global (integra com class-validator / class-transformer)
  app.useGlobalPipes(new ValidationPipe({ whitelist: true, transform: true }));

  // Configuração do Swagger / OpenAPI
  const config = new DocumentBuilder()
    .setTitle('Nest Swagger Lab')
    .setDescription('API de exemplo para laboratório com NestJS + Swagger')
    .setVersion('1.0.0')
    .addBearerAuth(
      { type: 'http', scheme: 'bearer', bearerFormat: 'JWT' },
      'jwt',
    )
    .build();

  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('docs', app, document, {
    swaggerOptions: {
      persistAuthorization: true,
    },
    customSiteTitle: 'Nest Swagger Lab',
  });

  await app.listen(3000);
  console.log('HTTP: http://localhost:3000');
  console.log('Swagger UI: http://localhost:3000/docs');
}
bootstrap();

