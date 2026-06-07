{
  description = "Polyjuice development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux"; 
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
            go_1_26
            golangci-lint
            dbmate
        ];

        shellHook = ''
          echo "🧪 Polyjuice development environment"
          echo "Go: $(go version)"
          echo "DBMate: $(dbmate --version)"
          echo "Running docker compose"
          docker compose -f ./deployments/docker-compose.yml up -d
          echo "Happy Hacking 😎"
        '';
      };
    };
}

