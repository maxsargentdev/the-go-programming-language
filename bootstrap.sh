#!/bin/bash

pushd ./the-go-programming-language

for n in {1..12}; do cobra-cli add "c1e$n"; done
for n in {1..5}; do cobra-cli add "c2e$n"; done
for n in {1..13}; do cobra-cli add "c3e$n"; done
for n in {1..14}; do cobra-cli add "c4e$n"; done
for n in {1..19}; do cobra-cli add "c5e$n"; done
for n in {1..5}; do cobra-cli add "c6e$n"; done
for n in {1..18}; do cobra-cli add "c7e$n"; done
for n in {1..15}; do cobra-cli add "c8e$n"; done
for n in {1..5}; do cobra-cli add "c9e$n"; done
for n in {1..4}; do cobra-cli add "c10e$n"; done
for n in {1..7}; do cobra-cli add "c11e$n"; done
for n in {1..13}; do cobra-cli add "c12e$n"; done
for n in {1..4}; do cobra-cli add "c13e$n"; done
