# Define variables
$imageName = "my-dev-environment"
$dockerFilePath = "Dockerfile" # Specify the correct Dockerfile
$workDir = (Get-Location).Path # Use the current directory
$containerName = "dev-container"
$containerWorkDir = "/home/devuser/app"

# Check if the Docker image exists
docker images --format "{{.Repository}}:{{.Tag}}" | findstr "$imageName:latest"

if ($LASTEXITCODE -eq 0) {
    Write-Host "Docker image '$imageName:latest' finnes."
} else {
    Write-Host "Docker image '$imageName:latest' finnes IKKE, bygger bildet..."
    docker build -t $imageName -f Dockerfile .
}
# Run the Docker container
docker run -it --rm `
    --name $containerName `
    --hostname $containerName `
    -v "${workDir}:${containerWorkDir}" `
    -p 8080:8080 `
    -p 7331:7331 `
    $imageName
