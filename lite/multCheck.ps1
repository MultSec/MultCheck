# Define function to check if file was removed
function Was-Removed {
    param (
        [string]$FilePath
    )

    $DeletionTime = 120  # Deletion time in seconds

    $StartTime = Get-Date
    while (((Get-Date) - $StartTime).TotalSeconds -lt $DeletionTime) {
        if (-not (Test-Path $FilePath)) {
            return $true
        }
        Start-Sleep -Seconds 1
    }
    return $false
}

# Define function to check original file
function Check-OriginalFile {
    param (
        [string]$FilePath,
        [string]$TempDirectory
    )

    # Check if file exists
    if (-not (Test-Path $FilePath)) {
        Write-Host "File not found at: $FilePath"
        exit
    }

    # Copy file to temporary directory and check if it's removed
    $TempFilePath = Join-Path -Path $TempDirectory -ChildPath (Split-Path $FilePath -Leaf)

    # Read file and write to temp directory
    try {
        $Data = [System.IO.File]::ReadAllBytes($FilePath)
    } catch {
        Write-Host "Error reading file: $_"
        exit
    }

    # Write file to temp directory
    try {
        [System.IO.File]::WriteAllBytes($TempFilePath, $Data)
    } catch {
        Write-Host "Error writing file: $_"
        exit
    }

    # Check if file was removed
    if (-not (Was-Removed -FilePath $TempFilePath)) {
        Write-Host "Clean"
        exit
    }
}

# Define function to split file
function Split-File {
    param (
        [string]$InputFile,
        [string]$OutputDirectory
    )

    # Read file
    try {
        $Data = [System.IO.File]::ReadAllBytes($InputFile)
    } catch {
        Write-Host "Error reading file: $_"
        return
    }

    $Mid = $Data.Length / 2
    $FirstHalf = $Data[0..($Mid - 1)]
    $SecondHalf = $Data[$Mid..($Data.Length - 1)]

    $FirstHalfPath = Join-Path -Path $OutputDirectory -ChildPath ("first_half_" + (Split-Path $InputFile -Leaf))
    $SecondHalfPath = Join-Path -Path $OutputDirectory -ChildPath ("second_half_" + (Split-Path $InputFile -Leaf))

    try {
        [System.IO.File]::WriteAllBytes($FirstHalfPath, $FirstHalf)
        [System.IO.File]::WriteAllBytes($SecondHalfPath, $SecondHalf)
    } catch {
        Write-Host "Error writing file: $_"
        return
    }

    return ,$FirstHalfPath, $SecondHalfPath
}

# Define main function
function multCheck {
    # Define script parameters
    $OriginalFilePath = "C:\Users\Public\Downloads\payload.exe"
    $TempDirectory = "C:\Users\Public\Documents"

    # Check original file
    Check-OriginalFile -FilePath $OriginalFilePath -TempDirectory $TempDirectory

    # Check for dynamic analysis
    $FirstHalfPath, $SecondHalfPath = Split-File -InputFile $OriginalFilePath -OutputDirectory $TempDirectory

    # Check if either file was removed
    if ((Was-Removed -FilePath $FirstHalfPath) -or (Was-Removed -FilePath $SecondHalfPath)) {
        Write-Host "Static"
        exit
    } else {
        Write-Host "Dynamic"
        exit
    }
}