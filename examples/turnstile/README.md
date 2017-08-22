# Turnstile

A [turnstile](https://en.wikipedia.org/wiki/Finite-state_machine#Example:_coin-operated_turnstile) is a a mechanical gate consisting of revolving horizontal arms fixed to a vertical post, allowing only one person at a time to pass through if the entry fee is paid.

This example is separated into three parts:

- [Basic](basic/): explains the basic usage of a state machine
- [Embedded](embedded/): embeds the state machine into the subject (turnstile) and exposes commands hiding the state machine
- [Pool](pool/): embeds a state machine pool into the subject. Although using a state machine should be safe for concurrent usage, delegates are out of control of this library and by using pools one can make sure that the state machine is in fact concurrent safe.
