"use strict";

import * as vscode from "vscode";

function cleanupAmazonUrl(url: string) {
  let searchExpression = new RegExp("((https?://(www|smile)?.?)?amazon.de/)[^/]*/?(dp/[^/]*)");
  var matches = url.match(searchExpression);
  if (matches == null) {
    return url;
  }

  let results = Array.from(matches);
  if (results == null) {
    return url;
  }

  return results[1] + results[4];
}

export function activate(context: vscode.ExtensionContext) {
  const disposable = vscode.commands.registerCommand("shorty.amazon", function () {
    // Get the active text editor
    const editor = vscode.window.activeTextEditor;
    if (!editor) return;

    const document = editor.document;
    if (!document) return;

    const selection = editor.selection;
    const text = document.getText(selection);

    if (!text) return;

    let newLines: string[] = [];
    let lines = text.split("\n");
    for (let line of lines) {
      newLines.push(cleanupAmazonUrl(line));
    }

    const newText = newLines.join("\n");

    editor.edit((editBuilder) => {
      editBuilder.replace(selection, newText);
    });
  });

  context.subscriptions.push(disposable);
}
