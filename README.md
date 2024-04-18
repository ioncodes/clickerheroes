# Clicker Heroes
Basically a server that allows you to spoof the redemption code submission response. By doing so you can use internal cheat codes that unlock various things (gems, gold, ...). This doesn't require modifications to the game. Tested on `v1.0e12-4699`.  

![image](https://github.com/ioncodes/clickerheroes/assets/18533297/55f750a6-da3e-4b1c-bcbc-15e09cba0353)

## How?
The tool hosts a server that always returns the following response:
```json
{
    "success": "success",
    "response": "layle was here"
}
```

By doing so we trick the game into thinking whatever code we submitted is valid. It then proceeds to do local checks to grant you whatever you requested. The information of what you get, and how much of it you get is entirely stored in the code you submit. There's also a "cheat code" that gives you everything (`Z` in the list below). The pattern is quite simple: `aaaaaaaaaaaBBBBBBBBBC`, where:
* `a` can be ignored. We just need a 11 character long filler
* `B` is the amount. Max size/amount is `long` for Rubies and `BigNumber` (allows for scientific notation afaik) for Gold and Souls 
* `C` is the item type (`R` for Rubies, `G` for Gold, `S` for Souls, `Z` to unlock everything)

Specifically this means (`gluesniffer...` would be the redemption code):
```
gluesniffer1337G       -> 1337 gold
gluesniffer6969696969R -> 6969696969 rubies
gluesniffer420420420S  -> 420420420 souls
gluesniffer1e69S       -> 1e69 souls
gluesnifferZ           -> "cheat code"
```

What does the cheat code (`Z`) do?
* Completes the current zone
* Gives you `1e300` gold
* Gives you `1e5` souls
* Gives you `10000` rubies
* Gives you 2 world resets
* Whatever `zoneController.SetCurZoneByHeight` and `userDataController.EnableDebug` do

## Setup
1. Generate public and private keys for TLS and import them into your trust store. Set the common name to `savedgames.clickerheroes.com`
2. Copy the certificates to the project folder as `server.crt` and `server.key`
3. Create an entry in your hosts file for `127.0.0.1 savedgames.clickerheroes.com`
4. `go run .` and start the game
4. Profit...
