{
  pkgs,
  gomod2nix,
  goEnv,
  goVersion,
}:

{
  default = pkgs.mkShell {
    nativeBuildInputs = [
      goVersion
      pkgs.gopls
      pkgs.gotools
      pkgs.go-tools
      pkgs.goimports-reviser
      gomod2nix
      goEnv

      (pkgs.writeShellScriptBin "update-mod" ''
        go get -u
        go mod tidy
        gomod2nix generate --outdir nix/pkgs/fabric
      '')
    ];

    shellHook = ''
      export GOTOOLCHAIN=local
      echo -e "\033[0;32;4mHelper commands:\033[0m"
      echo "'update-mod' instead of 'go get -u && go mod tidy && gomod2nix generate --outdir nix/pkgs/fabric'"
    '';
  };
}
