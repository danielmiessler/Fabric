{
  projectRootFile = "flake.nix";

  programs = {
    deadnix.enable = true;
    statix.enable = true;
    nixfmt.enable = true;

    goimports.enable = true;
    gofmt.enable = true;
  };
}
