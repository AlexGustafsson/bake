function _commaSeparated(rule) {
  return seq(rule, repeat(seq(',', rule)));
}

function commaSeparated(rule) {
  return optional(_commaSeparated(rule));
}

module.exports = grammar({
  name: 'bake',

  extras: $ => [
    $.comment,
  ],

  word: $ => $.identifier,

  rules: {
    program: $ => repeat($.statement),

    statement: $ => choice(
      $.import_statement,
      $.package_statement,
    ),

    package_statement: $ => seq(
      "package",
      /[^ ]+/,
    ),

    import_statement: $ => seq(
      'import',
      '(',
      commaSeparated($.string),
      ')',
    ),

    function_declaration: $ => seq(
      optional("export"),
      "func",
      $.identifier,
      optional(seq(
        '(',
        commaSeparated($.identifier),
        ')'
      )),
      '{',
      $.body,
      '}',
    ),

    rule_declaration: $ => seq(
      optional("export"),
      "rule",
      $.identifier,
      optional(seq(
        '(',
        commaSeparated($.identifier),
        ')'
      )),
      '{',
      $.body,
      '}',
    ),

    body: $ => repeat1(
      choice(
        "..."
      )
    ),

    identifier: $ => /[a-zA-Z_][a-zA-Z0-9_]*/,

    string: $ => choice(
      seq(
        '"',
        repeat(choice(
          token.immediate(prec(1, /[^'\\]+/)),
          $.escape_sequence
        )),
        '"'
      ),
    ),

    number: $ => token(/[0-9]+/),

    escape_sequence: $ => token.immediate(seq(
      '\\',
      choice(
        /[^xu0-7]/,
        /[0-7]{1,3}/,
        /x[0-9a-fA-F]{2}/,
        /u[0-9a-fA-F]{4}/,
        /u{[0-9a-fA-F]+}/
      )
    )),

    comment: $ => token(
      seq('//', /.*/)
    ),
  }

})
