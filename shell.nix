{ pkgs ? import <nixpkgs> {} }:

with pkgs;

mkShell {
  buildInputs = [
    gofumpt
    golint
    SDL2
  ];
}
