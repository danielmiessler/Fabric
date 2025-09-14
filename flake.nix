{
  description = "Fabric is an open-source framework for augmenting humans using AI. It provides a modular framework for solving specific problems using a crowdsourced set of AI prompts that can be used anywhere";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      systems,
      treefmt-nix,
      gomod2nix,
      ...
    }:
    let
      forAllSystems = nixpkgs.lib.genAttrs (import systems);

      getGoVersion = system: nixpkgs.legacyPackages.${system}.go_latest;

      treefmtEval = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        treefmt-nix.lib.evalModule pkgs (
          { ... }:
          {
            imports = [ ./nix/treefmt.nix ];
            # Set environment variable to prevent Go toolchain auto-download
            settings.global.excludes = [ ];
          }
        )
      );
    in
    {
      formatter = forAllSystems (system: treefmtEval.${system}.config.build.wrapper);

      checks = forAllSystems (system: {
        formatting = treefmtEval.${system}.config.build.check self;
      });

      devShells = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          goVersion = getGoVersion system;
          goEnv = gomod2nix.legacyPackages.${system}.mkGoEnv {
            pwd = ./.;
            go = goVersion;
          };
        in
        import ./nix/shell.nix {
          inherit pkgs goEnv goVersion;
          inherit (gomod2nix.legacyPackages.${system}) gomod2nix;
        }
      );

      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          goVersion = getGoVersion system;
        in
        {
          default = self.packages.${system}.fabric;
          fabric = pkgs.callPackage ./nix/pkgs/fabric {
            go = goVersion;
            inherit self;
            inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
          };
          inherit (gomod2nix.legacyPackages.${system}) gomod2nix;
        }
      );
    };
}
