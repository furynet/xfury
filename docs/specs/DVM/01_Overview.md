# **Overview**

The FURY-Network chain relies on external off-chain data of matches and other markets. To push this data reliably to the chain, some kind of origin verification is required. The `DVM module` essentially fills this role in the FURY-Network chain. The `DVM Module` verifies the following data:

- Market data pushed by the House to the chain
- Validity of Odds data using proposed ticket data during bet placement by user
