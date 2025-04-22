
4. Evite Reinventar a Roda

    Em vez de criar uma solução personalizada do zero, veja se um padrão existente resolve o problema de maneira eficaz. Isso pode poupar tempo e garantir que a solução seja robusta e testada.


https://en.wikipedia.org/wiki/Software_design_pattern

https://en.wikipedia.org/wiki/Decorator_pattern


docker run \
  -p 9090:9090 \
  -v ./volume-outubro-07/prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus



    | Padrão de Projeto | Tipo | Descrição | Exemplo de Uso |
|---|---|---|---|
| **Abstract Factory** | Criação | Fornece uma interface para criar famílias de objetos relacionados ou dependentes sem especificar suas classes concretas. |  Um sistema que precisa criar diferentes tipos de documentos (HTML, PDF, etc.) pode usar uma Abstract Factory para fornecer as classes concretas de cada tipo de documento. |
| **Builder** | Criação | Separa a construção de um objeto complexo da sua representação, permitindo que o mesmo processo de construção crie diferentes representações. |  Construir um objeto `Carro` com vários opcionais (motor, rodas, acessórios) de forma flexível. |
| **Factory Method** | Criação | Define uma interface para criar um objeto, mas deixa as subclasses decidirem qual classe instanciar. Permite a uma classe delegar a instanciação para subclasses. |  Um framework de logging que permite que as subclasses definam qual tipo de logger usar (console, arquivo, etc.). |
| **Prototype** | Criação | Especifica os tipos de objetos a serem criados usando uma instância prototípica e cria novos objetos copiando este protótipo. |  Clonar objetos complexos em um jogo, como unidades inimigas com diferentes características. |
| **Singleton** | Criação | Garante que uma classe tenha somente uma instância e fornece um ponto global de acesso para ela. |  Um gerenciador de configuração que deve ser acessado por todo o sistema. |
| **Adapter** | Estrutural | Converte a interface de uma classe em outra interface que os clientes esperam. Permite que classes com interfaces incompatíveis trabalhem juntas. |  Integrar um sistema legado com uma nova biblioteca que possui uma interface diferente. |
| **Bridge** | Estrutural | Desacopla uma abstração da sua implementação, permitindo que as duas variem independentemente. |  Um sistema de desenho que suporta diferentes tipos de formas (círculo, quadrado) e diferentes APIs de renderização (OpenGL, DirectX). |
| **Composite** | Estrutural | Compõe objetos em estruturas de árvore para representar hierarquias parte-todo. Permite que os clientes tratem objetos individuais e composições de objetos de maneira uniforme. |  Representar uma estrutura de diretórios e arquivos em um sistema de arquivos. |
| **Decorator** | Estrutural | Adiciona responsabilidades a um objeto dinamicamente. Fornece uma alternativa flexível à herança para estender a funcionalidade. |  Adicionar diferentes tipos de molhos e acompanhamentos a um pedido de pizza. |
| **Facade** | Estrutural | Fornece uma interface simplificada para um subsistema complexo. |  Esconder a complexidade de um sistema de banco de dados atrás de uma interface mais simples. |
| **Flyweight** | Estrutural | Usa compartilhamento para suportar grandes quantidades de objetos de granularidade fina de forma eficiente. |  Representar caracteres em um editor de texto, onde muitos caracteres são repetidos. |
| **Proxy** | Estrutural | Fornece um substituto ou espaço reservado para outro objeto para controlar o acesso a ele. |  Controlar o acesso a um objeto remoto ou carregar um objeto sob demanda. |
| **Chain of Responsibility** | Comportamental | Evita acoplar o remetente de uma solicitação ao seu receptor, dando a mais de um objeto a chance de tratar a solicitação. Encadeia os objetos receptores e passa a solicitação ao longo da cadeia até que um objeto a trate. |  Um sistema de aprovação de pedidos, onde diferentes níveis de aprovação são necessários dependendo do valor do pedido. |
| **Command** | Comportamental | Encapsula uma solicitação como um objeto, permitindo parametrizar clientes com diferentes solicitações, enfileirar ou registrar solicitações e suportar operações que podem ser desfeitas. |  Implementar operações de desfazer/refazer em um editor de texto. |
| **Interpreter** | Comportamental | Dada uma linguagem, define uma representação para sua gramática juntamente com um interpretador que usa a representação para interpretar sentenças na linguagem. |  Criar uma linguagem de consulta para um banco de dados. |
| **Iterator** | Comportamental | Fornece uma maneira de acessar os elementos de um objeto agregado sequencialmente sem expor sua representação subjacente. |  Percorrer os elementos de uma lista sem precisar conhecer sua implementação interna. |
| **Mediator** | Comportamental | Define um objeto que encapsula como um conjunto de objetos interage. Promove o acoplamento fraco ao evitar que os objetos se refiram uns aos outros explicitamente e permite variar suas interações independentemente. |  Controlar a comunicação entre diferentes janelas em uma interface gráfica. |
| **Memento** | Comportamental | Sem violar o encapsulamento, captura e externaliza um estado interno de um objeto para que o objeto possa ser restaurado para esse estado mais tarde. |  Salvar e restaurar o estado de um jogo. |
| **Observer** | Comportamental | Define uma dependência um-para-muitos entre objetos para que quando um objeto muda de estado, todos os seus dependentes sejam notificados e atualizados automaticamente. |  Notificar os usuários sobre novas postagens em um blog. |
| **State** | Comportamental | Permite que um objeto altere seu comportamento quando seu estado interno muda. O objeto parecerá ter mudado de classe. |  Um objeto `ContaBancaria` que pode estar nos estados "aberta", "bloqueada" ou "encerrada", cada um com comportamentos diferentes. |
| **Strategy** | Comportamental | Define uma família de algoritmos, encapsula cada um e os torna intercambiáveis. Permite que o algoritmo varie independentemente dos clientes que o usam. |  Implementar diferentes algoritmos de ordenação em uma aplicação. |
| **Template Method** | Comportamental | Define o esqueleto de um algoritmo em uma operação, postergando alguns passos para as subclasses. Permite que as subclasses redefinam certos passos de um algoritmo sem alterar a estrutura do algoritmo. |  Definir o processo geral de fazer uma bebida (preparar, misturar, servir) em uma classe abstrata, permitindo que as subclasses definam os ingredientes específicos. |
| **Visitor** | Comportamental | Representa uma operação a ser realizada nos elementos de uma estrutura de objeto. Permite definir uma nova operação sem alterar as classes dos elementos nos quais opera. |  Aplicar diferentes operações (formatação, cálculo de impostos) em uma estrutura de dados complexa, como uma árvore sintática. |

