package main

import (
    "testing"
)

func TestSayHello(t *testing.T) {
    if SayHello(&Person{"Name", 0}) != "Hello Name" {
        t.Errorf("Error saying hello")
    }
}
