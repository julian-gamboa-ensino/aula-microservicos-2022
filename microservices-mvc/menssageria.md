## Reflexão: Comunicação via HTTP vs Mensageria

Neste projeto, a comunicação entre os microsserviços ocorre via HTTP, usando o modelo de rede interna do Docker, onde os serviços se comunicam diretamente entre si, como `http://login:80`. Isso representa um modelo de comunicação síncrona, onde os serviços aguardam a resposta de outros microsserviços para seguir o processamento.

**Desafio para os Alunos:**
- Quais são as vantagens e desvantagens dessa abordagem de comunicação direta via HTTP?
- Como a mensageria, como o **SNS** da AWS ou ferramentas equivalentes, pode ser usada para substituir ou complementar essa forma de comunicação?
- Como a mensageria assíncrona pode melhorar a resiliência, escalabilidade e desacoplamento dos microsserviços?

### Comunicação via HTTP
**Vantagens:**
- Simplicidade na implementação.
- Ferramentas comuns de debug, como o Swagger, tornam mais fácil visualizar e testar as interações entre serviços.

**Desvantagens:**
- Dependência direta entre os serviços (acoplamento mais forte).
- Pode aumentar a latência, pois o serviço chamado deve estar disponível para a resposta imediata.

### Comunicação via Mensageria (SNS, RabbitMQ, etc.)
**Vantagens:**
- Desacoplamento dos microsserviços, permitindo que se comuniquem de forma assíncrona.
- Resiliência e tolerância a falhas aumentada, já que os serviços não precisam estar online ao mesmo tempo.
- Escalabilidade: pode lidar com grandes volumes de mensagens sem sobrecarregar serviços individuais.

**Desvantagens:**
- Complexidade maior na implementação e monitoramento.
- Pode haver uma curva de aprendizado para entender e configurar os sistemas de mensageria.

**Perguntas para Reflexão:**
- Quando a comunicação assíncrona é mais adequada que a síncrona?
- Como a arquitetura baseada em eventos pode impactar a performance e manutenção do sistema?
