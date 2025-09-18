{ pkgs, ... }:
{
  projectRootFile = "flake.nix";

  programs = {
    deadnix.enable = true;
    statix.enable = true;
    nixfmt.enable = true;

    goimports = {
      enable = true;
      package = pkgs.writeShellScriptBin "goimports" ''
        export GOTOOLCHAIN=local
        exec ${pkgs.gotools}/bin/goimports "$@"
      '';
    };
    gofmt = {
      enable = true;
      package = pkgs.writeShellScriptBin "gofmt" ''
        export GOTOOLCHAIN=local
        exec ${pkgs.go}/bin/gofmt "$@"
      '';
    };
  };
}
