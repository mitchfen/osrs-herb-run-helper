$ProgressPreference = "SilentlyContinue"
$url = 'https://prices.runescape.wiki/api/v1/osrs/1h'
$headers = @{'User-Agent'='herb-prices-powershell-script'}
$numberofHerbsPerSeed = 9.423
$numberofPatches = 9

try {
    $ErrorActionPreference = "Stop"
    $herbs = Get-Content $(Join-Path $PSScriptRoot "herbs.json") -Raw | ConvertFrom-Json
    $response = Invoke-WebRequest -Uri $url -Method GET -Headers $headers
    $latestPrices = $response.Content | ConvertFrom-Json

    foreach($herb in $herbs) {
        $herbPrices = $latestPrices.data.$($herb.Id)
        $seedPrices = $latestPrices.data.$($herb.SeedId)
        $herb.SeedPrice = [int](($seedPrices.avgHighPrice + $seedPrices.avgLowPrice) / 2 )
        $herb.HighPriceAverage = $herbPrices.avgHighPrice
        $herb.LowPriceAverage = $herbPrices.avgLowPrice
        $herb.HerbPrice = [int](($herbPrices.avgHighPrice + $herbPrices.avgLowPrice) / 2 )
        $herb.ExpectedProfit = [int](($herb.HerbPrice * $numberOfHerbsPerSeed) - $herb.SeedPrice)
        $herb.ExpectedProfit = $herb.ExpectedProfit
    }

    Write-Host ""
    Write-Host "This table assumes $numberOfHerbsPerSeed herbs harvested per seed."
    $herbs = $herbs | Sort-Object -Property ExpectedProfit -Descending
    $herbs | Format-Table Herb, ExpectedProfit

    Write-Host "If you plant " -NoNewline
    Write-Host "$($herbs[0].Herb) " -NoNewline -ForegroundColor Magenta
    Write-Host "in $numberofPatches patches you can expect " -NoNewLine
    Write-Host "$([int]($herbs[0].ExpectedProfit / 1000) * $numberofPatches)K profit" -ForegroundColor Magenta
    Write-Host ""

} catch {
    Write-Error $_
}
