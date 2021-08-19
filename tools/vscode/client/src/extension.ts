import { statSync, watchFile } from "fs";
import {
  ExtensionContext,
  languages,
  window,
  workspace
} from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind
} from "vscode-languageclient/node";

let client: LanguageClient;

export function activate(context: ExtensionContext) {
  languages.setLanguageConfiguration("bake", { wordPattern: /[a-zA-Z_]\w*/ });

  const config = workspace.getConfiguration("bagels");
  const bagelsCommand = config.get<string>("command")!

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

  // In debug mode, automatically restart the server on change
  const isDebugging = process.env.NODE_ENV === "dev";
  console.log("bagels - isDebugging", isDebugging)
  const serverExists = statSync(bagelsCommand).isFile()
  console.log("bagels - serverExists", serverExists)
  if (isDebugging && serverExists) {
    watchFile(bagelsCommand, async current => {
      if (current.isFile()) {
        window.setStatusBarMessage("bagels has changed, restarting it", 2000)
        await client.stop();
        client.start();
      }
    });
  }
}

export function deactivate(): Thenable<void> | void {
  if (!client) {
    return;
  }

  return client.stop();
}
