# Проверяем, передан ли аргумент
if ($args.Length -eq 0) {
    Write-Host "Usage: .\rename_files.ps1 <directory>"
    exit 1
}

# Директория, переданная пользователем
$TARGET_DIR = $args[0]

# Проверяем, существует ли директория
if (-Not (Test-Path $TARGET_DIR -PathType Container)) {
    Write-Host "Error: Directory '$TARGET_DIR' does not exist."
    exit 1
}

# Удаляем .gen из имен файлов в указанной директории
Get-ChildItem -Path $TARGET_DIR -Recurse -File -Filter "*.gen.*" | ForEach-Object {
    # Новый путь к файлу без .gen
    $newName = $_.FullName -replace '\.gen', ''

    # Переименование файла
    Rename-Item $_.FullName -NewName $newName
    Write-Host "Renamed: $($_.FullName) -> $newName"
}

Write-Host "Done renaming files in directory: $TARGET_DIR"
