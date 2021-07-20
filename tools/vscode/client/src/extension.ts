import * as path from "path";
import {
  ExtensionContext,
  languages
} from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient;

export function activate(context: ExtensionContext) {
  languages.setLanguageConfiguration("bake", { wordPattern: /[a-zA-Z_]\w*/ });

  let serverModule = context.asAbsolutePath(
    path.join("server", "out", "server.js")
  );

  // The debug options for the server
  // --inspect=6009: runs the server in Node"s Inspector mode so VS Code can attach to the server for debugging
  let debugOptions = { execArgv: ["--nolazy", "--inspect=6009"] };

  // If the extension is launched in debug mode then the debug server options are used
  // Otherwise the run options are used
  let serverOptions: ServerOptions = {
    run: { module: serverModule, transport: TransportKind.ipc },
    debug: {
      module: serverModule,
      transport: TransportKind.ipc,
      options: debugOptions
    }
  };

  let clientOptions: LanguageClientOptions = {
    documentSelector: [
      { scheme: "file", language: "bake" },
      { scheme: "untitled", language: "bake" },
    ],
  };

  client = new LanguageClient(
    "bake",
    "Bake",
    serverOptions,
    clientOptions
  );

  client.start();
}

export function deactivate(): Thenable<void> | void {
  if (!client) {
    return;
  }

  return client.stop();
}
