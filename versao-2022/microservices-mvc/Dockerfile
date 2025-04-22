FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app

EXPOSE 5001
EXPOSE 5002
EXPOSE 5003

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["MyApi/MyApi.csproj", "MyApi/"]
RUN dotnet restore "MyApi/MyApi.csproj"
COPY . .
WORKDIR "/src/MyApi"

#RUN dotnet build "MyApi.csproj" -c Release -o /app/build

RUN dotnet build "MyApi.csproj" -o /app/build

FROM build AS publish
RUN dotnet publish "MyApi.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "MyApi.dll"]