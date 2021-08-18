import * as path from "path";
import {
  ExtensionContext,
  languages,
  workspace
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

  const config = workspace.getConfiguration("bagels");
  const bagelsCommand = config.get<string>("command")!

  // If the extension is launched in debug mode then the debug server options are used
  // Otherwise the run options are used
  let serverOptions: ServerOptions = {
    run: {command: bagelsCommand, transport: TransportKind.stdio},
    debug: {command: bagelsCommand, transport: TransportKind.stdio},
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
