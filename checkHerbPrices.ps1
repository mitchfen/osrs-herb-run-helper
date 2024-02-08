$ProgressPreference = "SilentlyContinue"
$url = 'https://prices.runescape.wiki/api/v1/osrs/1h'
$headers = @{'User-Agent'='herb-prices-powershell-script'}
$numberofHerbsPerSeed = 9.423
$numberofPatches = 9
$highlightColor = "Green"

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

    $herbs = $herbs | Sort-Object -Property ExpectedProfit -Descending
    $herbs | Format-Table Herb, ExpectedProfit, SeedPrice, HerbPrice

    Write-Host "The above table assumes $numberOfHerbsPerSeed herbs harvested per seed."
    Write-Host "If you plant " -NoNewline
    Write-Host "$($herbs[0].Herb) " -NoNewline -ForegroundColor $highlightColor
    Write-Host "in $numberofPatches patches you can expect " -NoNewLine
    Write-Host "$([int]($herbs[0].ExpectedProfit / 1000) * $numberofPatches)K profit" -ForegroundColor $highlightColor

    $cost = $numberofPatches * $($herbs[0].SeedPrice)
    Write-Host "If you plant " -NoNewline
    Write-Host "$($herbs[0].Herb) " -NoNewline -ForegroundColor $highlightColor
    Write-Host "in $numberofPatches patches you need to harvest at least " -NoNewLine
    Write-Host "$([int]($cost / $($herbs[0].HerbPrice))) herbs" -ForegroundColor $highlightColor -NoNewLine
    Write-Host " to break even."
    Write-Host ""

} catch {
    Write-Error $_
}
