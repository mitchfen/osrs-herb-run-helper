$url = 'https://prices.runescape.wiki/api/v1/osrs/1h'

try {
    $herbsFile = Join-Path $PSScriptRoot "herbs.json"
    $herbs = Get-Content $herbsFile -Raw | ConvertFrom-Json

    # From https://oldschool.runescape.wiki/w/Herb#Farming_herbs
    $numberofHerbsPerSeed = 9.423 # Assumes magic secateurs and farming cape equipped 
    $numberofPatches = 8

    $headers = @{'User-Agent'='herb-prices-powershell-script'}
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

    Write-Host
    Write-Host "Here are the latest herb prices."
    Write-Host "This table assumes $numberOfHerbsPerSeed herbs harvested per seed."

    $herbs = $herbs | Sort-Object -Property ExpectedProfit -Descending
    $herbs | Format-Table Herb, ExpectedProfit, SeedPrice, HerbPrice
    Write-Host "If you plant $($herbs[0].Herb) in $numberofPatches patches you can expect " -NoNewLine
    Write-Host "$([int]($herbs[0].ExpectedProfit / 1000) * $numberofPatches)K profit" -ForegroundColor Green

} catch {
    Write-Error $_
}
