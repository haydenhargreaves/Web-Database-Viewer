{
  description = "Go development flake. Adjust as needed.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        # Define the development shell.
        # When you run `nix develop` (or direnv activates), you'll enter this shell.
        devShells.default = pkgs.mkShell {
          # List all the development tools you need available in this shell's PATH.
          packages = with pkgs; [
            go
            gopls
            go-tools
            gcc_multi
            glibc_multi
          ];
          
          name = "Go";

          # Define the shell that will be executed.
          # Here, we explicitly use zsh.
          # Note: pkgs.zsh needs to be included in `packages` or `nativeBuildInputs`
          # for it to be found in the shell's environment. `inherit pkgs.zsh;` is concise.
          inherit (pkgs) zsh;

          # Environment variables and commands to run when the shell starts.
          shellHook = ''
            # Use the .local directory instead of home
            export GOPATH="$HOME/.local/go"
            echo "Settings GOPATH to: $HOME/.local/go "

            export GOOS=linux
            export GOARCH=amd64
            export CGO_CFLAGS=-Wno-error=cpp;

            # Exec zsh to replace the current shell process with zsh.
            # This ensures your prompt and zsh configurations load correctly.
            exec zsh
          '';
        };
      }
    );
}
