from pygments.lexer import RegexLexer, bygroups
from pygments.token import *

import re

__all__ = ['BakeLexer']


class BakeLexer(RegexLexer):
    name = 'Bake'
    aliases = ['Bake']
    filenames = ['*.bke']
    flags = re.MULTILINE | re.UNICODE

    tokens = {
        'root': [
            (u'(//.*)', bygroups(Comment)),
            (u'(package)([ \\t]+)([a-zA-Z0-9_]+)',
             bygroups(Name.Namespace, String, String)),
            (u'(import)([ \\t]*)(\\()', bygroups(Name.Namespace,
             String, Punctuation), 'import__1'),
            (u'(import)([ \\t]*)(\\\")', bygroups(Name.Namespace,
             String, Punctuation), 'import__2'),
            (u'(export)?([ \\t]+)?(func)([ \\t]+)([a-zA-Z0-9_]+)', bygroups(
                Name.Function, String, Keyword, String, Keyword), 'function_declaration__1'),
            (u'(export)?([ \\t]+)?(rule)([ \\t]+)([a-zA-Z0-9_]+)', bygroups(
                Name.Function, String, Keyword, String, Keyword), 'rule_declaration__1'),
            (u'(@)', bygroups(Keyword), 'hook__1'),
            (u'([a-zA-Z0-9-_.]+)', bygroups(String)),
            (u'(\")', bygroups(Punctuation), 'string__1'),
            (u'(\\[)', bygroups(Punctuation), 'file_list__1'),
            (u'(\\{)', bygroups(Punctuation), 'rule__1'),
            (u'(:)', bygroups(Punctuation), 'rule__2'),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'file_list__1': [
            (u'([a-zA-Z0-9-_.]+)', bygroups(String)),
            (u'(\")', bygroups(Punctuation), 'string__1'),
            (u'(,)', bygroups(Punctuation)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'function_call__1': [
            (u'(\n|\r|\r\n)', bygroups(String), '#pop'),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'function_declaration__1': [
            (u'(\n|\r|\r\n)', bygroups(String), '#pop'),
            (u'(\\()', bygroups(Punctuation), 'function_declaration__2'),
            (u'(\\{)', bygroups(Punctuation), 'function_declaration__3'),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'function_declaration__2': [
            (u'([a-zA-Z0-9_]+)', bygroups(Name.Variable)),
            (u'(,)', bygroups(Punctuation)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'function_declaration__3': [
            (u'(^[ \\t]+)', bygroups(String)),
            (u'(//.*)', bygroups(Comment)),
            (u'(shell)([ \\t]+)', bygroups(Keyword, String), 'shell__1'),
            (u'([a-zA-Z_]+)', bygroups(Keyword), 'function_call__1'),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'hook__1': [
            (u'(\n|\r|\r\n)', bygroups(String), '#pop'),
            (u'([a-zA-Z0-9_]+)', bygroups(Keyword)),
            (u'(\\.)', bygroups(Punctuation)),
            (u'([^ \\t{\\n\\r]+)', bygroups(String)),
            (u'(\\{)', bygroups(Punctuation), 'hook__2'),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'hook__2': [
            (u'(^[ \\t]+)', bygroups(String)),
            (u'(//.*)', bygroups(Comment)),
            (u'(shell)([ \\t]+)', bygroups(Keyword, String), 'shell__1'),
            (u'([a-zA-Z_]+)', bygroups(Keyword), 'function_call__1'),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'import__1': [
            (u'(\")', bygroups(Punctuation), 'string__1'),
            (u'(//.*)', bygroups(Comment)),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'import__2': [
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule__1': [
            (u'(^[ \\t]+)', bygroups(String)),
            (u'(//.*)', bygroups(Comment)),
            (u'(shell)([ \\t]+)', bygroups(Keyword, String), 'shell__1'),
            (u'([a-zA-Z_]+)', bygroups(Keyword), 'function_call__1'),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule__2': [
            (u'(\n|\r|\r\n)', bygroups(String), '#pop'),
            (u'([a-zA-Z0-9_]+)', bygroups(Keyword)),
            (u'(\\()', bygroups(Punctuation), 'rule__3'),
            (u'(\\{)', bygroups(Punctuation), 'rule__4'),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule__3': [
            (u'(\")', bygroups(Punctuation), 'string__1'),
            (u'(\\b\\d+)', bygroups(Number)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule__4': [
            (u'(^[ \\t]+)', bygroups(String)),
            (u'(//.*)', bygroups(Comment)),
            (u'(shell)([ \\t]+)', bygroups(Keyword, String), 'shell__1'),
            (u'([a-zA-Z_]+)', bygroups(Keyword), 'function_call__1'),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule_declaration__1': [
            (u'(\n|\r|\r\n)', bygroups(String), '#pop'),
            (u'(\\[)', bygroups(Punctuation), 'rule_declaration__2'),
            (u'(\\{)', bygroups(Punctuation), 'rule_declaration__3'),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule_declaration__2': [
            (u'([a-zA-Z0-9_]+)', bygroups(Name.Variable)),
            (u'(,)', bygroups(Punctuation)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'rule_declaration__3': [
            (u'(^[ \\t]+)', bygroups(String)),
            (u'(//.*)', bygroups(Comment)),
            (u'(shell)([ \\t]+)', bygroups(Keyword, String), 'shell__1'),
            (u'([a-zA-Z_]+)', bygroups(Keyword), 'function_call__1'),
            (u'([^\\s\\n\\r])', bygroups(Generic.Error)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'shell__1': [
            (u'(\n|\r|\r\n)', bygroups(String), '#pop'),
            (u'(\\$[^ \\n\\r]+)', bygroups(Name.Variable)),
            (u'([^\\s\\n\\r]+)', bygroups(String)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ],
        'string__1': [
            (u'(\\\\(?:\\\\|\"))', bygroups(String)),
            (u'([^\"\\\\\\n\\r]+)', bygroups(String)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ]
    }
