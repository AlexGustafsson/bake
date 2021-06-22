#include <tree_sitter/parser.h>

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 13
#define STATE_COUNT 23
#define LARGE_STATE_COUNT 2
#define SYMBOL_COUNT 27
#define ALIAS_COUNT 0
#define TOKEN_COUNT 19
#define EXTERNAL_TOKEN_COUNT 0
#define FIELD_COUNT 0
#define MAX_ALIAS_SEQUENCE_LENGTH 5
#define PRODUCTION_ID_COUNT 1

enum {
  sym_identifier = 1,
  anon_sym_package = 2,
  aux_sym_package_statement_token1 = 3,
  anon_sym_import = 4,
  anon_sym_LPAREN = 5,
  anon_sym_COMMA = 6,
  anon_sym_RPAREN = 7,
  anon_sym_export = 8,
  anon_sym_func = 9,
  anon_sym_LBRACE = 10,
  anon_sym_RBRACE = 11,
  anon_sym_rule = 12,
  anon_sym_DOT_DOT_DOT = 13,
  anon_sym_DQUOTE = 14,
  aux_sym_string_token1 = 15,
  sym_number = 16,
  sym_escape_sequence = 17,
  sym_comment = 18,
  sym_program = 19,
  sym_statement = 20,
  sym_package_statement = 21,
  sym_import_statement = 22,
  sym_string = 23,
  aux_sym_program_repeat1 = 24,
  aux_sym_import_statement_repeat1 = 25,
  aux_sym_string_repeat1 = 26,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_identifier] = "identifier",
  [anon_sym_package] = "package",
  [aux_sym_package_statement_token1] = "package_statement_token1",
  [anon_sym_import] = "import",
  [anon_sym_LPAREN] = "(",
  [anon_sym_COMMA] = ",",
  [anon_sym_RPAREN] = ")",
  [anon_sym_export] = "export",
  [anon_sym_func] = "func",
  [anon_sym_LBRACE] = "{",
  [anon_sym_RBRACE] = "}",
  [anon_sym_rule] = "rule",
  [anon_sym_DOT_DOT_DOT] = "...",
  [anon_sym_DQUOTE] = "\"",
  [aux_sym_string_token1] = "string_token1",
  [sym_number] = "number",
  [sym_escape_sequence] = "escape_sequence",
  [sym_comment] = "comment",
  [sym_program] = "program",
  [sym_statement] = "statement",
  [sym_package_statement] = "package_statement",
  [sym_import_statement] = "import_statement",
  [sym_string] = "string",
  [aux_sym_program_repeat1] = "program_repeat1",
  [aux_sym_import_statement_repeat1] = "import_statement_repeat1",
  [aux_sym_string_repeat1] = "string_repeat1",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_identifier] = sym_identifier,
  [anon_sym_package] = anon_sym_package,
  [aux_sym_package_statement_token1] = aux_sym_package_statement_token1,
  [anon_sym_import] = anon_sym_import,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_export] = anon_sym_export,
  [anon_sym_func] = anon_sym_func,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [anon_sym_rule] = anon_sym_rule,
  [anon_sym_DOT_DOT_DOT] = anon_sym_DOT_DOT_DOT,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [aux_sym_string_token1] = aux_sym_string_token1,
  [sym_number] = sym_number,
  [sym_escape_sequence] = sym_escape_sequence,
  [sym_comment] = sym_comment,
  [sym_program] = sym_program,
  [sym_statement] = sym_statement,
  [sym_package_statement] = sym_package_statement,
  [sym_import_statement] = sym_import_statement,
  [sym_string] = sym_string,
  [aux_sym_program_repeat1] = aux_sym_program_repeat1,
  [aux_sym_import_statement_repeat1] = aux_sym_import_statement_repeat1,
  [aux_sym_string_repeat1] = aux_sym_string_repeat1,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym_identifier] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_package] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_package_statement_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_import] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_export] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_func] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_rule] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT_DOT_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_string_token1] = {
    .visible = false,
    .named = false,
  },
  [sym_number] = {
    .visible = true,
    .named = true,
  },
  [sym_escape_sequence] = {
    .visible = true,
    .named = true,
  },
  [sym_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_program] = {
    .visible = true,
    .named = true,
  },
  [sym_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_package_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_import_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_string] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_program_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_import_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_string_repeat1] = {
    .visible = false,
    .named = false,
  },
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
};

static const uint16_t ts_non_terminal_alias_map[] = {
  0,
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(12);
      if (lookahead == '"') ADVANCE(27);
      if (lookahead == '(') ADVANCE(27);
      if (lookahead == ')') ADVANCE(27);
      if (lookahead == ',') ADVANCE(27);
      if (lookahead == '.') ADVANCE(23);
      if (lookahead == '/') ADVANCE(24);
      if (lookahead == '\\') ADVANCE(4);
      if (lookahead == '{') ADVANCE(27);
      if (lookahead == '}') ADVANCE(27);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(25);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(26);
      if (lookahead != 0 &&
          lookahead != '\'') ADVANCE(27);
      END_STATE();
    case 1:
      if (lookahead == '"') ADVANCE(27);
      if (lookahead == '/') ADVANCE(24);
      if (lookahead == '\\') ADVANCE(4);
      if (lookahead != 0 &&
          lookahead != '\'') ADVANCE(27);
      END_STATE();
    case 2:
      if (lookahead == '/') ADVANCE(31);
      END_STATE();
    case 3:
      if (lookahead == '/') ADVANCE(14);
      if (lookahead != 0 &&
          lookahead != ' ') ADVANCE(15);
      END_STATE();
    case 4:
      if (lookahead == 'u') ADVANCE(5);
      if (lookahead == 'x') ADVANCE(10);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(30);
      if (lookahead != 0) ADVANCE(28);
      END_STATE();
    case 5:
      if (lookahead == '{') ADVANCE(9);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(7);
      END_STATE();
    case 6:
      if (lookahead == '}') ADVANCE(28);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(6);
      END_STATE();
    case 7:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(10);
      END_STATE();
    case 8:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(28);
      END_STATE();
    case 9:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(6);
      END_STATE();
    case 10:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(8);
      END_STATE();
    case 11:
      if (eof) ADVANCE(12);
      if (lookahead == '"') ADVANCE(20);
      if (lookahead == '(') ADVANCE(16);
      if (lookahead == ')') ADVANCE(18);
      if (lookahead == ',') ADVANCE(17);
      if (lookahead == '/') ADVANCE(2);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(19);
      END_STATE();
    case 12:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 13:
      ACCEPT_TOKEN(aux_sym_package_statement_token1);
      if (lookahead == '\n') ADVANCE(15);
      if (lookahead == ' ') ADVANCE(31);
      if (lookahead != 0) ADVANCE(13);
      END_STATE();
    case 14:
      ACCEPT_TOKEN(aux_sym_package_statement_token1);
      if (lookahead == '/') ADVANCE(13);
      if (lookahead != 0 &&
          lookahead != ' ') ADVANCE(15);
      END_STATE();
    case 15:
      ACCEPT_TOKEN(aux_sym_package_statement_token1);
      if (lookahead != 0 &&
          lookahead != ' ') ADVANCE(15);
      END_STATE();
    case 16:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 17:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 18:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 19:
      ACCEPT_TOKEN(sym_identifier);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(19);
      END_STATE();
    case 20:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 21:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (lookahead == '\n') ADVANCE(27);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(21);
      END_STATE();
    case 22:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (lookahead == '.') ADVANCE(27);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(27);
      END_STATE();
    case 23:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (lookahead == '.') ADVANCE(22);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(27);
      END_STATE();
    case 24:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (lookahead == '/') ADVANCE(21);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(27);
      END_STATE();
    case 25:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(25);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(27);
      END_STATE();
    case 26:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(26);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(27);
      END_STATE();
    case 27:
      ACCEPT_TOKEN(aux_sym_string_token1);
      if (lookahead != 0 &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(27);
      END_STATE();
    case 28:
      ACCEPT_TOKEN(sym_escape_sequence);
      END_STATE();
    case 29:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(28);
      END_STATE();
    case 30:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(29);
      END_STATE();
    case 31:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(31);
      END_STATE();
    default:
      return false;
  }
}

static bool ts_lex_keywords(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (lookahead == 'e') ADVANCE(1);
      if (lookahead == 'f') ADVANCE(2);
      if (lookahead == 'i') ADVANCE(3);
      if (lookahead == 'p') ADVANCE(4);
      if (lookahead == 'r') ADVANCE(5);
      END_STATE();
    case 1:
      if (lookahead == 'x') ADVANCE(6);
      END_STATE();
    case 2:
      if (lookahead == 'u') ADVANCE(7);
      END_STATE();
    case 3:
      if (lookahead == 'm') ADVANCE(8);
      END_STATE();
    case 4:
      if (lookahead == 'a') ADVANCE(9);
      END_STATE();
    case 5:
      if (lookahead == 'u') ADVANCE(10);
      END_STATE();
    case 6:
      if (lookahead == 'p') ADVANCE(11);
      END_STATE();
    case 7:
      if (lookahead == 'n') ADVANCE(12);
      END_STATE();
    case 8:
      if (lookahead == 'p') ADVANCE(13);
      END_STATE();
    case 9:
      if (lookahead == 'c') ADVANCE(14);
      END_STATE();
    case 10:
      if (lookahead == 'l') ADVANCE(15);
      END_STATE();
    case 11:
      if (lookahead == 'o') ADVANCE(16);
      END_STATE();
    case 12:
      if (lookahead == 'c') ADVANCE(17);
      END_STATE();
    case 13:
      if (lookahead == 'o') ADVANCE(18);
      END_STATE();
    case 14:
      if (lookahead == 'k') ADVANCE(19);
      END_STATE();
    case 15:
      if (lookahead == 'e') ADVANCE(20);
      END_STATE();
    case 16:
      if (lookahead == 'r') ADVANCE(21);
      END_STATE();
    case 17:
      ACCEPT_TOKEN(anon_sym_func);
      END_STATE();
    case 18:
      if (lookahead == 'r') ADVANCE(22);
      END_STATE();
    case 19:
      if (lookahead == 'a') ADVANCE(23);
      END_STATE();
    case 20:
      ACCEPT_TOKEN(anon_sym_rule);
      END_STATE();
    case 21:
      if (lookahead == 't') ADVANCE(24);
      END_STATE();
    case 22:
      if (lookahead == 't') ADVANCE(25);
      END_STATE();
    case 23:
      if (lookahead == 'g') ADVANCE(26);
      END_STATE();
    case 24:
      ACCEPT_TOKEN(anon_sym_export);
      END_STATE();
    case 25:
      ACCEPT_TOKEN(anon_sym_import);
      END_STATE();
    case 26:
      if (lookahead == 'e') ADVANCE(27);
      END_STATE();
    case 27:
      ACCEPT_TOKEN(anon_sym_package);
      END_STATE();
    default:
      return false;
  }
}

static const TSLexMode ts_lex_modes[STATE_COUNT] = {
  [0] = {.lex_state = 0},
  [1] = {.lex_state = 11},
  [2] = {.lex_state = 11},
  [3] = {.lex_state = 11},
  [4] = {.lex_state = 1},
  [5] = {.lex_state = 1},
  [6] = {.lex_state = 1},
  [7] = {.lex_state = 11},
  [8] = {.lex_state = 11},
  [9] = {.lex_state = 11},
  [10] = {.lex_state = 11},
  [11] = {.lex_state = 11},
  [12] = {.lex_state = 11},
  [13] = {.lex_state = 11},
  [14] = {.lex_state = 11},
  [15] = {.lex_state = 11},
  [16] = {.lex_state = 11},
  [17] = {.lex_state = 11},
  [18] = {.lex_state = 11},
  [19] = {.lex_state = 11},
  [20] = {.lex_state = 3},
  [21] = {.lex_state = 11},
  [22] = {.lex_state = 11},
};

static const uint16_t ts_parse_table[LARGE_STATE_COUNT][SYMBOL_COUNT] = {
  [0] = {
    [ts_builtin_sym_end] = ACTIONS(1),
    [sym_identifier] = ACTIONS(1),
    [anon_sym_package] = ACTIONS(1),
    [anon_sym_import] = ACTIONS(1),
    [anon_sym_LPAREN] = ACTIONS(1),
    [anon_sym_COMMA] = ACTIONS(1),
    [anon_sym_RPAREN] = ACTIONS(1),
    [anon_sym_export] = ACTIONS(1),
    [anon_sym_func] = ACTIONS(1),
    [anon_sym_LBRACE] = ACTIONS(1),
    [anon_sym_RBRACE] = ACTIONS(1),
    [anon_sym_rule] = ACTIONS(1),
    [anon_sym_DOT_DOT_DOT] = ACTIONS(1),
    [anon_sym_DQUOTE] = ACTIONS(1),
    [aux_sym_string_token1] = ACTIONS(1),
    [sym_number] = ACTIONS(1),
    [sym_escape_sequence] = ACTIONS(1),
    [sym_comment] = ACTIONS(3),
  },
  [1] = {
    [sym_program] = STATE(21),
    [sym_statement] = STATE(3),
    [sym_package_statement] = STATE(8),
    [sym_import_statement] = STATE(8),
    [aux_sym_program_repeat1] = STATE(3),
    [ts_builtin_sym_end] = ACTIONS(5),
    [anon_sym_package] = ACTIONS(7),
    [anon_sym_import] = ACTIONS(9),
    [sym_comment] = ACTIONS(11),
  },
};

static const uint16_t ts_small_parse_table[] = {
  [0] = 6,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(13), 1,
      ts_builtin_sym_end,
    ACTIONS(15), 1,
      anon_sym_package,
    ACTIONS(18), 1,
      anon_sym_import,
    STATE(2), 2,
      sym_statement,
      aux_sym_program_repeat1,
    STATE(8), 2,
      sym_package_statement,
      sym_import_statement,
  [21] = 6,
    ACTIONS(7), 1,
      anon_sym_package,
    ACTIONS(9), 1,
      anon_sym_import,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(21), 1,
      ts_builtin_sym_end,
    STATE(2), 2,
      sym_statement,
      aux_sym_program_repeat1,
    STATE(8), 2,
      sym_package_statement,
      sym_import_statement,
  [42] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(23), 1,
      anon_sym_DQUOTE,
    STATE(4), 1,
      aux_sym_string_repeat1,
    ACTIONS(25), 2,
      aux_sym_string_token1,
      sym_escape_sequence,
  [56] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(28), 1,
      anon_sym_DQUOTE,
    STATE(4), 1,
      aux_sym_string_repeat1,
    ACTIONS(30), 2,
      aux_sym_string_token1,
      sym_escape_sequence,
  [70] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(32), 1,
      anon_sym_DQUOTE,
    STATE(5), 1,
      aux_sym_string_repeat1,
    ACTIONS(34), 2,
      aux_sym_string_token1,
      sym_escape_sequence,
  [84] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(36), 3,
      ts_builtin_sym_end,
      anon_sym_package,
      anon_sym_import,
  [93] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(38), 3,
      ts_builtin_sym_end,
      anon_sym_package,
      anon_sym_import,
  [102] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(40), 3,
      ts_builtin_sym_end,
      anon_sym_package,
      anon_sym_import,
  [111] = 4,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(42), 1,
      anon_sym_RPAREN,
    ACTIONS(44), 1,
      anon_sym_DQUOTE,
    STATE(11), 1,
      sym_string,
  [124] = 4,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(46), 1,
      anon_sym_COMMA,
    ACTIONS(48), 1,
      anon_sym_RPAREN,
    STATE(13), 1,
      aux_sym_import_statement_repeat1,
  [137] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(50), 3,
      ts_builtin_sym_end,
      anon_sym_package,
      anon_sym_import,
  [146] = 4,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(46), 1,
      anon_sym_COMMA,
    ACTIONS(52), 1,
      anon_sym_RPAREN,
    STATE(15), 1,
      aux_sym_import_statement_repeat1,
  [159] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(54), 3,
      ts_builtin_sym_end,
      anon_sym_package,
      anon_sym_import,
  [168] = 4,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(56), 1,
      anon_sym_COMMA,
    ACTIONS(59), 1,
      anon_sym_RPAREN,
    STATE(15), 1,
      aux_sym_import_statement_repeat1,
  [181] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(61), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [189] = 3,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(44), 1,
      anon_sym_DQUOTE,
    STATE(19), 1,
      sym_string,
  [199] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(63), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [207] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(59), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [215] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(65), 1,
      aux_sym_package_statement_token1,
  [222] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(67), 1,
      ts_builtin_sym_end,
  [229] = 2,
    ACTIONS(11), 1,
      sym_comment,
    ACTIONS(69), 1,
      anon_sym_LPAREN,
};

static const uint32_t ts_small_parse_table_map[] = {
  [SMALL_STATE(2)] = 0,
  [SMALL_STATE(3)] = 21,
  [SMALL_STATE(4)] = 42,
  [SMALL_STATE(5)] = 56,
  [SMALL_STATE(6)] = 70,
  [SMALL_STATE(7)] = 84,
  [SMALL_STATE(8)] = 93,
  [SMALL_STATE(9)] = 102,
  [SMALL_STATE(10)] = 111,
  [SMALL_STATE(11)] = 124,
  [SMALL_STATE(12)] = 137,
  [SMALL_STATE(13)] = 146,
  [SMALL_STATE(14)] = 159,
  [SMALL_STATE(15)] = 168,
  [SMALL_STATE(16)] = 181,
  [SMALL_STATE(17)] = 189,
  [SMALL_STATE(18)] = 199,
  [SMALL_STATE(19)] = 207,
  [SMALL_STATE(20)] = 215,
  [SMALL_STATE(21)] = 222,
  [SMALL_STATE(22)] = 229,
};

static const TSParseActionEntry ts_parse_actions[] = {
  [0] = {.entry = {.count = 0, .reusable = false}},
  [1] = {.entry = {.count = 1, .reusable = false}}, RECOVER(),
  [3] = {.entry = {.count = 1, .reusable = false}}, SHIFT_EXTRA(),
  [5] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_program, 0),
  [7] = {.entry = {.count = 1, .reusable = true}}, SHIFT(20),
  [9] = {.entry = {.count = 1, .reusable = true}}, SHIFT(22),
  [11] = {.entry = {.count = 1, .reusable = true}}, SHIFT_EXTRA(),
  [13] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_program_repeat1, 2),
  [15] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_program_repeat1, 2), SHIFT_REPEAT(20),
  [18] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_program_repeat1, 2), SHIFT_REPEAT(22),
  [21] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_program, 1),
  [23] = {.entry = {.count = 1, .reusable = false}}, REDUCE(aux_sym_string_repeat1, 2),
  [25] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_string_repeat1, 2), SHIFT_REPEAT(4),
  [28] = {.entry = {.count = 1, .reusable = false}}, SHIFT(18),
  [30] = {.entry = {.count = 1, .reusable = true}}, SHIFT(4),
  [32] = {.entry = {.count = 1, .reusable = false}}, SHIFT(16),
  [34] = {.entry = {.count = 1, .reusable = true}}, SHIFT(5),
  [36] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_import_statement, 3),
  [38] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_statement, 1),
  [40] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_package_statement, 2),
  [42] = {.entry = {.count = 1, .reusable = true}}, SHIFT(7),
  [44] = {.entry = {.count = 1, .reusable = true}}, SHIFT(6),
  [46] = {.entry = {.count = 1, .reusable = true}}, SHIFT(17),
  [48] = {.entry = {.count = 1, .reusable = true}}, SHIFT(12),
  [50] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_import_statement, 4),
  [52] = {.entry = {.count = 1, .reusable = true}}, SHIFT(14),
  [54] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_import_statement, 5),
  [56] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_import_statement_repeat1, 2), SHIFT_REPEAT(17),
  [59] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_import_statement_repeat1, 2),
  [61] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_string, 2),
  [63] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_string, 3),
  [65] = {.entry = {.count = 1, .reusable = false}}, SHIFT(9),
  [67] = {.entry = {.count = 1, .reusable = true}},  ACCEPT_INPUT(),
  [69] = {.entry = {.count = 1, .reusable = true}}, SHIFT(10),
};

#ifdef __cplusplus
extern "C" {
#endif
#ifdef _WIN32
#define extern __declspec(dllexport)
#endif

extern const TSLanguage *tree_sitter_bake(void) {
  static const TSLanguage language = {
    .version = LANGUAGE_VERSION,
    .symbol_count = SYMBOL_COUNT,
    .alias_count = ALIAS_COUNT,
    .token_count = TOKEN_COUNT,
    .external_token_count = EXTERNAL_TOKEN_COUNT,
    .state_count = STATE_COUNT,
    .large_state_count = LARGE_STATE_COUNT,
    .production_id_count = PRODUCTION_ID_COUNT,
    .field_count = FIELD_COUNT,
    .max_alias_sequence_length = MAX_ALIAS_SEQUENCE_LENGTH,
    .parse_table = &ts_parse_table[0][0],
    .small_parse_table = ts_small_parse_table,
    .small_parse_table_map = ts_small_parse_table_map,
    .parse_actions = ts_parse_actions,
    .symbol_names = ts_symbol_names,
    .symbol_metadata = ts_symbol_metadata,
    .public_symbol_map = ts_symbol_map,
    .alias_map = ts_non_terminal_alias_map,
    .alias_sequences = &ts_alias_sequences[0][0],
    .lex_modes = ts_lex_modes,
    .lex_fn = ts_lex,
    .keyword_lex_fn = ts_lex_keywords,
    .keyword_capture_token = sym_identifier,
  };
  return &language;
}
#ifdef __cplusplus
}
#endif
