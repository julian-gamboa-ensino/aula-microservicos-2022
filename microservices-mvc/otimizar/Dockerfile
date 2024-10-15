FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app

EXPOSE 5001
EXPOSE 5002
EXPOSE 5003

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["otimizar/otimizar.csproj", "otimizar/"]
RUN dotnet restore "otimizar/otimizar.csproj"
COPY . .
WORKDIR "/src/otimizar"


RUN dotnet build "otimizar.csproj" -o /app/build

FROM build AS publish
RUN dotnet publish "otimizar.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "otimizar.dll"]