# Tutorial: Executando SonarQube com Maven Wrapper sem SCM (Git)

## Introdução
O SonarQube é uma poderosa ferramenta de análise estática de código que ajuda a identificar problemas de qualidade, vulnerabilidades e códigos duplicados. Quando usamos o Maven Wrapper (`mvnw`) para rodar o SonarQube, podemos nos deparar com problemas em projetos que não estão versionados com Git. Neste tutorial, aprenderemos a executar o SonarQube ignorando o SCM (Source Code Management), permitindo sua análise mesmo sem um repositório Git.

## Pré-requisitos
Antes de iniciar, verifique se você tem os seguintes requisitos atendidos:

- **Java JDK 11 ou superior** instalado.
- **Maven Wrapper (`mvnw`)** disponível no projeto.
- **SonarQube Server** em execução (local ou remoto).
- **Token de autenticação** gerado no SonarQube.

## Passo 1: Iniciar o SonarQube (se estiver rodando localmente)
Se estiver usando o SonarQube localmente, certifique-se de que o serviço está rodando. Caso esteja usando o Docker, inicie o contêiner com o seguinte comando:

```sh
 docker run -d --name sonar -p 9000:9000 sonarqube
```

Aguarde alguns segundos e acesse `http://localhost:9000` para garantir que o SonarQube está ativo.

## Passo 2: Configurar o SonarQube no `pom.xml`
Se o seu projeto ainda não estiver configurado para análise com o SonarQube, adicione o seguinte plugin ao `pom.xml` dentro da seção `<plugins>`:

```xml
<plugin>
    <groupId>org.sonarsource.scanner.maven</groupId>
    <artifactId>sonar-maven-plugin</artifactId>
    <version>3.9.1.2184</version>
</plugin>
```

## Passo 3: Executar a Análise sem SCM (Git)
Agora, para rodar a análise do SonarQube sem exigir um repositório Git, utilize o seguinte comando:

```sh
./mvnw sonar:sonar -Dsonar.scm.disabled=true -Dsonar.login=SEU_TOKEN_AQUI
```

> Substitua `SEU_TOKEN_AQUI` pelo token gerado no SonarQube (Configurações > Segurança > Tokens).

O parâmetro `-Dsonar.scm.disabled=true` instrui o SonarQube a ignorar a necessidade de um repositório Git, permitindo que a análise prossiga normalmente.

## Passo 4: Conferir o Resultado
Após a execução do comando, acesse o painel do SonarQube (`http://localhost:9000`) e verifique os resultados da análise na interface web.

Se a análise não aparecer no SonarQube, verifique os logs de saída do Maven (`mvnw`) e corrija qualquer erro identificado.

## Conclusão
Seguindo esses passos, você conseguirá rodar a análise do SonarQube em um projeto **Spring Boot** usando o Maven Wrapper, mesmo sem um repositório Git. Isso é útil para projetos em desenvolvimento local ou que ainda não foram versionados. 🚀

Caso tenha dúvidas ou problemas, verifique os logs do Maven e garanta que o SonarQube está rodando corretamente.

---

📌 *Dica:* Para rodar o SonarQube em um branch específico, use `-Dsonar.branch.name=nome-do-branch`. Exemplo:
```sh
./mvnw sonar:sonar -Dsonar.branch.name=develop -Dsonar.login=SEU_TOKEN_AQUI
```

Boa análise de código! ✅

