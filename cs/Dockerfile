FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build

WORKDIR /app

# Copy csproj and restore as distinct layers
COPY FlattenDirectory.csproj ./
RUN dotnet restore

# Copy the remaining source code
COPY FlattenDirectory.cs ./

# Build the application
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/runtime:5.0

WORKDIR /app

# Copy the published application from the build stage
COPY --from=build /app/out ./

ENTRYPOINT ["dotnet", "FlattenDirectory.dll"]
