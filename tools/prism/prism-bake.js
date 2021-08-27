(function (Prism) {
	Prism.languages.bake = {
		"comment": /\/\/.*/,
		"raw-string": {
			pattern: /`[\s\S\n]*`/,
			greedy: true,
			alias: "string",
		},
		"keyword": /\b(package|import|func|rule|export|if|else|for|in|return|shell|let)\b/,
		"boolean": /\b(?:true|false)\b/,
		"number": /\b((0x[0-9a-fA-F]+)|(0[0-7]+i?)|(\d+([Ee]\d+)?i?)|(\d+[Ee][-+]\d+i?))\b/i,
		"function": /[a-zA-Z_]\w*(?=\()/,
		"punctuation": /[()[\]{}:,.$@]/,
		"operator": /--|\+\+|==|!=|<=|>=|&&|\|\||!|[+-/*%?]?=|\+|\-|\*|\/|%/,
		"builtin": /\b(context)\b/,
	};

	var evaluation = {
		"template-punctuation": {
			pattern: /^"|"$/,
			alias: "string",
		},
		"interpolation": {
			pattern: /((?:^|[^\\])(?:\\{2})*)\$\((?:[^()]|\((?:[^()]|\([^)]*\))*\))+\)/,
			lookbehind: true,
			inside: {
				"interpolation-punctuation": {
					pattern: /^\$\(|\)$/,
					alias: "punctuation",
				},
				rest: Prism.languages.bake
			}
		},
		"entity": /\\([0-7]{3}|[abfnrtv"$]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8})/,
		"string": /[\s\S]+/,
	};

	Prism.languages.insertBefore("bake", "raw-string", {
		"evaluated-string": {
			pattern: /"(?:\\[\s\S]|\$\((?:[^()]|\((?:[^()]|\([^)]*\))*\))+\)|(?!\$\()[^\\"])*"/,
			inside: evaluation,
		},
		// TODO: Doesn't work for context.shell.stdout.string etc. as shell* will be marked as a string. Add an assert for only whitespace before the shell expression
		"shell-block": {
			pattern: /(shell\s*\{)[^}]*\}/,
			lookbehind: true,
			alias: "string",
			inside: {
				"punctuation": {
					pattern: /^\{|}$/,
					alias: "punctuation",
				},
				...evaluation,
			},
		},
		// TODO: Doesn't work for context.shell.stdout.string etc. as shell* will be marked as a string. Add an assert for only whitespace before the shell expression
		"shell": {
			pattern: /(shell)(?!\s*\{).*/,
			greedy: true,
			lookbehind: true,
			inside: evaluation,
			alias: "string",
		},
	});
}(Prism));
