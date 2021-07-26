package parsing

import (
	"strings"

	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
)

func parseSourceFile(parser *Parser) (*nodes.SourceFile, error) {
	parser.require(lexing.ItemStartOfInput)

	sourceFile := nodes.CreateSourceFile(0)

	sourceFile.PackageDeclaration = parsePackageDeclaration(parser)

	for {
		token, ok := parser.expect(lexing.ItemComment)
		if ok {
			content := strings.Replace(token.Value, "//", "", 1)
			comment := nodes.CreateComment(nodes.NodePosition(token.Start), content)
			sourceFile.TopLevelDeclarations = append(sourceFile.TopLevelDeclarations, comment)
		} else {
			break
		}
	}

	parser.require(lexing.ItemEndOfInput)

	return sourceFile, nil
}

func parsePackageDeclaration(parser *Parser) *nodes.PackageDeclaration {
	if _, ok := parser.expectPeek(lexing.ItemKeywordPackage); !ok {
		return nil
	}

	startToken := parser.require(lexing.ItemKeywordPackage)
	identifier := parser.require(lexing.ItemIdentifier)
	return nodes.CreatePackageDeclaration(nodes.NodePosition(startToken.Start), identifier)
}
