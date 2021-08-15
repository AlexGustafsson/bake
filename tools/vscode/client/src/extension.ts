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

  let serverModule = "bagels";

  // If the extension is launched in debug mode then the debug server options are used
  // Otherwise the run options are used
  let serverOptions: ServerOptions = {
    run: {command: serverModule, transport: TransportKind.stdio},
    debug: {command: serverModule, transport: TransportKind.stdio},
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
