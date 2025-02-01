# Define variables
$imageName = "my-dev-environment"
$dockerFilePath = "Dockerfile"
$workDir = (Get-Location).Path
$containerName = "dev-container"
$containerWorkDir = "/home/devuser/app"

# Check if the Docker image exists
Write-Host "Checking if image '$imageName' exists..."
$imageList = docker images --format "{{.Repository}}:{{.Tag}}" | Out-String

if ($imageList -match "$imageName:latest") {
    Write-Host "Docker image '$imageName:latest' exists."
} else {
    Write-Host "Docker image '$imageName:latest' does not exist. Building..."
    docker build -t $imageName -f $dockerFilePath $workDir

    if ($LASTEXITCODE -ne 0) {
        Write-Error "Failed to build Docker image. Exiting..."
        exit 1
    }

    Write-Host "Docker image '$imageName' built successfully."
}

# Run the Docker container
docker run -it --rm `
    --name $containerName `
    --hostname $containerName `
    -v "${workDir}:${containerWorkDir}" `
    -p 8080:8080 `
    -p 7331:7331 `
    --user 1000:1000 ` # Run the container as devuser (if that user exists in the image)
    $imageName